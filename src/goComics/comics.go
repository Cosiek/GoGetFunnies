package main

import (
	"fmt"
	"strings"
	"time"
)


func Buttersafe(date time.Time)string{
	url := "http://buttersafe.com/"
	fmt.Printf("buttersafe....")
	imgSources := GetImagesSrcList(url)
	fmt.Println("OK")
	return imgSources[5]
}


func HagarTheHorrible(date time.Time)string{
	url := "http://hagarthehorrible.com/"
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
