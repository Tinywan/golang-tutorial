4 - 类型
========================

上一节：[第三篇  变量](/docs/golang_tutorial_03.md)   
下一节：[第五篇  常量](/docs/golang_tutorial_05.md)   

这是本Golang系列教程的第四篇。  

下面是 Go 支持的基本类型  

* bool  
* Numeric Types  
  * int8, int16, int32, int64, int  
  * uint8, uint16, uint32, uint64, uint  
  * float32, float64  
  * complex64, complex128  
  * byte  
  * rune  
* string

## bool  

bool 类型表示一个布尔值，值为 true 或者 false。  

```golang
package main

import "fmt"

func main() {  
    a := true
    b := false
    fmt.Println("a:", a, "b:", b)
    c := a && b
    fmt.Println("c:", c)
    d := a || b
    fmt.Println("d:", d)
}
```

在上面的程序中，a 赋值为 true，b 赋值为 false。  

c 赋值为 a && b。仅当 a 和 b 都为 true 时，操作符 && 才返回 true。因此，在这里 c 为 false。

当 a 或者 b 为 true 时，操作符 || 返回 true。在这里，由于 a 为 true，因此 d 也为 true。

我们将会得到如下输出：  

```golang
a: true b: false  
c: false  
d: true 
```

## 有符号整型  

int8：表示 8 位有符号整型  
大小：8 位  
范围：`-128～127`  

int16：表示 16 位有符号整型  
大小：16 位  
范围：`-32768～32767`  

int32：表示 32 位有符号整型  
大小：32 位  
范围：`-2147483648～2147483647`  

int64：表示 64 位有符号整型  
大小：64 位  
范围：`-9223372036854775808～9223372036854775807`  

int：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型。除非对整型的大小有特定的需求，否则你通常应该使用 int 表示整型。  
大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。  
范围：在 32 位系统下是 `-2147483648～2147483647`，而在 64 位系统是 `-9223372036854775808～9223372036854775807`。  

```golang
package main

import "fmt"

func main() {  
    var a int = 89
    b := 95  // 简短
    fmt.Println("value of a is", a, "and b is", b)
}
```

如下输出：  

```golang
value of a is 89 and b is 95
```

在上述程序中，a 是 int 类型，而 b 的类型通过赋值（95）推断得出。上面我们提到，int 类型的大小在 32 位系统下是 32 位，而在 64 位系统下是 64 位。接下来我们会证实这种说法。  

在 Printf 方法中，使用 `%T` 格式说明符（Format Specifier），可以打印出变量的**类型**。Go 的 unsafe 包提供了一个 Sizeof 函数，该函数接收变量并返回它的字节大小。unsafe 包应该小心使用，因为使用 unsafe 包可能会带来可移植性问题。不过出于本教程的目的，我们是可以使用的。  

下面程序会输出变量 a 和 b 的类型和大小。格式说明符 %T 用于打印类型，而 %d 用于打印字节大小。  

```golang
package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
    var a int = 89
    b := 95
    fmt.Println("value of a is", a, "and b is", b)
    fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) //type and size of a
    fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b
}
```

以上程序会输出：  

```golang
value of a is 89 and b is 95  
type of a is int, size of a is 4  
type of b is int, size of b is 4  
```

从上面的输出，我们可以推断出 a 和 b 为 int 类型，且大小都是 32 位（4 字节）。如果你在 64 位系统上运行上面的代码，会有不同的输出。在 64 位系统下，a 和 b 会占用 64 位（8 字节）的大小。  

以下为windows 64位系统运行结果（非翻译案例）：  

```golang
value of a is 89 and b is 95
type of a is int, size of a is 8
type of b is int, size of b is 8 
```

## 无符号整型  

uint8：表示 8 位无符号整型  
大小：8 位  
范围：0～255  

uint16：表示 16 位无符号整型  
大小：16 位  
范围：0～65535  

uint32：表示 32 位无符号整型  
大小：32 位  
范围：0～4294967295  

uint64：表示 64 位无符号整型  
大小：64 位  
范围：0～18446744073709551615  

uint：根据不同的底层平台，表示 32 或 64 位无符号整型。  
大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。  
范围：在 32 位系统下是 0～4294967295，而在 64 位系统是 0～18446744073709551615。  

##  浮点型  

float32：32 位浮点数  
float64：64 位浮点数  

下面一个简单程序演示了整型和浮点型的运用。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    a, b := 5.67, 8.97
    fmt.Printf("type of a %T b %T\n", a, b)
    sum := a + b
    diff := a - b
    fmt.Println("sum", sum, "diff", diff)

    no1, no2 := 56, 89
    fmt.Println("sum", no1+no2, "diff", no1-no2)
}
```

a 和 b 的类型根据赋值推断得出。在这里，a 和 b 的类型为 float64（float64 是浮点数的默认类型）。我们把 a 和 b 的和赋值给变量 sum，把 b 和 a 的差赋值给 diff，接下来打印 sum 和 diff。no1 和 no2 也进行了相同的计算。上述程序将会输出：  

```golang
type of a float64 b float64
sum 14.64 diff -3.3000000000000007
sum 145 diff -33
```

## 复数类型  

complex64：实部和虚部都是 float32 类型的的复数。  
complex128：实部和虚部都是 float64 类型的的复数。  

内建函数 complex 用于创建一个包含实部和虚部的复数。complex 函数的定义如下：  

```golang
func complex(r, i FloatType) ComplexType
```

该函数的参数分别是实部和虚部，并返回一个复数类型。实部和虚部应该是相同类型，也就是 float32 或 float64。如果实部和虚部都是 float32 类型，则函数会返回一个 complex64 类型的复数。如果实部和虚部都是 float64 类型，则函数会返回一个 complex128 类型的复数。  

还可以使用简短语法来创建复数：  

```golang
c := 6 + 7i
```

下面我们编写一个简单的程序来理解复数：

```golang
package main

import (  
    "fmt"
)

func main() {  
    c1 := complex(5,7)
    c2 := 8 + 27i
    cadd := c1 + c2
    fmt.Println("sum:",cadd)

    cmul := c1 * c2
    fmt.Println("mul :",cmul)
}
```

在上面的程序里，c1 和 c2 是两个复数。c1的实部为 5，虚部为 7。c2 的实部为8，虚部为 27。c1 和 c2 的和赋值给 cadd ，而 c1 和 c2 的乘积赋值给 cmul。该程序将输出：  

```golang
sum: (13+34i)
mul : (-149+191i)
```

##  其他数字类型  

byte 是 uint8 的别名。  
rune 是 int32 的别名。  

在学习字符串的时候，我们会详细讨论 byte 和 rune。  

## string 类型  

在 Golang 中，字符串是字节的集合。如果你现在还不理解这个定义，也没有关系。我们可以暂且认为一个字符串就是由很多字符组成的。我们后面会在一个教程中深入学习字符串。 下面编写一个使用字符串的程序。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    first := "Wan"
    last := "ShaoBo"
    name := first +" "+ last
    fmt.Println("My name is",name)
}
```

上面程序中，first 赋值为字符串 "Naveen"，last 赋值为字符串 "Ramanathan"。+ 操作符可以用于拼接字符串。我们拼接了 first、空格和 last，并将其赋值给 name。  

上述程序将打印输出:

```golang
My name is Wan ShaoBo
```

还有许多应用于字符串上面的操作，我们将会在一个单独的教程里看见它们。  

## 类型转换  

Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换。我们通过一个例子说明这意味着什么。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    i := 55      //int
    j := 67.8    //float64
    sum := i + j //不允许 int + float64
    fmt.Println(sum)
}
```

上面的代码在 C 语言中是完全合法的，然而在 Go 中，却是行不通的。i 的类型是 int ，而 j 的类型是 float64 ，我们正试图把两个不同类型的数相加，Go 不允许这样的操作。如果运行程序，你会得到   

```golang
# command-line-arguments
demo04\demo04-07.go:10:14: invalid operation: i + j (mismatched types int and float64)
```

要修复这个错误，i 和 j 应该是相同的类型。在这里，我们把 j 转换为 int 类型。把 v 转换为 T 类型的语法是 T(v)。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    i := 55      //int
    j := 67.8    //float64
    sum := i + int(j) //j is converted to int
    fmt.Println(sum)
}
```

现在，当你运行上面的程序时，会看见输出

```golang
122
```

赋值的情况也是如此。把一个变量赋值给另一个不同类型的变量，需要显式的类型转换。下面程序说明了这一点。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    i := 10
    var j float64 = float64(i) // 若没有显式转换，该语句会报错
    fmt.Println("j", j)
}
```

在第 9 行，i 转换为 float64 类型，接下来赋值给 j。如果不进行类型转换，当你试图把 i 赋值给 j 时，编译器会抛出错误。  

本文由 [GCTT](https://github.com/studygolang/GCTT) 原创翻译，[Go语言中文网](https://studygolang.com/)首发。













