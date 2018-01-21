package proxy

import (
	"fmt"
	"net/http"
	"bufio"
)

func Article(w http.ResponseWriter, r *http.Request) {
	//url := chi.URLParam(request, "q") // URL encoded url we want to proxy serve
	url := "http://www.bbc.com/news/health-42736764"
	fmt.Println("url", url)
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
