package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)


type TemplateContext struct {
	Date time.Time
	Comics []Comic
}


func main() {
	// definitions
	definitions := make([]Comic, 0)
	definitions = append(definitions, GetComic("buttersafe", "http://buttersafe.com/", Buttersafe))
	definitions = append(definitions, GetComic("Hagar the Horrible", "http://hagarthehorrible.com/", HagarTheHorrible))
	// gathering data
	date := time.Now()
	fmt.Println("Starting")
	for _, comic := range definitions{
		comic.HTML = comic.Function(date)
		fmt.Println(comic.Name, comic.HTML)
	}
	fmt.Println("Rendering output")
	templ, err := template.ParseFiles("main_template.html")
	if err != nil { panic(err) }
	outFile, err := os.Create("komiksy.html")
	if err != nil { panic(err) }
	ctx := TemplateContext{date, definitions}
	err = templ.Execute(outFile, ctx)
	if err != nil { panic(err) }
	outFile.Close()
	fmt.Println("Done")
}
