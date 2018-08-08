package main

import (
	"fmt"
)

func modify(sls []int) {  
    sls[0] = 90
}

func main() {  
    a := [3]int{89, 90, 91}
    modify(a[:])
	fmt.Println(a)
	
	b := []int{100,200,300}
	fmt.Println(b[:])
	fmt.Println(b[:1])
}