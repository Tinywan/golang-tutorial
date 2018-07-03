16 - 结构体  
========================

上一节：[第十篇 if else 语句](/docs/golang_tutorial_10.md)   
下一节：[第十二篇 包](/docs/golang_tutorial_12.md)  

这是本Golang系列教程的第16篇。   

## 什么是结构体？  

结构体（struct）是用户自定义的类型，它代表若干字段的集合。有些时候将多个数据看做一个整体要比单独使用这些数据更有意义，这种情况下就适合使用结构体。  

比如将一个员工的 firstName, lastName 和 age 三个属性打包在一起成为一个 `employee` 结构就是很有意义的。  

## 结构体的声明  

```golang
type Employee struct {  
    firstName string
    lastName  string
    age       int
}
```

上面的代码片段声明了一个名为 `Employee` 的结构体类型，它拥有 firstName，lastName 和 age 三个字段。同一类型的多个字段可以合并到一行（用逗号分隔），并将类型放在后面。上面的结构体中 firstName 与 lastName 都是 string 类型，因此可以将它们写在一起。   

```golang
type Employee struct {  
    firstName, lastName string
    age, salary         int
}
```

上面的结构体 `Employee` 是一个**具名结构体`（named structure）`**，因为它创建了一个具有名字的结构体类型： `Employee` 。我们可以定义具名结构体类型的变量。  

我们也可以定义一个没有类型名称的结构体，这种结构体叫做**匿名结构体（anonymous structures）**。  

```golang
var employee struct {  
        firstName, lastName string
        age int
}
```  

上面的代码片段声明了一个匿名结构体变量 employee。  

## 定义具名结构体变量  

下面的程序说明了如何定义一个具名结构体 Employee 的变量。  

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {

    //creating structure using field names
    emp1 := Employee{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Anderson",
    }

    //creating structure without using field names
    emp2 := Employee{"Thomas", "Paul", 29, 800}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)
}
```  

在上面的程序中，我们定义了一个名为 `Employee` 的结构体类型。我们可以通过指定字段的名称和对应的值来创建一个结构体变量，比如在第15行，我们就是通过这种方式定义了 `Employee` 类型的一个结构体变量 `empl`。这里字段名称的顺序没必要和声明结构体类型时的一致。例如这里我们将 `lastName` 放在了最后，程序同样正确运行。  

在定义结构体变量时也可以忽略字段名称，例如在第 23 行，我们定义 emp2 时没有指定字段名称。但是通过这种方式定义的结构体变量时，字段值的顺序必须与声明结构体类型时字段的顺序保持一致。  

上面的程序输出如下：  

```golang
Employee 1 {Sam Anderson 25 500}  
Employee 2 {Thomas Paul 29 800} 
```  

## 定义匿名结构体变量  

```golang
package main

import (  
    "fmt"
)

func main() {  
    emp3 := struct {
        firstName, lastName string
        age, salary         int
    }{
        firstName: "Andreah",
        lastName:  "Nikola",
        age:       31,
        salary:    5000,
    }

    fmt.Println("Employee 3", emp3)
}
```  

在上面的程序中，第3行定义了一个 匿名结构体变量 emp3。正如我们提到的一样，这种结构体成为匿名结构体，因为它只创建了一个新的结构体变量 emp3，而没有定义新的结构体类型。

程序的输出为：  

```golang
Employee 3 {Andreah Nikola 31 5000}  
```  

## 结构体变量的 0 值  

当定义一个结构体变量，但是没有给它提供初始值，则对应的字段被赋予它们各自类型的0值。  

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    var emp4 Employee //zero valued structure
    fmt.Println("Employee 4", emp4)
}
```  

上面的程序定义了 emp4 但是没有赋予任何初值。因此 firstName 和 lastName 被赋值为 string 类型的0值，也就是空字符串。age 和 salary 被赋值为 int 类型的0值，也就是 0 。程序的输出为：  

```golang
Employee 4 {  0 0} 
```  

可以指定一些字段而忽略一些字段。在这种情况下，被忽略的字段被赋予相应类型的 0 值。  

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp5 := Employee{
        firstName: "John",
        lastName:  "Paul",
    }
    fmt.Println("Employee 5", emp5)
}
```  

在上面的程序中，第14和15行，firstName 和 lastName 被提供了初始值，而 age 和 salary 没有。因此 age 和 salary 被指定为0值。程序的输出为：  

```golang
Employee 5 {John Paul 0 0} 
```  

## 访问结构体中的字段  

使用点` . `操作符来访问结构体中的字段。  

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp6 := Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", emp6.firstName)
    fmt.Println("Last Name:", emp6.lastName)
    fmt.Println("Age:", emp6.age)
    fmt.Printf("Salary: $%d", emp6.salary)
}
```  

在上面的程序中，通过` emp6.firstName `访问 emp6 中的字段 firstName。程序的输出为：  

```golang
First Name: Sam  
Last Name: Anderson  
Age: 55  
Salary: $6000
```  

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

## 知识扩展  

[Go编程基础视频教程笔记](https://study.163.com/course/courseLearn.htm?courseId=306002#/learn/video?lessonId=421019&courseId=306002)  

```golang
package main

import (
	"fmt"
	"sort"
)

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

希望你喜欢阅读。请留下宝贵的意见和反馈:)  