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
	for _, element := range imgSources{
		if strings.Contains(element, "safr.kingfeatures.com/"){
			fmt.Println("OK")
			return element
		}
	}
	fmt.Println("Fial")
	return "Nope :("
}
