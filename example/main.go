package main

import (
	"fmt"
	bar "processBar"
	"time"
)

func main() {

	bar := bar.NewBar(100, time.Microsecond*200, "#")
	go bar.Run()
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Done <- i
	}
	time.Sleep(time.Second * 1)
	fmt.Printf("every thing ready!\n")
}
