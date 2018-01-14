package firego

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type aBool struct {
	addr *int32
}

func (a *aBool) set(v bool) {
	val := int32(0)
	if v {
		val = 1
	}
	atomic.StoreInt32(a.addr, val)
}

func (a *aBool) val() bool {
	v := atomic.LoadInt32(a.addr)
	if v == 1 {
		return true
	}
	return false
}

func newABool() *aBool {
	return &aBool{
		addr: new(int32),
	}
}

func TestTransaction(t *testing.T) {
	storedVal := newABool()
	hitConflict := newABool()

	fbCounter := new(int64)
	atomic.StoreInt64(fbCounter, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		counter := atomic.LoadInt64(fbCounter)
		val := strconv.FormatInt(counter, 10)
		valBytes := []byte(val)
		etag := base64.StdEncoding.EncodeToString(valBytes)

		if req.Method == http.MethodGet {
			w.Header().Set("Etag", etag)
			w.Write(valBytes)
			return
		}

		if req.Header.Get("if-match") != etag {
			hitConflict.set(true)
			w.Header().Set("Etag", etag)
			w.WriteHeader(http.StatusConflict)
			w.Write(valBytes)
			return
		}

		storedVal.set(true)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	fb := New(server.URL, nil)
	err := fb.Transaction(func(currentSnapshot interface{}) (interface{}, error) {
		counter, ok := currentSnapshot.(float64)
		require.True(t, ok, "counter is not of type float64")
		if !hitConflict.val() {
			// set some random value so that we can  test out the conflict logic
			atomic.StoreInt64(fbCounter, 123)
		}
		return counter + 1, nil
	})
	assert.NoError(t, err)
	assert.True(t, storedVal.val())
	assert.True(t, hitConflict.val())
}
