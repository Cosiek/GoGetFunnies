package main

import "time"

type Comic struct {
  Name string
  Url string
  Function func(date time.Time)string
}
