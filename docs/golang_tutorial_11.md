11 - 数组和切片  
========================

上一节：[第十篇 if else 语句](/docs/golang_tutorial_10.md)   
下一节：[第十二篇 包](/docs/golang_tutorial_12.md)  

这是本Golang系列教程的第11篇。   

## 数组  

数组是类型相同的元素的集合。例如，整数 5, 8, 9, 79, 76 的集合就构成了一个数组。Go不允许在数组中混合使用不同类型的元素（比如整数和字符串）。  

## 申明  

数组的类型为` n[T]`，其中 `n` 表示数组中元素的个数，`T` 表示数组中元素的类型。元素的个数 `n`也是数组类型的一部分（我们将在稍后详细讨论）。  

有很多声明数组的方式，让我们一个一个地介绍。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    var a [3]int //int array with length 3
    fmt.Println(a)
}
```

`var a [3]int `声明了一个长度为 3 的整型数组。数组中的所有元素都被自动赋值为元素类型的 0 值。比如这里 a 是一个整型数组，因此 a 中的所有元素都被赋值为 0（即整型的 0 值）。运行上面的程序，输出为：`[0 0 0]`。

数组的索引从 0 开始到 `length - 1` 结束。下面让我们给上面的数组赋一些值。

```golang
package main

import (  
    "fmt"
)

func main() {  
    var a [3]int //int array with length 3
    a[0] = 12 // array index starts at 0
    a[1] = 78
    a[2] = 50
    fmt.Println(a)
}
```

`a[0]`表示数组中的第一个元素。程序的输出为：`[12 78 50]`。

（译者注：可以用下标运算符（`[]`）来访问数组中的元素，下标从 0 开始，例如 `a[0]` 表示数组 a 的第一个元素，`a[1] `表示数组 a 的第二元素，以此类推。）  

可以利用**速记声明（shorthand declaration）**的方式来创建同样的数组：  

```golang
package main 

import (  
    "fmt"
)

func main() {  
    a := [3]int{12, 78, 50} // shorthand declaration to create array
    fmt.Println(a)
}
```

上面的程序输出为：`[12 78 50]`。

（译者注：这个例子给出了速记声明的方式：在数组类型后面加一对大括号（`{}`），在大括号里面写元素初始值列表，多个值用逗号分隔。）  

在速记声明中，没有必要为数组中的每一个元素指定初始值。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    a := [3]int{12} 
    fmt.Println(a)
}
```

上面程序的第 8 行：`a := [3]int{12} `声明了一个长度为 3 的数组，但是只提供了一个初值 12。剩下的两个元素被自动赋值为 0。程序 的输出为：`[12 0 0]`。  

在声明数组时你可以忽略数组的长度并用` ... `代替，让编译器为你自动推导数组的长度。比如下面的程序：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    a := [...]int{12, 78, 50} // ... makes the compiler determine the length
    fmt.Println(a)
}
```

上面已经提到，数组的长度是数组类型的一部分。因此 `[5]int` 和 `[25]int`是两个不同类型的数组。正是因为如此，一个数组不能动态改变长度。不要担心这个限制，因为切片（`slices`）可以弥补这个不足。  

```golang
package main

func main() {  
    a := [3]int{5, 78, 8}
    var b [5]int
    b = a //not possible since [3]int and [5]int are distinct types
}
```

在上面程序的第 6 行，我们试图将一个` [3]int `类型的数组赋值给一个 `[5]int `类型的数组，这是不允许的。编译会报错：`main.go:6: cannot use a (type [3]int) as type [5]int in assignment。`  

## 数组是值类型  

在 Go 中数组是值类型而不是引用类型。这意味着当数组变量被赋值时，将会获得原数组（*译者注：也就是等号右面的数组*）的拷贝。新数组中元素的改变不会影响原数组中元素的值。  

```golang
package main

import "fmt"

func main() {  
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a // a copy of a is assigned to b
    b[0] = "Singapore"
    fmt.Println("a is ", a)
    fmt.Println("b is ", b) 
}
```  

上面程序的第 7 行，将数组 a 的拷贝赋值给数组 b。第 8 行，b 的第一个元素被赋值为 `Singapore`。这将不会影响到原数组 a。程序的输出为：  

```golang
a is [USA China India Germany France]  
b is [Singapore China India Germany France]  
```  

同样的，如果将数组作为参数传递给函数，仍然是值传递，在函数中对（作为参数传入的）数组的修改不会造成原数组的改变。  

```golang
package main

import "fmt"

func changeLocal(num [5]int) {  
    num[0] = 55
    fmt.Println("inside function ", num)

}
func main() {  
    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)
}
```  

上面程序的第 13 行，数组 num 是通过值传递的方式传递给函数 `changeLocal` 的，因此该函数执行过程中不会造成 num 的改变。程序输出如下：  

```golang
before passing to function  [5 6 7 8 8]  
inside function  [55 6 7 8 8]  
after passing to function  [5 6 7 8 8]  
```  

## 数组的长度  

希望你喜欢阅读。请留下宝贵的意见和反馈:)  