package main

import (
	"fmt"
)

// 具体结构体
type Employee struct {
	firstName, lastName string
	age, salary 		int
}

func main() {
	emp1 := Employee{}
	fmt.Println(emp1)

	// 匿名结构体
	emp2 := struct{
		firstName, lastName string
		age, salary			int
	}{
		firstName:"wanshaobo",
		lastName: "wan",
		age: 24,
		salary: 1500,
	}
	fmt.Println(emp2);
}