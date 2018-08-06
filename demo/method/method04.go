package main

import (
	"fmt"
)

type rectangle struct {
	length int
	width int
}

func Area(r rectangle)  {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) Area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func main(){
	r := rectangle{
		length:10,
		width:20,
	}
	Area(r)
	r.Area()

	p := &r
	//Area(p)
	p.Area()
}