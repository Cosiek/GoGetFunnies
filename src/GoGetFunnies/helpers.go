package main

import (
	"bytes"
	"io"
	"io/ioutil"
	//"fmt"
	"net/http"
	"os"
	"text/template"

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

func renderStandardTemplate(ctx StdComicTemplateCtx)string{
	templ, err := template.New("segment").Parse(SEGMENT_TEMPLATE)
	if err != nil { panic(err) }
	var rendered bytes.Buffer
	err = templ.Execute(&rendered, ctx)
	if err != nil { panic(err) }
	return rendered.String()
}


func renderStd(comic Comic, imgSrc string, title string, errMsg string)(string){
	ctx := StdComicTemplateCtx{comic, imgSrc, title, errMsg}
	return renderStandardTemplate(ctx)
}


func writeCssFile(targetDirName string){
	targetFilePath := targetDirName + "comics.css"
	// check if file already exists
	if _, err := os.Stat(targetFilePath); err == nil {
	  return
	}
	// creata a file
	outFile, err := os.Create(targetFilePath)
	if err != nil { panic(err) }
	defer outFile.Close()
	// write
	outFile.WriteString(CSS)
	outFile.Sync()
}
