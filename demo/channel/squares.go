package main

import (
	"fmt"
)

// 计算 number 每一位的平方和
func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum // 将结果发送给信道 squareop
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 598
	squareop := make(chan int)
	cubeop := make(chan int)
	go calcSquares(number, squareop)
	go calcCubes(number, cubeop)
	squares, cubes := <-squareop, <-cubeop
	fmt.Println("Final output", squares+cubes)
}
