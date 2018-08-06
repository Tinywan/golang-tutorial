package main

import (
	"fmt"
)

type Square struct {
	width int
}

func area(s *Square) {
	fmt.Printf("Area Function result: %d\n", (s.width * s.width))
}

func (s * Square) area() {
	fmt.Printf("Area Method result: %d\n", (s.width * s.width))
}

func main(){
	r := Square{
		width:20,
	}
	p := &r // poniter
	area(p)
	p.area()

	//area(r)
	r.area()
}