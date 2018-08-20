package main

import (
	"fmt"
)

// 打印字符串中的字节
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

// 打印字符串中的字符
func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
}

func main() {
	name := "Hello World"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)

	fmt.Printf("\n")

	// 特殊字符
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)

	// 打印汉字
	fmt.Printf("\n")
	name = "萬少波"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
}
