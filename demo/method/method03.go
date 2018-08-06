package main

import (
	"fmt"
)

type Employee struct{
	name string
	age int
}

// 传值
func (e Employee) changeName(name string)  {
	e.name = name
}

// 传指针
func (e *Employee) changeAge(age int)  {
	e.age = age
}

func main()  {
	e := Employee{
        name: "Mark Tinywan",
        age:  50,
	}
	
	fmt.Printf("before name is %s \n", e.name)
	e.changeName("Tinyaiai")
	fmt.Printf("after name is %s \n", e.name)

	fmt.Printf("before name is %d \n", e.age)
	(&e).changeAge(24) // 用 e.changeAge(36) 也是可以的
	fmt.Printf("after age is %d \n", e.age)

}