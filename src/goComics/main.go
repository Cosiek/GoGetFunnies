package main

import (
	"io/ioutil"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


func getResponse(url string) *http.Response {
	fmt.Println("Ściągam " + url)
	resp, err := http.Get(url)
	if err != nil { panic(err) }
	// TODO: check response status
	return resp
}


func getHTML(url string)[]byte{
	resp := getResponse(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}


func getDocument(url string) (*goquery.Document, error){
	resp := getResponse(url)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	return doc, err
}


func main() {
	url := "http://buttersafe.com/"
	doc, err := getDocument(url)
	if err != nil { panic(err) }
	doc.Find("img").Each(func(i int, s *goquery.Selection){
		fmt.Println(s.Attr("src"))
	})
	fmt.Println("Done.")
}
