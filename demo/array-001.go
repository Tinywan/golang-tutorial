package main

import "fmt"

func main() {
	arr := [...]float64{23.2,34.12,45.22,55.6}
	fmt.Println("length of a is  ",len(arr)); // print 4

	for i := 0; i < len(arr); i++ { //looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, arr[i])
	}
	
	sum := float64(0);
	for i,v := range arr {
		fmt.Printf("%d the element of a is %.2f\n ",i,v);
		sum += v
	}
	fmt.Println("\n of all element of a ",sum);
}