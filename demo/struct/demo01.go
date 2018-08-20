package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
	sex  bool
}

func (p *person) isSex() {
	if p.age > 25 {
		p.sex = true
	}

	if p.sex {
		fmt.Println("Who is ", p.name)
	}
}

func main() {
	p1 := person{name: "Tinywan", age: 24}
	p1.isSex()
	p2 := person{name: "Shaobo", age: 26}
	p2.isSex()
	p3 := person{name: "MaLi", age: 32}
	p3.isSex()
}
