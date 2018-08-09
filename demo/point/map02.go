package main

import (
	"fmt"
)

func main(){
	// [1] 实例化一个 集合
	var personSalary = make(map[string]int)
	personSalary["Tinywan"] = 10000
	personSalary["Tinyaiai"] = 50000
	personSalary["VCKL"] = 200
	fmt.Println("personSalary map contents:",personSalary)

	// [2] 直接初始化
	animalSalary := map[string]int {
		"Pig":100,
		"Dog":200,
	}

	animalSalary["Fire"] = 400
	fmt.Println("animalSalary map contents:",animalSalary)

	employee := "Tinywan"
	fmt.Println("Salary of", employee, "is", personSalary[employee])
	fmt.Println("Salary of joe is", personSalary["joe"])

	// 检测一个特定的键是否存在于 map 中
	newEmp := "Tinyaiai"
	value,ok := personSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is ",value)
	}else{
		fmt.Println(newEmp,"not found")
	}
	
	fmt.Println("----------------遍历 map 中所有的元素---------------------")
	for key,value := range personSalary {
		fmt.Printf("personSalary[%s] = %d \n",key,value)
	}
}