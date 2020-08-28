
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

