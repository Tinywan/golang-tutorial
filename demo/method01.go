package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	schoole string
}

// mechtod
func (p Person) ShowName() {
	fmt.Printf("name is %s, age is %s", p.name,p.age);
}

func main(){

}