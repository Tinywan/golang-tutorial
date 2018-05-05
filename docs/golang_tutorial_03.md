3 - 变量
========================

上一节：[第二篇  Hello World](/docs/golang_tutorial_02.md)   
下一节：[第四篇  类型](/docs/golang_tutorial_04.md)   

这是本Golang系列教程的第三篇。  

除了亲自写代码以外没有更好的方式来学习一门新的编程语言。在这篇教程中我们将开始编写我们的第一个程序。  

## 什么是变量  

变量（Variable）是给某个内存地址起的一个名字。我们用变量来存储某个特定类型的值。在 Go 中有多种声明变量的语法。  

##  声明单一变量  

声明一个变量的语法为：`var name type`，例如：  

```golang
package main

import "fmt"

func main() {  
    var age int // variable declaration
    fmt.Println("my age is", age)
}
```

语句 `var age int` 声明了一个类型为 int，名称为 age 的变量。在这里我们没有给它赋任何值。如果一个变量没有被赋予任何值，Go 会自动将这个变量初始化为其类型的 0值，比如这里的 age 将被赋值为 0（`译者注：int的0值为0`）。运行这个程序，将得到如下输出：  

```golang
my age is 0
```

一个变量可以被赋予其类型的任何值。例如，在上例中，age 可以被赋予任何整型值:  

```golang
package main

import "fmt"

func main() {  
    var age int // variable declaration
    fmt.Println("my age is ", age)
    age = 24 //assignment
    fmt.Println("my age is", age)
    age = 28 //assignment
    fmt.Println("my new age is", age)
}
```

上面的程序输出如下：  
```golang
my age is  0
my age is 24
my new age is 28
```

## 声明一个带初值的变量  

在声明变量时可以指定其初始值。  

声明一个带初值的变量的语法为：`var name type = initialvalue`，例如：  

```golang
package main

import "fmt"

func main() {  
    var age int = 29 // variable declaration with initial value

    fmt.Println("my age is", age)
}
```

在上面的程序中， age 是一个 int 型的变量，初始值为 29。运行上面的程序，输出如下。可以看到 age 的初始值确实为 29。  

```golang
my age is 29
```

## 类型推导

如果声明一个变量时提供了初始值，Go可以根据该初始值来自动推导变量的类型。因此如果声明变量时提供了初始值，就可以不必指定其类型。  

也就是说，如果声明变量的形式为：`var name = initialvalue`，Go将根据 initialvalue 自动推导变量 name 的类型。  

在下面的例子中，你可以看到声明变量 age 时并没有指定其类型。因为 age 的初值为 29，Go 自动推断其类型为 int。  

```golang
package main

import "fmt"

func main() {  
    var age = 29 // type will be inferred

    fmt.Println("my age is", age)
}
```

## 多变量声明

多个变量可以在一条语句中声明。  

多变量声明的语法为：`var name1, name2 type = initialvalue1, initialvalue2`，例如：  

```golang
package main

import "fmt"

func main() {  
    var width, height int = 100, 50 //declaring multiple variables

    fmt.Println("width is", width, "height is", height)
}
```

如果指定了初始值，**则 type 可以省略**。下面的程序利用类型推导声明了多个变量：  

```golang
package main

import "fmt"

func main() {  
    var width, height = 100, 50 //"int" is dropped

    fmt.Println("width is", width, "height is", height)
}
```

运行上面的程序，可以看到输出结果为 :  

```golang
width is 100 height is 50
```

正如你猜想的那样，如果不指定 width 和 height 的初值，它们将自动被赋值为 0，也就是说它们将以 0 作为初值：  

```golang
package main

import "fmt"

func main() {  
    var width, height int
    fmt.Println("width is", width, "height is", height)
    width = 100
    height = 50
    fmt.Println("new width is", width, "new height is ", height)
}
```

上面的程序将会输出：  

```golang
width is 0 height is 0  
new width is 100 new height is  50 
```

有些时候我们需要在一条语句中**声明多个不同类型的变量**。我们可以使用下面的语法达到此目的：  
```golang
var (  
      name1 = initialvalue1,
      name2 = initialvalue2
)
```

下面的程序就用上面的语法声明了多个不同类型的变量：  

```golang
package main

import "fmt"

func main() {  
    var (
        name   = "naveen"
        age    = 29
        height int
    )
    fmt.Println("my name is", name, ", age is", age, "and height is", height)
}
```

这里我们声明了一个字符串类型的变量 name，以及两个整型的变量 age 和 height。（我们将在下一篇教程中讨论 Golang 中可用的类型）。运行上面的程序将会产生如下输出：  

```golang
my name is naveen , age is 29 and height is 0
```

## 简短声明  

Go 也支持一种声明变量的简洁形式，称为简短声明（Short Hand Declaration），该声明使用了 `:=` 操作符。   

声明变量的简短语法是 `name := initialvalue`。例如：

```golang
package main

import "fmt"

func main() {  
    name, age := "naveen", 29 //short hand declaration

    fmt.Println("my name is", name, "age is", age)
}
```

运行上面的程序，输出如下：  

```golang
my name is naveen age is 29
```

简短声明要求 ` := ` 操作符左边的所有变量都有初始值。下面程序将会抛出错误 `cannot assign 1 values to 2 variables`，这是因为 age 没有被赋值。  

```golang
package main

import "fmt"

func main() {  
    name, age := "naveen" //error

    fmt.Println("my name is", name, "age is", age)
}
```

简短声明的语法要求 `:= `操作符的左边至少有一个变量是尚未声明的。考虑下面的程序：  

```golang
package main

import "fmt"

func main() {
    a, b := 20, 30 // 声明变量a和b
    fmt.Println("a is", a, "b is", b)
    
    b, c := 40, 50 // b已经声明，但c尚未声明
    fmt.Println("b is", b, "c is", c)
    
    b, c = 80, 90 // 给已经声明的变量b和c赋新值
    fmt.Println("changed b is", b, "c is", c)
}
```

上面的程序中，在 `b, c := 40, 50 ` 这一行，虽然变量 b 在之前已经被声明了，但是 c 却是新声明的变量，**因此这是合法的**。该程序的输出为：    

```golang
a is 20 b is 30  
b is 40 c is 50  
changed b is 80 c is 90 
```

但是当我们运行下面的程序  

```golang
package main

import "fmt"

func main() {  
    a, b := 20, 30 //a and b declared
    fmt.Println("a is", a, "b is", b)
    a, b := 40, 50 //error, no new variables
}
```

将会报错：`no new variables on left side of :=` 。这是因为变量 a 和变量 b 都是已经声明过的变量，在` :=` 左侧并没有新的变量被声明。（重要）  

**注意：简短声明只能用在函数中**  

变量可以被赋予运行时产生的值。考虑下面的程序：  

```golang
package main

import (  
    "fmt"
    "math"
)

func main() {  
    a, b := 145.8, 543.8
    c := math.Min(a, b) // 运行时产生的值
    fmt.Println("minimum value is ", c)
}
```

在上面的程序中， c 的值为 a 和 b 的最小值，该值是在运行时计算的。  

一个变量不能被赋予与其类型不同的值。下面的程序将报错：`cannot use "naveen" (type string) as type int in assignment`。这是因为 age 被声明为 int 类型，但是我们试图将 string 类型的值赋给它。  

```golang
package main

func main() {  
    age := 29      // age is int
    age = "naveen" // error since we are trying to assign a string to a variable of type int
}
```

本文由 [GCTT](https://github.com/studygolang/GCTT) 原创翻译，[Go语言中文网](https://studygolang.com/)首发。

