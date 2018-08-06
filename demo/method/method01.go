package main

import (
	"fmt"
)

// 定义一个 Class Person
type Person struct {
	name string
	age int
	schoole string
}

// 定义一个方法
func (p Person) ShowInfoMechod() {
	fmt.Printf("name is %s, age is %d \n", p.name,p.age)
}

// 定义一个函数
func ShowInfoFun(p Person){
	fmt.Printf("name is %s, age is %d \n", p.name,p.age)
}

func main(){
	p := Person{
		name:"Tinywan",
		age:24,
		schoole:"QingHua",
	}
	// 通过方法访问
	p.ShowInfoMechod();  // name is Tinywan, age is 24

	p1 := Person{
		name:"Tinyaiai",
		age:22,
		schoole:"BeiJing",
	}
	// 通过函数访问
	ShowInfoFun(p1)
}