package main

import "fmt"

func main() {
	a := [10]int{1,2,3,4,5,6,7,8,9}
	var s1  []int
	fmt.Println(s1) // 输出：[]

	// 只取一个元素
	s2 := a[5]
	fmt.Println(s2) // 输出：6

	// 只取前5个元素
	s3 := a[:5]
	fmt.Println(s3) // 输出：[1 2 3 4 5]

	// 只取后5个元素
	s4 := a[5:]
	fmt.Println(s4) // 输出：[6 7 8 9 0]

	// 截取某一段元素，不包括最后一个
	s5 := a[5:8]
	fmt.Println(s5) // 输出：[6 7 8]

	// 使用make创建一个切片
	s6 := make([]int, 3, 10) //  func make([]T, len, cap) []T 可以用来创建切片
	fmt.Println(len(s6),cap(s6));

	s7 := []byte{'a','b','c','d','e','f','g','h','i','j','k'} // 切片底层对应的数组

	slice_a := s7[2:5]
	fmt.Println(slice_a) //	输出的assica码 值 [99 100 101]
	fmt.Println(string(slice_a)) // 格式化为字符串输出
	fmt.Println(len(slice_a),cap(slice_a))

	slice_b := s7[3:5]
	fmt.Println(string(slice_b)) // 格式化为字符串输出

	// append 函数使用
	s8 := make([]int, 3, 6) // 3个元素，容量为6的切片
	fmt.Printf("%p\n", s8)  // 打印内存地址：0xc042074030
	s8 = append(s8, 12, 48)
	fmt.Printf("%v %p", s8, s8) // 格式化打印值和内存地址：[0 0 0 12 48] 0xc042074030

	// 追加的元素如果没有超多切片容量，则切片的地址是不变的，否则内存地址会变
	s8 = append(s8, 66, 88)
	fmt.Printf("%v %p\n", s8, s8) // [0 0 0 12 48 66 88] 0xc042048060

	// copy 使用
	s9 := []int{1,2,3,4,5,6,7}
	s10 := []int{33,44,55}
	copy(s9,s10) // copy(拷贝，被拷贝)
	fmt.Println(s9) //[33 44 55 4 5 6 7]

	copy(s10,s9) 
	fmt.Println(s10) // [33 44 55]

	copy(s9[2:4],s10[0:2]) 
	fmt.Println(s9) // [1 2 33 44 5 6 7]

	s11 := s7[:]
	fmt.Println(s11) // copy一个数组的所有
}