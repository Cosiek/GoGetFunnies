package main

import (
	"errors"
	// "fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// comic functions

func Buttersafe(date time.Time, comic Comic)(string, error){
	url := comic.Url

	doc, err := GetDocument(url)
	if err != nil {
		ctx := StdComicTemplateCtx{comic, "", "", err.Error()}
		return renderStandardTemplate(ctx), err
	}

	imgUrl := "#"
	title := "Fial :("
	doc.Find("img").Each(func(i int, s *goquery.Selection){
		_url, _ := s.Attr("src")
		if strings.Contains(_url, "buttersafe.com/comics/"){
			imgUrl = _url
			title, _ = s.Attr("alt")
		}
	})
	ctx := StdComicTemplateCtx{comic, imgUrl, title, ""}
	return renderStandardTemplate(ctx), nil
}


func Abstrusegoose(date time.Time, comic Comic)(string, error){
	doc, err := GetDocument(comic.Url)
	if err != nil {
		ctx := StdComicTemplateCtx{comic, "", "", err.Error()}
		return renderStandardTemplate(ctx), err
	}

	imgUrl := "#"
	title := "Fial :("
	alt := "Alt"
	doc.Find("img").Each(func(i int, s *goquery.Selection){
		_url, _ := s.Attr("src")
		if strings.Contains(_url, "https://abstrusegoose.com/strips"){
			imgUrl = _url
			title, _ = s.Attr("title")
			alt, _ = s.Attr("alt")
		}
	})
	ctx := StdComicTemplateCtx{comic, imgUrl, title + " - " + alt, ""}
	return renderStandardTemplate(ctx), nil
}


func Oglaf(date time.Time, comic Comic)(string, error){
	doc, err := GetDocument(comic.Url)
	if err != nil {
		ctx := StdComicTemplateCtx{comic, "", "", err.Error()}
		return renderStandardTemplate(ctx), err
	}
	// get node
	found := doc.Find("#strip")
	node := found.Nodes[0]
	// get data
	var imgUrl, title, alt string
	for i := 0; i < len(node.Attr); i++{
		if node.Attr[i].Key == "src"{
			imgUrl = node.Attr[i].Val
		} else if node.Attr[i].Key == "alt"{
			alt = node.Attr[i].Val
		} else if node.Attr[i].Key == "title"{
			title = node.Attr[i].Val
		}
	}
	ctx := StdComicTemplateCtx{comic, imgUrl, title + " - " + alt, ""}
	return renderStandardTemplate(ctx), nil
}


func HagarTheHorrible(date time.Time, comic Comic)(string, error){
	imgSources := GetImagesSrcList(comic.Url)
	for _, imgUrl := range imgSources{
		if strings.Contains(imgUrl, "safr.kingfeatures.com/"){
			ctx := StdComicTemplateCtx{comic, imgUrl, "", ""}
			return renderStandardTemplate(ctx), nil
		}
	}
	errorMsg := "No image url containing \"safr.kingfeatures.com/\" found."
	ctx := StdComicTemplateCtx{comic, "#", "", errorMsg}
	return renderStandardTemplate(ctx), errors.New(errorMsg)
}


func GoComics(date time.Time, comic Comic)(string, error){
	// make right url for passed date
	// (like https://www.gocomics.com/calvinandhobbes/2018/06/16)
	url := comic.Url + date.Format("2006/01/02")
	// get html
	doc, err := GetDocument(url)
	if err != nil {
		ctx := StdComicTemplateCtx{comic, "", "", err.Error()}
		return renderStandardTemplate(ctx), err
	}
	// get picture data
	found := doc.Find("meta[name$=image]")
	if len(found.Nodes) == 0 {
		errorMsg := "Nie znalazłem obrazka. :("
		ctx := StdComicTemplateCtx{comic, "", "", errorMsg}
		return renderStandardTemplate(ctx), errors.New(errorMsg)
	}
	node := found.Nodes[0]
	var imgUrl string
	for i := 0; i < len(node.Attr); i++{
		if node.Attr[i].Key == "content"{
			imgUrl = node.Attr[i].Val
		}
	}
	// render standard template
	ctx := StdComicTemplateCtx{comic, imgUrl, "", ""}
	return renderStandardTemplate(ctx), nil
}


func Xkcd(date time.Time, comic Comic)(string, error){
	// get document
	doc, err := GetDocument(comic.Url)
	if err != nil {
		ctx := StdComicTemplateCtx{comic, "", "", err.Error()}
		return renderStandardTemplate(ctx), err
	}
	// get node
	found := doc.Find("img")
	node := found.Nodes[1]
	// get data
	var imgUrl, title, alt string
	for i := 0; i < len(node.Attr); i++{
		if node.Attr[i].Key == "src"{
			imgUrl = "https:" + node.Attr[i].Val
		} else if node.Attr[i].Key == "alt"{
			alt = node.Attr[i].Val
		} else if node.Attr[i].Key == "title"{
			title = node.Attr[i].Val
		}
	}
	// render standard template
	ctx := StdComicTemplateCtx{comic, imgUrl, alt + " - " + title, ""}
	return renderStandardTemplate(ctx), nil
}


func Sinfest(date time.Time, comic Comic)(string, error){
	// build url
	imgUrl := "http://www.sinfest.net/btphp/comics/" + date.Format("2006-01-02") + ".gif"
	// render standard template
	ctx := StdComicTemplateCtx{comic, imgUrl, "", ""}
	return renderStandardTemplate(ctx), nil
}


func TheSystem(date time.Time, comic Comic)(string, error){
	// get document
	doc, err := GetDocument(comic.Url)
	if err != nil {
		ctx := StdComicTemplateCtx{comic, "", "", err.Error()}
		return renderStandardTemplate(ctx), err
	}
	// get node
	found := doc.Find("img")
	node := found.Nodes[0]
	// get data
	var imgUrl, alt string
	for i := 0; i < len(node.Attr); i++{
		if node.Attr[i].Key == "src"{
			imgUrl = node.Attr[i].Val
		} else if node.Attr[i].Key == "alt"{
			alt = node.Attr[i].Val
		}
	}
	// render standard template
	ctx := StdComicTemplateCtx{comic, imgUrl, alt, ""}
	return renderStandardTemplate(ctx), nil
}
