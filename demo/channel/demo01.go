package main

import (
	"fmt"
)

func main() {
	var m map[string]string
	if m == nil {
		fmt.Println("this is nil map")
	}
	m = make(map[string]string)
	m["name"] = "Tinywan"
	fmt.Println(m)

	m1 := map[string]int{}
	fmt.Println(m1)
	m1["age"] = 24
	m1["dateTime"] = 20180909
	fmt.Println(m1)
}
