package main

import "fmt"

func main(){
	// 方式一
	var m map[int]string // 声明一个map
	fmt.Println(m)
	m = map[int]string{} // 初始化一个map
	fmt.Println(m)

	// 方式二
	var m2 map[int]string = map[int]string{}
	fmt.Println(m2)

	// 方式三
	m3 := map[int]string{}
	fmt.Println(m3)

	// 设置、获取、删除
	m3[1] = "Tinywan"
	a := m3[1]
	fmt.Println(m3) // map[1:Tinywan]
	fmt.Println(a)  // Tinywan

	delete(m3,1)  // 删除一个map
	fmt.Println(m3) // map[]
}