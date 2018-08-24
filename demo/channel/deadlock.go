package main

func main() {
	m := make(chan int)
	m <- 5
}
