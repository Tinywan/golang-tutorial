package main

import "fmt"

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

type Person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

// 组合的使用，嵌入结构的使用，被嵌入的结构体相当于把自己所有属性给予了需要嵌入的结构体中
type Hourman struct {
	Sex int
}

type Teacher struct {
	Hourman // 继承 Hourman 所有
	Name    string
	Age     int
}

type Student struct {
	Hourman // 继承 Hourman 所有
	Name    string
	Age     int
}

func main() {
	sr1 := Person{}
	fmt.Println(sr1)

	sr2 := Person{}
	sr2.Name = "Tinywan"
	sr2.Age = 24
	fmt.Println(sr2)

	sr3 := Person{
		Name: "Tinywan Default",
		Age:  25,
	}
	fmt.Println(sr3) // {Tinywan Default 25}
	A(sr3)           // 传递值 {A Tinywan 25}
	fmt.Println(sr3) // {Tinywan Default 25}

	//传递指针
	sr4 := Person{
		Name: "Tinywan Default",
		Age:  25,
	}
	fmt.Println(sr4) // {Tinywan Default 25}
	B(&sr4)          // 传递地址 &{Pointer B 25}
	fmt.Println(sr4) // {Pointer B 25}

	// 一般在结构体初始化的时候进行【取地址符号】，这时候被赋值的变量将会直接指向结构体的【指针】
	// 【推荐改写法】
	sr5 := &Person{ // 这里直接取地址
		Name: "Default",
		Age:  24,
	}
	fmt.Println("Update Before :", sr5)
	sr5.Name = "Update"
	B(sr5)
	C(sr5)
	fmt.Println("Update After :", sr5)

	// 匿名
	sr6 := &struct {
		Name string
		Age  int
	}{
		Name: "nothing",
		Age:  26,
	}
	fmt.Println("%v %p", sr6, sr6)

	//匿名结构使用
	sr7 := Person{Name: "Name01", Age: 24}
	sr7.Contact.Phone = "13669361192"
	sr7.Contact.City = "GanSu"
	fmt.Println(sr7)

	// 赋值操作
	sr8 := Person{Name: "Name008", Age: 26}
	var b Person // 这里必须声明为为Person的类型
	b = sr8
	fmt.Println(sr8) // {Name008 26 { }}
	fmt.Println(b)   // {Name008 26 { }}

	// 嵌入和组合的使用
	te := Teacher{Name: "Student Teacher", Age: 48, Hourman: Hourman{Sex: 1}} //初始化操作
	st := Student{Name: "Wan Student", Age: 24, Hourman: Hourman{Sex: 2}}
	// update
	te.Name = "2018 Name"
	te.Age = 12
	te.Sex = 0
	fmt.Println(te, st) // {{0} 2018 Name 12} {{2} Wan Student 24}

}

// 传值引用
func A(person Person) {
	person.Name = "A Tinywan"
	fmt.Println(person) // {A Tinywan 25}
}

//
func B(person *Person) {
	person.Name = "Pointer B "
	fmt.Println(person)
}

func C(person *Person) {
	person.Name = "Pointer C "
	fmt.Println(person)
}
