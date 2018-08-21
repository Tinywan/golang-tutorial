package main

import (
	"fmt"
)

func producer(num chan int) {
	for i := 0; i < 10; i++ {
		num <- i
	}
	close(num)
}

func main() {
	ch := make(chan int)
	go producer(ch)

	// for {
	// 	v, ok := <-ch
	// 	if ok == false { // 如果 ok 是 false 表示信道已经被关闭
	// 		break // 通过 break 退出循环
	// 	}
	// 	fmt.Println("Received", v, ok)
	// }

	for v := range ch {
		fmt.Println("Received ", v)
	}
}
