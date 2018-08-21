package main

import ()
import "fmt"

func sendData(d chan<- int) {
	d <- 5
}

func main() {
	d := make(chan<- int)
	go sendData(d)
	//fmt.Println(<-d) //试图从一个只写信道中接收数据

	m := make(chan int) // 将双向信道转换为只写或只读信道
	go sendData(m)
	fmt.Println(<-m)
}
