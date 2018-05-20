package main

import (
	"fmt"
	"time"
)


func Buttersafe(date time.Time)string{
	url := "http://buttersafe.com/"
	fmt.Printf("buttersafe....")
	imgSources := GetImagesSrcList(url)
	fmt.Println("OK")
	return imgSources[5]
}
