13 - Map 集合  
========================

上一节：[第十篇 if else 语句](/docs/golang_tutorial_10.md)   
下一节：[第十二篇 包](/docs/golang_tutorial_12.md)  

这是本Golang系列教程的第13篇。   

## 什么是 map？  

Map 是 Go 中的内置类型，它将键与值绑定到一起。可以通过键获取相应的值。  

## 如何创建 map？  

可以通过将键和值的类型传递给内置函数 `make` 来创建一个 `map`。语法为：`make(map[KeyType]ValueType)`。（译者注：`map` 的类型表示为 `map[KeyType]ValueType`）例如：   

```golang
personSalary := make(map[string]int) 
```

上面的代码创建了一个名为 `personSalary` 的 map。其中键的类型为 string，值的类型为 int。  

**map 的 0 值为 `nil`。试图给一个 nil map 添加元素给会导致运行时错误。因此 map 必须通过 make 来初始化** （译者注：也可以使用速记声明来创建 map，见下文）。    

```golang
package main

import (  
    "fmt"
)

func main() {  
    var personSalary map[string]int
    if personSalary == nil {
        fmt.Println("map is nil. Going to make one.")
        personSalary = make(map[string]int)
    }
}
```

上面的程序中，`personSalary` 为 `nil`，因此使用 make 初始化它。程序的输出为：`map is nil. Going to make one`.   

## 向 map 中插入元素  

插入元素给 map 的语法与数组相似。下面的代码插入一些新的元素给 `map personSalary`。  
```golang
package main

import (  
    "fmt"
)

func main() {  
    personSalary := make(map[string]int)
    personSalary["steve"] = 12000
    personSalary["jamie"] = 15000
    personSalary["mike"] = 9000
    fmt.Println("personSalary map contents:", personSalary)
}
```

上面的程序输出：`personSalary map contents: map[steve:12000 jamie:15000 mike:9000]`。  

也可以在声明时初始化一个数组：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int {
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("personSalary map contents:", personSalary)
}
```

上面的程序在声明 personSalary 的同时向其中插入了两个元素。接着插入了一个以 "mike" 为键的元素。程序的输出为：  

```golang
personSalary map contents: map[steve:12000 jamie:15000 mike:9000] 
```

`string` 并不是可以作为键的唯一类型，其他所有可以比较的类型，比如，布尔类型，整型，浮点型，复数类型都可以作为键。如果你想了解更多关于可比较类型的话，请参阅：http://golang.org/ref/spec#Comparison_operators  

## 访问 map 中的元素  

现在我们已经添加了一些元素给 map，现在让我们学习如何从 map 中提取它们。根据键获取值的语法为：`map[key]`，例如：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    employee := "jamie"
    fmt.Println("Salary of", employee, "is", personSalary[employee])
}
```

上面的程序非常简单。员工 `jamie` 的工资被取出并打印。程序的输出为：`Salary of jamie is 15000`。  

如果一个键不存在会发生什么？`map` 会返回值类型的 `0 `值。比如如果访问了 `personSalary` 中的不存在的键，那么将返回 `int` 的 0 值，也就是 0。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    employee := "jamie"
    fmt.Println("Salary of", employee, "is", personSalary[employee])
    fmt.Println("Salary of joe is", personSalary["joe"])
}
```

上面的程序输出为:  

```golang
Salary of jamie is 15000  
Salary of joe is 0  
```

上面的程序返回 `joe` 的工资为` 0`。我们没有得到任何运行时错误说明键 joe 在 `personSalary` 中不存在。

我们如何检测一个键是否存在于一个 map 中呢？可以使用下面的语法：

```golang
value, ok := map[key]  
```

上面的语法可以检测一个特定的键是否存在于 map 中。如果 `ok` 是 `true`，则键存在，value 被赋值为对应的值。如果 `ok` 为 `false`，则表示键不存在。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    newEmp := "joe"
    value, ok := personSalary[newEmp]
    if ok == true {
        fmt.Println("Salary of", newEmp, "is", value)
    } else {
        fmt.Println(newEmp,"not found")
    }

}
```

在上面的程序中，第 15 行，`ok` 应该为 `false` ，因为 `joe` 不存在。因此程序的输出为：

```golang
joe not found
```

range for 可用于遍历 map 中所有的元素（译者注：这里 range 操作符会返回 map 的键和值）。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("All items of a map")
    for key, value := range personSalary {
        fmt.Printf("personSalary[%s] = %d\n", key, value)
    }
}
```

上面的程序输出如下：  

```golang
All items of a map  
personSalary[mike] = 9000  
personSalary[steve] = 12000  
personSalary[jamie] = 15000
```

值得注意的是，因为 map 是无序的，因此对于程序的每次执行，不能保证使用 range for 遍历 map 的顺序总是一致的。  

## 删除元素  

`delete(map, key) `用于删除 map 中的 key。delete 函数没有返回值。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("map before deletion", personSalary)
    delete(personSalary, "steve")
    fmt.Println("map after deletion", personSalary)

}
```

上面的程序删除以 `steve` 为键的元素。程序输出为：  

```golang
map before deletion map[steve:12000 jamie:15000 mike:9000]  
map after deletion map[mike:9000 jamie:15000] 
```

## map 的大小  

用内置函数 `len` 获取 map 的大小：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("length is", len(personSalary))

}
```

上面程序中，`len(personSalary) `获取 `personSalary` 的大小。上面的程序输出：`length is 3`。  

##  map 是引用类型  

与切片一样，map 是引用类型。当一个 map 赋值给一个新的变量，它们都指向同一个内部数据结构。因此改变其中一个也会反映到另一个：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("Original person salary", personSalary)
    newPersonSalary := personSalary
    newPersonSalary["mike"] = 18000
    fmt.Println("Person salary changed", personSalary)
}
```

上面的程序中，第 14 行，`personSalary` 赋值给 `newPersonSalary`。下一行，将 `newPersonSalary` 中 `mike` 的工资改为 `18000`。那么在 `personSalary` 中 `mike` 的工资也将变为 `18000`。程序的输出如下：  

```golang
Original person salary map[steve:12000 jamie:15000 mike:9000]  
Person salary changed map[jamie:15000 mike:18000 steve:12000] 
```

将 map 作为参数传递给函数也是一样的。在函数中对 map 的任何修改都会影响在调用函数中看到。  

##  比较 map  

map 不能通过 `== `操作符比较是否相等。`== `操作符只能用来检测 map 是否为 nil。  

```golang
package main

func main() {  
    map1 := map[string]int{
        "one": 1,
        "two": 2,
    }

    map2 := map1

    if map1 == map2 {
    }
}
```

上面的程序将会报错：`invalid operation: map1 == map2 (map can only be compared to nil)`。  

比较两个 map 是否相等的方式是一一比较它们的元素是否相等。我会鼓励你为此编写一个程序，使其工作：）  

我（原文作者）已经将我们讨论的所有概念汇总到一个程序中，你可以从 [github](https://github.com/golangbot/arraysandslices) 下载。  

希望你喜欢阅读。请留下宝贵的意见和反馈:)   

## 以下为扩展知识  

[Go编程基础视频教程笔记](https://study.163.com/course/courseLearn.htm?courseId=306002#/learn/video?lessonId=421019&courseId=306002)  

```golang
package main

import (
	"fmt"
	"sort"
)

func main(){
	// 方式一 
	var m map[int]string // 声明一个map，此时的 map == nil
	fmt.Println(m)
	m = map[int]string{} // 初始化一个map，此时的 map != nil，是map[]
	fmt.Println(m)
	// 以上两种的区别在于有没有被初始化容量  

	// 方式二
	var m2 map[int]string = map[int]string{}
	fmt.Println(m2)

	// 方式三
	m3 := map[int]string{}
	fmt.Println(m3)
	
	// 方式四
	m4 := map[string]string{
        "name":"Tinywan",
        "school":"BAT_UN"
	}
	fmt.Println(m4)
	m5 := make(map[string][string])
	m5["name"] = "Linux"
	m5["school"] = "Unix"
	// 注意：m4和m5两种初始化的方式等价
	
	// 设置、获取、删除
	m3[1] = "Tinywan"
	a := m3[1]
	fmt.Println(m3) // map[1:Tinywan]
	fmt.Println(a)  // Tinywan

	delete(m3,1)  // 删除一个map
	fmt.Println(m3) // map[]

	// 复杂map 的操作
	var m5 map[int]map[int]string // 定义
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
	m_b,ok = m5[2][1]
	fmt.Println(m_b, ok) // OK b true

	// 迭代操作
	s_map := make([]map[int]string,5) // 以 map 为元素的slice 使用 make 创建一个切片,元素的slic
	for _,v := range s_map {
		v = make(map[int]string) // v 是值的拷贝
		v[1] = "OK"
		fmt.Println(v);
	}
	fmt.Println(s_map)

	// 针对一个 map 直接操作
	for i := range s_map {
		s_map[i] = make(map[int]string) 
		s_map[i][1] = "OK"
		fmt.Println(s_map[i]);
	}
	fmt.Println(s_map)

	// map 的间接排序
	// map 集合
	map01 := map[int]string{1:"a", 2:"b", 3:"n", 4:"c", 5:"p", 6:"f"}
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
```