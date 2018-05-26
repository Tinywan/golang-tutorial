package main

import (
	"fmt"
)

func main() {
	personSalary := map[string]int{
		"steve": 1200,
		"jamin": 1500,
	}
	personSalary["Tinywan"] = 20000
	fmt.Println("print All map")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d \n", key, value)
	}
}
