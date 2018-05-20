package main

import (
	"fmt"
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
	fmt.Println("Done")
}
