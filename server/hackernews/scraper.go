package hackernews

import (
	"github.com/PuerkitoBio/goquery"

	"log"
)

func GetPageBody(url string) ([]byte, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	html, err := doc.Find("body").Html()
	if err != nil {
		return nil, err
	}

	return []byte(html), nil
}
