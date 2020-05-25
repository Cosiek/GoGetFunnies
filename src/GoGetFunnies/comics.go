package main

import (
	"errors"
	//"fmt"
	//"reflect"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// comic functions

func Buttersafe(date time.Time, comic Comic)(string, error){
	url := comic.Url

	doc, err := GetDocument(url)
	if err != nil {
		return renderStd(comic, "", "", err.Error()), err
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
	return renderStd(comic, imgUrl, title, ""), nil
}


func Abstrusegoose(date time.Time, comic Comic)(string, error){
	doc, err := GetDocument(comic.Url)
	if err != nil {
		return renderStd(comic, "", "", err.Error()), err
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
	return renderStd(comic, imgUrl, title + " - " + alt, ""), nil
}


func Oglaf(date time.Time, comic Comic)(string, error){
	doc, err := GetDocument(comic.Url)
	if err != nil {
		return renderStd(comic, "", "", err.Error()), err
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
	return renderStd(comic, imgUrl, title + " - " + alt, ""), nil
}


func HagarTheHorrible(date time.Time, comic Comic)(string, error){
	imgSources := GetImagesSrcList(comic.Url)
	for _, imgUrl := range imgSources{
		if strings.Contains(imgUrl, "safr.kingfeatures.com/"){
			return renderStd(comic, imgUrl, "", ""), nil
		}
	}
	errorMsg := "No image url containing \"safr.kingfeatures.com/\" found."
	return renderStd(comic, "#", "", errorMsg), errors.New(errorMsg)
}


func GoComics(date time.Time, comic Comic)(string, error){
	// make right url for passed date
	// (like https://www.gocomics.com/calvinandhobbes/2018/06/16)
	url := comic.Url + date.Format("2006/01/02")
	// get html
	doc, err := GetDocument(url)
	if err != nil {
		return renderStd(comic, "", "", err.Error()), err
	}
	// get picture data
	found := doc.Find("meta[name$=image]")
	if len(found.Nodes) == 0 {
		errorMsg := "Nie znalazłem obrazka. :("
		return renderStd(comic, "", "", errorMsg), errors.New(errorMsg)
	}
	node := found.Nodes[0]
	var imgUrl string
	for i := 0; i < len(node.Attr); i++{
		if node.Attr[i].Key == "content"{
			imgUrl = node.Attr[i].Val
		}
	}
	// render standard template
	return renderStd(comic, imgUrl, "", ""), nil
}


func PHDComic(date time.Time, comic Comic)(string, error){
	// get html
	doc, err := GetDocument(comic.Url)
	if err != nil { return renderStd(comic, "", "", err.Error()), err }
	// get picture data
	imgUrl := ""
	doc.Find("meta").Each(func(i int, s *goquery.Selection){
		_url, _ := s.Attr("content")
		if strings.Contains(_url, "phdcomics.com/comics/archive"){
			imgUrl = _url
		}
	})
	if len(imgUrl) == 0 {
		errorMsg := "Nie znalazłem obrazka. :("
		return renderStd(comic, "", "", errorMsg), errors.New(errorMsg)
	}
	// render standard template
	return renderStd(comic, imgUrl, "", ""), nil
}


func Xkcd(date time.Time, comic Comic)(string, error){
	// get document
	doc, err := GetDocument(comic.Url)
	if err != nil {
		return renderStd(comic, "", "", err.Error()), err
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
	return renderStd(comic, imgUrl, alt + " - " + title, ""), nil
}


func MonkeyUser(date time.Time, comic Comic)(string, error){
    // get document
	doc, err := GetDocument(comic.Url)
	if err != nil {
		return renderStd(comic, "", "", err.Error()), err
	}
	// get node
	found := doc.Find("img")
	node := found.Nodes[3]
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
	// render standard template
	return renderStd(comic, imgUrl, alt + " - " + title, ""), nil
}


func Astronomy_Picture_of_the_Day(date time.Time, comic Comic)(string, error){
	// get document
	doc, err := GetDocument(comic.Url)
	if err != nil { return renderStd(comic, "", "", err.Error()), err }
	// get picture
	found := doc.Find("img")
	pictureNode := found.Nodes[0]
	imgUrl := "http://apod.nasa.gov/apod/"
	for i := 0; i < len(pictureNode.Attr); i++{
		if pictureNode.Attr[i].Key == "src"{
			imgUrl += pictureNode.Attr[i].Val
			break
		}
	}
	// get description
	found = doc.Find("p")
	textNode := goquery.NewDocumentFromNode(found.Nodes[2])
	description, _ := textNode.Html()
	// render standard template
	return renderStd(comic, imgUrl, description, ""), nil
}


func Dilbert(date time.Time, comic Comic)(string, error){
	// get document
	doc, err := GetDocument(comic.Url)
	if err != nil { return renderStd(comic, "", "", err.Error()), err }
	// get node
	found := doc.Find("img")
	node := found.Nodes[2]
	// get data
	var imgUrl string
	for i := 0; i < len(node.Attr); i++{
		if node.Attr[i].Key == "src"{
			imgUrl = node.Attr[i].Val
			break
		}
	}
	imgUrl = "https:" + imgUrl
	// render standard template
	return renderStd(comic, imgUrl, "", ""), nil
}


func DilbertCzech(date time.Time, comic Comic)(string, error){
	// get document
	doc, err := GetDocument(comic.Url)
	if err != nil { return renderStd(comic, "", "", err.Error()), err }
	// get node
	found := doc.Find("img")
	node := found.Nodes[0]
	// get data
	var imgUrl string
	for i := 0; i < len(node.Attr); i++{
		if node.Attr[i].Key == "src"{
			imgUrl = node.Attr[i].Val
			break
		}
	}
	// render standard template
	return renderStd(comic, imgUrl, "", ""), nil
}


func BC(date time.Time, comic Comic)(string, error){
	// build url
	imgUrl := comic.Url + "bcstrips/"
	if int(date.Weekday()) == 0 {
		imgUrl += date.Format("2006/January/bc06_0102sc.jpg")
	} else {
		imgUrl += date.Format("2006/January/bc010206dc.jpg")
	}
	imgUrl = strings.ToLower(imgUrl)
	// render standard template
	return renderStd(comic, imgUrl, "", ""), nil
}


func WizardOfId(date time.Time, comic Comic)(string, error){
	// build url
	imgUrl := comic.Url + "wizardofidstrips/"
	if int(date.Weekday()) == 0 {
		imgUrl += date.Format("2006/January/wiz06_0102sc.jpg")
	} else {
		imgUrl += date.Format("2006/January/wiz010206dc.jpg")
	}
	imgUrl = strings.ToLower(imgUrl)
	// render standard template
	return renderStd(comic, imgUrl, "", ""), nil
}


func Sinfest(date time.Time, comic Comic)(string, error){
	// build url
	imgUrl := "http://www.sinfest.net/btphp/comics/" + date.Format("2006-01-02") + ".gif"
	// render standard template
	return renderStd(comic, imgUrl, "", ""), nil
}
