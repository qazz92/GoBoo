package crawler

import (
	"net/http"
	"github.com/suapapa/go_hangul/encoding/cp949"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func GETDoc(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer res.Body.Close()
	utfBody, err := cp949.NewReader(res.Body)
	if err != nil{

	}
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Fatal()
	}

	return doc
}
