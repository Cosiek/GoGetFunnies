package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)


func main() {
	// definitions
	definitions := map[string] func(date time.Time)string{
		"buttersafe": Buttersafe,
		"Hagar the Horrible": HagarTheHorrible,
	}
	// gathering data
	date := time.Now()
	fmt.Println("Starting")
	for key, f := range definitions{
		fmt.Println(key, f(date))
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
