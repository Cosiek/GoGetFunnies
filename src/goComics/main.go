package main

import (
	"fmt"
	"time"
)


func main() {
	date := time.Now()
	fmt.Println("Starting")
	fmt.Println(Buttersafe(date))
	fmt.Println("Done")
}
