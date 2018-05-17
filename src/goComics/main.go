package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)


func getHTML(url string)[]byte{
	fmt.Println("Ściągam " + url)
	resp, err := http.Get(url)
	if err != nil { panic(err) }
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}


func main() {
	url := "https://gobyexample.com/"
	body := getHTML(url)
	ioutil.WriteFile("output.txt", body, 0644)
}
