/*
Package firetest provides utilities for Firebase testing

*/
package firetest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

var (
	missingJSONExtension = []byte("append .json to your request URI to use the REST API")
	missingBody          = []byte(`{"error":"Error: No data supplied."}`)
	invalidJSON          = []byte(`{"error":"Invalid data; couldn't parse JSON object, array, or value. Perhaps you're using invalid characters in your key names."}`)
	invalidAuth          = []byte(`{"error" : "Could not parse auth token."}`)
)

// Firetest is a Firebase server implementation
type Firetest struct {
	// URL of form http://ipaddr:port with no trailing slash
	URL string
	// Secret used to authenticate with server
	Secret string

	listener net.Listener
	db       *notifyDB

	requireAuth *int32
}

// New creates a new Firetest server
func New() *Firetest {
	secret := []byte(fmt.Sprint(time.Now().UnixNano()))
	return &Firetest{
		db:          newNotifyDB(),
		Secret:      base64.URLEncoding.EncodeToString(secret),
		requireAuth: new(int32),
	}
}

// Start starts the server
func (ft *Firetest) Start() {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		if l, err = net.Listen("tcp6", "[::1]:0"); err != nil {
			panic(fmt.Errorf("failed to listen on a port: %v", err))
		}
	}
	ft.listener = l

	s := http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ft.serveHTTP(w, req)
	})}
	go func() {
		if err := s.Serve(l); err != nil {
			log.Printf("error serving: %s", err)
		}

		ft.Close()
	}()
	ft.URL = "http://" + ft.listener.Addr().String()
}

// Close closes the server
func (ft *Firetest) Close() {
	if ft.listener != nil {
		ft.listener.Close()
	}
}

func (ft *Firetest) serveHTTP(w http.ResponseWriter, req *http.Request) {
	if !strings.HasSuffix(req.URL.Path, ".json") {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(missingJSONExtension))
		return
	}

	if atomic.LoadInt32(ft.requireAuth) == 1 {
		var authenticated bool
		authHeader := req.URL.Query().Get("auth")
		switch {
		case strings.Contains(authHeader, "."):
			authenticated = ft.validJWT(authHeader)
		default:
			authenticated = authHeader == ft.Secret
		}

		if !authenticated {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(invalidAuth)
			return
		}
	}

	switch req.Method {
	case "PUT":
		ft.set(w, req)
	case "PATCH":
		ft.update(w, req)
	case "POST":
		ft.create(w, req)
	case "GET":
		switch req.Header.Get("Accept") {
		case "text/event-stream":
			ft.sse(w, req)
		default:
			ft.get(w, req)
		}
	case "DELETE":
		ft.del(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("not implemented yet")
	}
}

func decodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}

func (ft *Firetest) validJWT(val string) bool {
	parts := strings.Split(val, ".")
	if len(parts) != 3 {
		return false
	}

	// validate header
	hb, err := decodeSegment(parts[0])
	if err != nil {
		log.Println("error decoding header", err)
		return false
	}
	var header map[string]string
	if err := json.Unmarshal(hb, &header); err != nil {
		log.Println("error unmarshaling header", err)
		return false
	}
	if header["alg"] != "HS256" || header["typ"] != "JWT" {
		return false
	}

	// validate claim
	cb, err := decodeSegment(parts[1])
	if err != nil {
		log.Println("error decoding claim", err)
		return false
	}
	var claim map[string]interface{}
	if err := json.Unmarshal(cb, &claim); err != nil {
		log.Println("error unmarshaling claim", err)
		return false
	}
	if e, ok := claim["exp"]; ok {
		// make sure not expired
		exp, ok := e.(float64)
		if !ok {
			log.Println("expiration not a number")
			return false
		}
		if int64(exp) < time.Now().Unix() {
			log.Println("token expired")
			return false
		}
	}
	// ensure uid present
	data, ok := claim["d"]
	if !ok {
		log.Println("missing data in claim")
		return false
	}

	d, ok := data.(map[string]interface{})
	if !ok {
		log.Println("claim['data'] is not map")
		return false
	}

	if _, ok := d["uid"]; !ok {
		log.Println("claim['data'] missing uid")
		return false
	}

	if sig, err := decodeSegment(parts[2]); err == nil {
		hasher := hmac.New(sha256.New, []byte(ft.Secret))
		signedString := strings.Join(parts[:2], ".")
		hasher.Write([]byte(signedString))

		if !hmac.Equal(sig, hasher.Sum(nil)) {
			log.Println("invalid jwt signature")
			return false
		}
	}

	return true
}

func (ft *Firetest) set(w http.ResponseWriter, req *http.Request) {
	body, v, ok := unmarshal(w, req.Body)
	if !ok {
		return
	}

	ft.Set(req.URL.Path, v)
	w.Write(body)
}

func (ft *Firetest) update(w http.ResponseWriter, req *http.Request) {
	body, v, ok := unmarshal(w, req.Body)
	if !ok {
		return
	}
	ft.Update(req.URL.Path, v)
	w.Write(body)
}

func (ft *Firetest) create(w http.ResponseWriter, req *http.Request) {
	_, v, ok := unmarshal(w, req.Body)
	if !ok {
		return
	}

	name := ft.Create(req.URL.Path, v)
	rtn := map[string]string{"name": name}
	if err := json.NewEncoder(w).Encode(rtn); err != nil {
		log.Printf("Error encoding json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ft *Firetest) del(w http.ResponseWriter, req *http.Request) {
	ft.Delete(req.URL.Path)
}

func (ft *Firetest) get(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	v := ft.Get(req.URL.Path)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("Error encoding json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (ft *Firetest) sse(w http.ResponseWriter, req *http.Request) {
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming is not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")

	path := sanitizePath(req.URL.Path)
	c := ft.db.watch(path)
	defer ft.db.stopWatching(path, c)

	d := eventData{Path: path, Data: ft.db.get(path)}
	s, err := json.Marshal(d)
	if err != nil {
		fmt.Printf("Error marshaling node %s\n", err)
	}
	fmt.Fprintf(w, "event: put\ndata: %s\n\n", s)
	f.Flush()

	httpCloser := w.(http.CloseNotifier).CloseNotify()
	for {
		select {
		case <-httpCloser:
			return
		case <-time.After(30 * time.Second):
			fmt.Fprintf(w, "event: keep-alive\ndata: null\n\n")
			f.Flush()
			continue
		case n, ok := <-c:
			if !ok {
				return
			}

			s, err := json.Marshal(n.Data)
			if err != nil {
				fmt.Printf("Error marshaling node %s\n", err)
				continue
			}

			fmt.Fprintf(w, "event: %s\ndata: %s\n\n", n.Name, s)
			f.Flush()
		}
	}
}

func sanitizePath(p string) string {
	// remove slashes from the front and back
	//	/foo/.json -> foo/.json
	s := strings.Trim(p, "/")

	// remove .json extension
	//	foo/.json -> foo/
	s = strings.TrimSuffix(s, ".json")

	// trim an potential trailing slashes
	//	foo/ -> foo
	return strings.TrimSuffix(s, "/")
}

func unmarshal(w http.ResponseWriter, r io.Reader) ([]byte, interface{}, bool) {
	body, err := ioutil.ReadAll(r)
	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(missingBody)
		return nil, nil, false
	}

	var v interface{}
	if err := json.Unmarshal(body, &v); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(invalidJSON)
		return nil, nil, false
	}
	return body, v, true
}
