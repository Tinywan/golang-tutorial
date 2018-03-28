package main

import "fmt"

func main() {
	// for i := 1; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf(" 偶数-%d", i)
	// 		continue
	// 	}
	// 	fmt.Printf(" %d", i)
	// }

	i := 0
    for ;i <= 10; { // initialisation and post are omitted
        fmt.Printf("%d ", i)
        i += 2
    }
}
