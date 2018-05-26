package main

import (
	"fmt"
	"sort"
)

func main() {
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

	delete(m3, 1)   // 删除一个map
	fmt.Println(m3) // map[]

	// 复杂map 的操作
	var m5 map[int]map[int]string     // 定义
	m5 = make(map[int]map[int]string) // 通过 make 初始化 最外层的 map

	m5[1] = make(map[int]string) // 针对外层value 的map进行初始化
	m5[1][1] = "OK"
	m_a := m5[1][1]  // 取出map 的值赋予一个变量
	fmt.Println(m_a) // OK

	// 判断一个map 有没有被初始化，使用多返回值判断
	m_b, ok := m5[2][1]
	// 判断是否被初始化操作
	if !ok {
		m5[2] = make(map[int]string)
	}
	m5[2][1] = "OK b"
	m_b, ok = m5[2][1]
	fmt.Println(m_b, ok) // OK b true

	// 迭代操作
	s_map := make([]map[int]string, 5) // 以 map 为元素的slice 使用 make 创建一个切片,元素的slic
	for _, v := range s_map {
		v = make(map[int]string) // v 是值的拷贝
		v[1] = "OK"
		fmt.Println(v)
	}
	fmt.Println(s_map)

	// 针对一个 map 直接操作
	for i := range s_map {
		s_map[i] = make(map[int]string)
		s_map[i][1] = "OK"
		fmt.Println(s_map[i])
	}
	fmt.Println(s_map)

	// map 的间接排序
	// map 集合
	map01 := map[int]string{1: "a", 2: "b", 3: "n", 4: "c", 5: "p", 6: "f"}
	// 切片
	slice01 := make([]int, len(map01))
	i := 0
	for k, _ := range map01 {
		slice01[i] = k
		i++
	}

	fmt.Println(slice01) // 返回的是一个无序的数组:[5 6 1 2 3 4] [3 4 5 6 1 2]
	sort.Ints(slice01)
	fmt.Println(slice01) // 有序的数组:[1 2 3 4 5 6]
}
