package main

import "time"

type Comic struct {
  Name string
  Url string
  Function func(date time.Time)string
  HTML string
}

func GetComic(name string, url string, function func(date time.Time)string) Comic {
  return Comic{name, url, function, ""}
}
