package proxy

import (
	"net/http"
	"bufio"
)

func Article(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	url := queryValues.Get("q")
	if url == "" {
		http.Error(w, "Must specify the url to request", 400)
	}
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to retrieve article", 500)
	}
	if response.StatusCode >= 400 {
		if err != nil {
			http.Error(w, response.Status, response.StatusCode)
		}
	}
	reader := bufio.NewReader(response.Body)
	reader.WriteTo(w)
}
