package main

import (
	// "fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// comic functions

func Buttersafe(date time.Time, comic Comic)string{
	url := comic.Url

	doc, err := GetDocument(url)
	if err != nil { panic(err) }

	imgUrl := "#"
	title := "Fial :("
	doc.Find("img").Each(func(i int, s *goquery.Selection){
		_url, _ := s.Attr("src")
		if strings.Contains(_url, "buttersafe.com/comics/"){
			imgUrl = _url
			title, _ = s.Attr("alt")
		}
	})
	ctx := StdComicTemplateCtx{comic, imgUrl, title}
	return renderStandardTemplate(ctx)
}


func HagarTheHorrible(date time.Time, comic Comic)string{
	url := comic.Url
	imgSources := GetImagesSrcList(url)
	for _, imgUrl := range imgSources{
		if strings.Contains(imgUrl, "safr.kingfeatures.com/"){
			ctx := StdComicTemplateCtx{comic, imgUrl, ""}
			return renderStandardTemplate(ctx)
		}
	}
	return "Nope :("
}


func GoComics(date time.Time, comic Comic)string{
	// make tight url for passed date
	// (like https://www.gocomics.com/calvinandhobbes/2018/06/16)
	url := comic.Url + date.Format("2006/01/02")
	// get html
	doc, err := GetDocument(url)
	if err != nil { panic(err) }
	// get picture data
	found := doc.Find("meta[name$=image]")
	node := found.Nodes[0]
	var imgUrl string
	for i := 0; i < len(node.Attr); i++{
		if node.Attr[i].Key == "content"{
			imgUrl = node.Attr[i].Val
		}
	}
	// render standard template
	ctx := StdComicTemplateCtx{comic, imgUrl, ""}
	return renderStandardTemplate(ctx)
}


func Xkcd(date time.Time, comic Comic)string{
	// get document
	doc, err := GetDocument(comic.Url)
	if err != nil { panic(err) }
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
	ctx := StdComicTemplateCtx{comic, imgUrl, alt + " - " + title}
	return renderStandardTemplate(ctx)
}


func Sinfest(date time.Time, comic Comic)string{
	// build url
	imgUrl := "http://www.sinfest.net/btphp/comics/" + date.Format("2006-01-02") + ".gif"
	// render standard template
	ctx := StdComicTemplateCtx{comic, imgUrl, ""}
	return renderStandardTemplate(ctx)
}
