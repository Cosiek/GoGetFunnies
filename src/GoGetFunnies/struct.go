package main

import "time"

type Comic struct {
  Name string
  Url string
  Function func(date time.Time, comic Comic)(string, error)
  HTML string
  Nsfw bool
}

func GetComic(name string, url string, function func(date time.Time, comic Comic)(string, error)) Comic {
  return Comic{name, url, function, "", false}
}


type MainTemplateContext struct {
	Date time.Time
	Comics []Comic
}


type StdComicTemplateCtx struct {
  Comic Comic
  ImgSrc string
  Title string
  ErrorMsg string
}
