package main

import (
	"fmt"
	"math"
)

// 长方形
type Rectangle struct {
	length int
	width int
}

// 圆
type Circle struct {
	radius float64
}

// 长方形面积
func (r Rectangle) Area() int {
	return r.length * r.width
}

// 圆面积
func (c Circle) Area() float64  {
	return math.Pi * c.radius * c.radius
}

func main(){
	r := Rectangle{
		length:20,
		width:40,
	}
	fmt.Printf(" Area of Rectangle is %d \n", r.Area())

	c := Circle{
		radius:10,
	}
	fmt.Printf(" Area of Circle is %f \n", c.Area())
}