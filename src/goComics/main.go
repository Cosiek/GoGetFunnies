package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"text/template"
	"time"
)


func main() {
	// definitions
	definitions := make([]Comic, 0)
	definitions = append(definitions, GetComic("buttersafe", "http://buttersafe.com/", Buttersafe))
	definitions = append(definitions, GetComic("abstrusegoose", "http://abstrusegoose.com/", Abstrusegoose))
	definitions = append(definitions, GetComic("Hagar the Horrible", "http://hagarthehorrible.com/", HagarTheHorrible))
	definitions = append(definitions, GetComic("xkcd", "https://xkcd.com/", Xkcd))
	definitions = append(definitions, GetComic("Oglaf", "http://www.oglaf.com/", Oglaf))
	definitions = append(definitions, GetComic("Sinfest", "http://www.sinfest.net/", Sinfest))

	definitions = append(definitions, GetComic("Calvin and Hobbes", "https://www.gocomics.com/calvinandhobbes/", GoComics))
	// gathering data (async)
	date := time.Now()
	var comic Comic
	var wg sync.WaitGroup
	var err error
	fmt.Println("Starting")
	for i := 0; i < len(definitions); i++ {
		wg.Add(1)
		go func (i int)  {
			defer func ()  {
				if err := recover(); err != nil { /* just pass */ }
			}()
			defer wg.Done()

			comic = definitions[i]
			definitions[i].HTML, err = comic.Function(date, comic)
			if err != nil{
				fmt.Println(definitions[i].Name + "...Błąd")
			} else {
				fmt.Println(definitions[i].Name + "...OK")
			}
		}(i)
	}
	// wait for results
	wg.Wait()

	// rendering output file
	fmt.Println("Rendering output")
	templ, err := template.ParseFiles("main_template.html")
	if err != nil { panic(err) }
	outFile, err := os.Create("komiksy.html")
	if err != nil { panic(err) }
	ctx := MainTemplateContext{date, definitions}
	err = templ.Execute(outFile, ctx)
	if err != nil { panic(err) }
	outFile.Close()
	// saving an archive file
	os.Mkdir("archive/", os.ModeDir)
	dst := "archive/" + date.Format("2006-01-02") + ".html"
	err = CopyFile("komiksy.html", dst)
	if err != nil { panic(err) }
	// open the browser
	err = exec.Command("xdg-open", "komiksy.html").Start()
	if err != nil { panic(err) }
	fmt.Println("Done")
}
