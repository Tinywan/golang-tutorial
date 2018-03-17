package main

import (  
    "fmt"
)

func main() {  
    c1 := complex(5,7)
    c2 := 8 + 27i
    cadd := c1 + c2
    fmt.Println("sum:",cadd)

    cmul := c1 * c2
    fmt.Println("mul :",cmul)
}