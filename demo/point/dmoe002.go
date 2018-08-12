package main

import (
	"fmt"
	"reflect"
)

// import (
// 	"fmt"
// ) go_basic_courses

func main() {
	// 声明一个集合字典，默认为空
	a := 15
	fmt.Println("address of b is ", a)
	fmt.Println("address of b is ", &a)
	m := map[string]int{}
	fmt.Println("address of b i is ", m)
	fmt.Printf("Type of a is %T\n", m)
	fmt.Printf("m值是 %v \n", m == nil)
	var dir map[string]string
	fmt.Printf("dir字典的值是:   | %v \n", dir)
	fmt.Printf("dir字典是否为nil | %v \n", dir == nil)
	// nil map不能赋值,如果直接赋值会报错：“panic: assignment to entry in nil map”，下面语句将会报错
	// dir["name"] = "Tinywan"

	fmt.Println("----------------------------")
	// 使用make函数进行初始化创建一个非nil的map
	dir = make(map[string]string)
	// map是引用类型，未初始化的是指向nil，初始化了以后应该就有自己的内存空间了，所以不是nil
	fmt.Printf("dir字典的值是:   | %v \n", dir)
	fmt.Printf("dir字典是否为nil | %v \n", dir == nil)

	fmt.Println("----------------------------")
	// make之后分配内存了,一旦分配了内存地址就不为空了,可以赋值了
	dir["name"] = "Tinywan"
	fmt.Printf("dir字典的键:值 |%v \n", dir)
	fmt.Printf("dir字典的类型  |%v \n", reflect.TypeOf(dir))
	fmt.Printf("name的值是     |%v \n", dir["name"])
	// value, ok := s["ok"]
	// if ok == nil {
	// 	s := make(map[string]int)
	// 	s["ok"] = 19
	// 	fmt.Println(s["ok"])
	// }
	// fmt.Println(s["ok"])
}
