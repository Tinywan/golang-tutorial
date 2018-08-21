package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("This is hello goroutines")
	done <- true
}

func getName(name chan string) {
	fmt.Println("This is getName goroutines")
	name <- "Tinywan"
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done

	name := make(chan string)
	go getName(name)
	fmt.Println(<-name)

	fmt.Println("This is main function")
}
