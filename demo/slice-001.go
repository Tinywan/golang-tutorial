package main

import (  
    "fmt"
)

func main() {  
	arr := [5]int{23,44,55,66,77}
	var b []int = arr[1:4]
	fmt.Println(b)
	c :=[]int{7,8,9,99,12}
	fmt.Println(c)

    darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before",darr)
    for i := range dslice {
        dslice[i]++
    }
	fmt.Println("array after ",darr) 
	
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d \n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	
	var names []string //zero value of a slice is nil
    if names == nil {
        fmt.Println("slice is nil going to append")
        names = append(names, "John", "Sebastian", "Vinay")
        fmt.Println("names contents:",names)
	}
	
	sl1 := []string{"apple","orange","tag"}
	sl2 := []string{"fruit","person"}

	sl3 := append(sl1,sl2...);
	fmt.Println("sl3 : ",sl3)
}