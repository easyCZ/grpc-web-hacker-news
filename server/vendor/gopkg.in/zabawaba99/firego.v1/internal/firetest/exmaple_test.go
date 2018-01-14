package firetest_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/zabawaba99/firego/internal/firetest"
)

func Example() {
	ft := firetest.New()
	defer ft.Close()

	ft.Start()

	resp, err := http.Post(ft.URL+"/foo.json", "application/json", strings.NewReader(`{"bar":true}`))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Post Resp: %s\n", string(b))

	v := ft.Get("foo/bar")
	log.Printf("Value: %v", v)
}
