
15 - 指针  
========================

上一节：[第十四篇 字符串](/docs/golang_tutorial_14.md)   
下一节：[第十六篇 结构体](/docs/golang_tutorial_16.md)  

这是本Golang系列教程的第15篇。我们将学习 Go 语言中指针的作用，并且将 Go 指针与其他语言（如 C/C++）的指针进行区分。

## 什么是指针？

指针是存储一个变量的内存地址的变量。  

![pointer-explained](../images/pointer-explained.png)  

在上图中，变量 `b` 的值是 `156`，存储在地址为 `0x1040a124` 的内存中。变量 `a` 存储了变量 `b` 的地址，也可以说 `a` 指向 `b`。  

## 指针的声明  

指向类型 `T` 的指针用 `*T` 表示。  

让我们写一些代码。

```golang
package main

import (  
    "fmt"
)

func main() {  
    b := 255
    var a *int = &b
    fmt.Printf("Type of a is %T\n", a)
    fmt.Println("address of b is", a)
}
```

`&` 操作符用来获取一个变量的地址。在上面的程序中，第 9 行我们将 `b` 的地址赋给 `a`（`a` 的类型为 `*int`）。现在我们说 `a` 指向了 `b`。当我们打印 `a` 的值时，`b` 的地址将会被打印出来。程序的输出为：

```
Type of a is *int
address of b is 0x1040a124
```

你可能得到的是一个不同的 `b` 的地址，因为 `b` 可能被存储在内存的任意一个地方。  

## 指针的空值  

指针的空值为 `nil`。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    a := 25
    var b *int
    if b == nil {
        fmt.Println("b is", b)
        b = &a
        fmt.Println("b after initialization is", b)
    }
}
```

'b' 首先被初始化为 `nil`，之后被赋值为 `a` 的地址，程序输出为：

```golang
b is <nil>  
b after initialization is 0x1040a124
```

## 利用 `new` 函数创建指针

Go 还提供一个实用的函数 `new` 来创建指针。`new` 函数接收一个类型，并返回一个该类型的空值指针。

我们来看一个实例：

```golang
package main

import (  
    "fmt"
)

func main() {  
    size := new(int)
    fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
    *size = 85
    fmt.Println("New size value is", *size)
}
```

上面的程序中，我们在第 8 行利用 `new` 函数创建了一个指向 `int` 类型的指针。由于 `int` 类型的空值为 `0`，变量 `size` 是一个指向 `0` 的整数指针（`*int` 类型）。

程序输出：

```golang
Size value is 0, type is *int, address is 0x414020  
New size value is 85
```

## 指针的间接引用

间接引用（dereference）指针就是获取指针所指向的变量的值。间接引用 `a` 的语句是 `*a`。例如：

```golang
package main  
import (  
    "fmt"
)

func main() {  
    b := 255
    a := &b
    fmt.Println("address of b is", a)
    fmt.Println("value of b is", *a)
}
```

在第 10 行，我们间接引用了 `a` 并打印了它的值，也就是 `b` 的值。程序输出为：

```golang
address of b is 0x1040a124  
value of b is 255  
```

我们再编写一个利用指针改变变量 `b` 的值的程序：

```golang
package main

import (  
    "fmt"
)

func main() {  
    b := 255
    a := &b
    fmt.Println("address of b is", a)
    fmt.Println("value of b is", *a)
    *a++
    fmt.Println("new value of b is", b)
}
```

在第 12 行，我们将 `a` 所指向的值加上了 1，也就是改变了 `b` 的值（因为 `a` 指向 `b`）。因此 `b` 变成了256，程序输出：

```golang
address of b is 0x1040a124  
value of b is 255  
new value of b is 256  
```

