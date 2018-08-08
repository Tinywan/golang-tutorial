package main

import (
	"fmt"
)

func main(){
	// map[KeyType]ValueType  集合类型
	var personSalary map[string]int 
	// go的三个引用类型 map slice chan
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}
}