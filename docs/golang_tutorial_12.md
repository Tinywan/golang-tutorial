11 - 变参函数  
========================

上一节：[第十篇 if else 语句](/docs/golang_tutorial_10.md)   
下一节：[第十二篇 包](/docs/golang_tutorial_12.md)  

这是本Golang系列教程的第11篇。   

## 什么是变参函数？  

变参函数是指可以接受可变数量的参数的函数。  

## 语法？  

如果一个函数的最后一个参数由 `...T` 表示，则表示该函数可以接受任意数量的类型为 T 的参数。

请注意只有函数的最后一个参数才能指定为可变参数。  

## 案例  

你有没有想过为什么 `append `函数可以将任意数量的值追加到切片末尾？这是因为它是一个变参函数。`append` 的原型为 `func append(slice []Type, elems ...Type) []Type`，其中 `elems` 是一个可变参数。 

让我们来创建一个自己的变参函数。我们将编写一个程序来判断某个特定整数是否包含在某个整数列表中。  

```golang
package main

import (  
    "fmt"
)

func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {  
    find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)
}
```

上面的程序中，`func find(num int, nums ...int) `可以接受任意数量的参数。`...int `在内部表示为切片。在这里 `nums` 的类型为 `[]int`。  

第 10 行利用 `range for` 遍历 `nums` 切片，如果找到 num 则打印 num 所在位置，否则打印没有找到。  

上面的程序输出如下：  

```golang
type of nums is []int  
89 found at index 0 in [89 90 95]

type of nums is []int  
45 found at index 2 in [56 67 45 90 109]

type of nums is []int  
78 not found in  [38 56 98]

type of nums is []int  
87 not found in  []  
```

在第 25 行，`find `只有一个参数。我们没有传递任何参数给 `nums ...int`。这是合法的，（译者注：如果没有给可变参数传递任何值，则可变参数为 `nil `切片），在这里 `nums `是一个 `nil `切片，长度和容量都是`0`。  

## 传递切片给可变参数  

我们已经提到 `...T `在内部表示为类型是 `[]T` 切片。如果真是这样，可以传递一个切片给可变参数吗？让我们从下面的例子中寻找答案：  

```golang
package main

import (  
    "fmt"
)

func find(num int, nums ...int) {  
    fmt.Printf("type of nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main() {  
    nums := []int{89, 90, 95}
    find(89, nums)
}
```

在第 23 行，我们没有将若干数量的参数传递给 `find` 的最后一个参数， 而是传递了一个切片。这是非法的，我们不能传递一个切片给可变参数。上面的程序将报错：`main.go:23: cannot use nums (type []int) as type int in argument to find`。  

这里有一个语法糖可以用来将切片传递给变参函数。可以在切片后面加 `...`，这样会将切片展开为其中的各个元素并将它们传递给变参函数。这样该程序将正常工作。  

上面的程序如果将第23行的 `find(89, nums)` 改为` find(89, nums...)`，程序将通过编译，并输出如下：   

```golang
type of nums is []int
89 found at index 0 in [89 90 95]
```

变参函数的介绍到此结束。感谢阅读。  

上一节：[第十一篇 数组和切片](/docs/golang_tutorial_10.md)   
下一节：[第十三篇 Map](/docs/golang_tutorial_12.md)  
