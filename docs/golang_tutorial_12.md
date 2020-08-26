12 - 变参函数  
========================

上一节：[第十一篇 数组和切片](/docs/golang_tutorial_11.md)   
下一节：[第十三篇 Map](/docs/golang_tutorial_13.md)  

这是本Golang系列教程的第12篇。   

## 什么是变参函数？  

变参函数是指可以接收可变数量的参数的函数。  

如果一个函数的最后一个参数由 `...T` 表示，则表示该函数可以接收任意数量的类型为 T 的参数。

请注意只有函数的最后一个参数才能指定为可变参数。

## 语法

```golang
func hello(a int, b ...int){
}
```

上述函数中，`b` 就是一个可变参数，因为它的类型定义有 `...` 前缀，因而可以接收任意数量的输入。以下语句可以用来调用这个变参函数：

```golang
hello(1, 2) //passing one argument "2" to b  
hello(5, 6, 7, 8, 9) //passing arguments "6, 7, 8 and 9" to b 
```

第一行中，我们将 `2` 赋给了 `b` ；第二行中，我们将四个参数 `6, 7, 8, 9` 赋给了 `b`。

我们也可以选择不将参数传入变参函数（即：传入“零”个参数），例如：

```golang
hello(1)
```

我们调用了 `hello` 函数，但是没有为 `b` 传入任何参数，这也是完全可以的。

为什么变参函数中的可变参数必须在定义参数时放在最后？作为反例，我们来尝试将 `hello` 函数的第一个参数变成可变参数：

```golang
func hello(b ...int, a int){
}
```

以上函数中，我们将无法为 `a` 传入参数， 因为无论我们在函数中传入多少参数，这些参数都会被赋给可变参数 `b`。上述函数会出现编译错误：`syntax error: cannot use ... with non-final parameter b`，因此可变参数必须在定义参数时被放在最后。

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

上面的程序中，`func find(num int, nums ...int)` 可以为 `nums` 接收任意数量的参数。在函数内部，`nums` 的类型为 `[]int`，即整数切片类型。

**变参函数会将可变数量的参数转换成相应类型的切片，例如上述函数的第22行，传入 `find` 函数的可变参数是 89，90，95，随后这些参数会被编译器转换为整数切片 `[]int{89, 90, 95}` 然后被送入 `find` 函数内。**

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

在第 25 行，`find `只有一个参数。我们没有传递任何参数给 `nums ...int`。这是合法的，（译者注：如果没有给可变参数传递任何值，则可变参数为 `nil` 切片），在这里 `nums` 是一个 `nil` 切片，长度和容量都是`0`。  

## 切片 VS 可变参数

我们现在肯定会想到一个问题：既然传入的可变参数被转换成了切片，那么我们为何不直接将一个切片传入函数中呢？我们将之前的函数用切片参数来改写：

```golang
package main

import (  
    "fmt"
)

func find(num int, nums []int) {  
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
    find(89, []int{89, 90, 95})
    find(45, []int{56, 67, 45, 90, 109})
    find(78, []int{38, 56, 98})
    find(87, []int{})
}
```

用可变参数代替切片的优势如下：

1. 不用在每次调用函数时都创建一个新的切片。上述函数中，我们在第22、23、24、25行都创建了新的切片，如果使用变参函数，这些额外的切片创建就可以省略。

2. 在第25行，为了满足 `find` 函数的参数定义，我们创建了一个空的切片，在变参函数中，我们只需要写 `find(87)` 即可。

3. 我（原文作者）认为带变参函数的程序比带切片的函数可读性强 :)

## 将切片传入可变参数  

我们已经提到 `...T` 在内部表示为类型是 `[]T` 切片。如果真是这样，可以传递一个切片给可变参数吗？让我们从下面的例子中寻找答案：  

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

在第 23 行，我们没有将若干数量的参数传递给 `find` 函数的最后一个参数， 而是传递了一个切片。这是非法的，我们不能传递一个切片给可变参数。上面的程序将报错：`main.go:23: cannot use nums (type []int) as type int in argument to find`。  

**有一个语法糖可以用来将切片传递给变参函数：在切片后面加 `...`，这样会将切片展开为其中的各个元素并将它们传递给变参函数，程序将正常工作。**

上面的程序如果将第23行的 `find(89, nums)` 改为` find(89, nums...)`，程序将通过编译，并输出如下：   

```golang
type of nums is []int
89 found at index 0 in [89 90 95]
```

变参函数的介绍到此结束。感谢阅读。  
