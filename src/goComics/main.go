package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
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
	definitions = append(definitions, GetComic("xkcd", "https://xkcd.com/", Xkcd))
	definitions = append(definitions, GetComic("Sinfest", "http://www.sinfest.net/", Sinfest))

	definitions = append(definitions, GetComic("Calvin and Hobbes", "https://www.gocomics.com/calvinandhobbes/", GoComics))
	// gathering data (async)
	date := time.Now()
	var comic Comic
	var wg sync.WaitGroup
	fmt.Println("Starting")
	for i := 0; i < len(definitions); i++ {
		wg.Add(1)
		go func (i int)  {
			defer wg.Done()
			comic = definitions[i]
			definitions[i].HTML = comic.Function(date, comic)
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
	ctx := TemplateContext{date, definitions}
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
