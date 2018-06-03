package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)


func main() {
	// definitions
	definitions := make([]Comic, 0)
	definitions = append(definitions, GetComic("buttersafe", "http://buttersafe.com/", Buttersafe))
	definitions = append(definitions, GetComic("Hagar the Horrible", "http://hagarthehorrible.com/", HagarTheHorrible))
	// gathering data
	date := time.Now()
	fmt.Println("Starting")
	for _, comic := range definitions{
		fmt.Println(comic.Name)
	}
	fmt.Println("Rendering output")
	templ, err := template.ParseFiles("main_template.html")
	if err != nil { panic(err) }
	outFile, err := os.Create("komiksy.htm")
	if err != nil { panic(err) }
	err = templ.Execute(outFile, date)
	if err != nil { panic(err) }
	outFile.Close()
	fmt.Println("Done")
}
