package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"text/template"
	"time"
)


func main() {
	// definitions
	definitions := make([]Comic, 0)

	definitions = append(definitions, GetComic("Hagar the Horrible", "http://hagarthehorrible.com/", HagarTheHorrible))
	definitions = append(definitions, GetComic("xkcd", "https://xkcd.com/", Xkcd))
	definitions = append(definitions, GetComic("abstrusegoose", "http://abstrusegoose.com/", Abstrusegoose))
	definitions = append(definitions, GetComic("PHD", "http://phdcomics.com/comics.php", PHDComic))
	definitions = append(definitions, GetComic("APotD", "http://apod.nasa.gov/apod/", Astronomy_Picture_of_the_Day))
	definitions = append(definitions, GetComic("buttersafe", "http://buttersafe.com/", Buttersafe))
	definitions = append(definitions, GetComic("The System", "http://www.systemcomic.com/", TheSystem))
	definitions = append(definitions, GetComic("Sinfest", "http://www.sinfest.net/", Sinfest))
	definitions = append(definitions, GetComic("Oglaf", "http://www.oglaf.com/", Oglaf))
	definitions[len(definitions) - 1].Nsfw = true
	definitions = append(definitions, GetComic("Dilbert Czeski", "https://ekonomika.idnes.cz/dilbert.aspx", DilbertCzech))
	definitions = append(definitions, GetComic("Dilbert", "http://dilbert.com/", Dilbert))
	definitions = append(definitions, GetComic("B.C.", "https://johnhartstudios.com/", BC))
	definitions = append(definitions, GetComic("Wizard of ID", "https://johnhartstudios.com/", WizardOfId))
	definitions = append(definitions, GetComic("The Barn", "https://www.gocomics.com/thebarn/", GoComics))
	definitions = append(definitions, GetComic("Garfield", "https://www.gocomics.com/garfield/", GoComics))
	definitions = append(definitions, GetComic("Barney & Clyde", "http://www.gocomics.com/barneyandclyde/", GoComics))
	definitions = append(definitions, GetComic("Pickles", "https://www.gocomics.com/pickles/", GoComics))
	definitions = append(definitions, GetComic("Luann", "http://www.gocomics.com/luann/", GoComics))
	definitions = append(definitions, GetComic("The Argyle Sweater", "http://www.gocomics.com/theargylesweater/", GoComics))
	definitions = append(definitions, GetComic("Liberty Meadows", "http://www.gocomics.com/libertymeadows/", GoComics))
	definitions = append(definitions, GetComic("Pearls Before Swine", "http://www.gocomics.com/pearlsbeforeswine/", GoComics))
	definitions = append(definitions, GetComic("Dark Side of the Horse", "http://www.gocomics.com/darksideofthehorse/", GoComics))
	definitions = append(definitions, GetComic("Nonsequitur", "http://www.gocomics.com/nonsequitur/", GoComics))
	definitions = append(definitions, GetComic("Texts from Mittens", "https://www.gocomics.com/texts-from-mittens/", GoComics))
	definitions = append(definitions, GetComic("Calvin and Hobbes", "https://www.gocomics.com/calvinandhobbes/", GoComics))

	var err error
	// prepare dir
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil { panic(err) }
	targetDirName := dir + "/komiksy/"
	os.Mkdir(targetDirName, os.ModeDir)
	// prepare logger
	logFile, err := os.OpenFile(targetDirName + "log.txt",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil { panic(err) }
	defer logFile.Close()
	log.SetOutput(logFile)
	// gathering data (async)
	date := time.Now()
	var comic Comic
	var wg sync.WaitGroup
	fmt.Println("Starting")
	for i := 0; i < len(definitions); i++ {
		wg.Add(1)
		go func (i int)  {
			defer func ()  {
				err := recover()
				if err != nil {
					log.Println(definitions[i].Name, " - ", err)
					fmt.Println(definitions[i].Name + "...Błąd")
				}
				wg.Done()
			}()

			comic = definitions[i]
			definitions[i].HTML, err = comic.Function(date, comic)
			if err != nil{
				log.Println(definitions[i].Name, " - ", err)
				fmt.Println(definitions[i].Name + "...Błąd")
			} else {
				fmt.Println(definitions[i].Name + "...OK")
			}
		}(i)
	}
	// wait for results
	wg.Wait()
	// write css
	writeCssFile(targetDirName)
	// rendering output file
	fmt.Println("Rendering output")
	templ, err := template.New("main").Parse(MAIN_TEMPLATE)
	if err != nil { panic(err) }
	targetFileName := targetDirName + "komiksy.html"
	outFile, err := os.Create(targetFileName)
	if err != nil { panic(err) }
	defer outFile.Close()
	ctx := MainTemplateContext{date, definitions}
	err = templ.Execute(outFile, ctx)
	if err != nil { panic(err) }
	// saving an archive file
	dst := targetDirName + date.Format("2006-01-02") + ".html"
	err = CopyFile(targetFileName, dst)
	if err != nil { panic(err) }
	// open the browser
	err = exec.Command("xdg-open", targetFileName).Start()
	if err != nil { panic(err) }
	fmt.Println("Done")
}
