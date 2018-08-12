package main

import (
	"fmt"
)

func main() {
	b := 255
	var a = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	s := map[string]int{"age": 24}
	var t map[string]string // 正确
	t1 := map[string]string // 错误
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(t1)
}
