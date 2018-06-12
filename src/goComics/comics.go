package main

import (
	"fmt"
	"strings"
	"time"
)


func Buttersafe(date time.Time, comic Comic)string{
	url := comic.Url
	fmt.Printf("buttersafe....")
	imgSources := GetImagesSrcList(url)
	fmt.Println("OK")
	return imgSources[5]
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
