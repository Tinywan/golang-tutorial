package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("Name", "Tinywan")
	fmt.Println("Name = ", os.Getenv("Name"))
	fmt.Println("Bar", os.Getenv("bar"))
	fmt.Println()

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
go