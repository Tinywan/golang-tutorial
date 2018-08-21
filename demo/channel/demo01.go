package main

import (
	"fmt"
)

func main() {
	var c chan int // 信道 c 的值为 nil
	if c == nil {
		fmt.Println("channel a is nil, going to define it")
		c = make(chan int)
		fmt.Printf("Type of c is %T", c) // 打印通道类型
	}

	fmt.Printf("\n")
	a := make(chan int)
	fmt.Printf("%T", a) // chan int
}
