package main

import "time"

type Comic struct {
  Name string
  Url string
  Function func(date time.Time, comic Comic)string
  HTML string
}

func GetComic(name string, url string, function func(date time.Time, comic Comic)string) Comic {
  return Comic{name, url, function, ""}
}


type MainTemplateContext struct {
	Date time.Time
	Comics []Comic
}


type StdComicTemplateCtx struct {
  Comic Comic
  ImgSrc string
  Title string
}
