package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)


func main() {
	url := "http://buttersafe.com/"
	doc, err := GetDocument(url)
	if err != nil { panic(err) }
	doc.Find("img").Each(func(i int, s *goquery.Selection){
		fmt.Println(s.Attr("src"))
	})
	fmt.Println("Done.")
}
