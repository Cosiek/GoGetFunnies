package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)


func main() {
	url := "https://gobyexample.com/"
	resp, err := http.Get(url)
	if err != nil { panic(err) }
	defer resp.Body.Close()

	fmt.Println("Witaj świecie!")
	body, err := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile("output.txt", body, 0644)
}
