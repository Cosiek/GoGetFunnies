package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// comic functions

func Buttersafe(date time.Time, comic Comic)string{
	url := comic.Url
	fmt.Printf("buttersafe....")

	doc, err := GetDocument(url)
	if err != nil { panic(err) }
	fmt.Println("OK")

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
	fmt.Printf("Hagar the Horrible....")
	imgSources := GetImagesSrcList(url)
	for _, imgUrl := range imgSources{
		if strings.Contains(imgUrl, "safr.kingfeatures.com/"){
			fmt.Println("OK")
			ctx := StdComicTemplateCtx{comic, imgUrl, ""}
			return renderStandardTemplate(ctx)
		}
	}
	fmt.Println("Fial")
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
