package main

import (
	"io"
	"io/ioutil"
	//"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)


func GetResponse(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil { panic(err) }
	// TODO: check response status
	return resp
}


func GetHTML(url string)[]byte{
	resp := GetResponse(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}


func GetDocument(url string) (*goquery.Document, error){
	resp := GetResponse(url)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	return doc, err
}


func GetImagesSrcList(url string)[]string{
	doc, err := GetDocument(url)
	if err != nil { panic(err) }

	slice := make([]string, 0)
	doc.Find("img").Each(func(i int, s *goquery.Selection){
		imgUrl, _ := s.Attr("src")
		slice = append(slice, imgUrl)
	})
	return slice
}

func CopyFile(src, dst string) error {
    in, err := os.Open(src)
    if err != nil { return err }

    out, err := os.Create(dst)
    if err != nil { return err }

    _, err = io.Copy(out, in)
    if err != nil { return err }

		in.Close()
    out.Close()
		return nil
}
