5 - 常量
========================

上一节：[第四篇  类型](/docs/golang_tutorial_04.md)   
下一节：[第六篇  函数](/docs/golang_tutorial_06.md)   

这是本Golang系列教程的第五篇。  

## 定义常量  

常量（constant）表示固定的值，比如：`5，-89，"I love Go"，67.89` 等等。  

考虑如下程序：  

```golang
var a int = 50  
var b string = "I love Go"  
```

上面的程序中，` a `和` b` 分别被赋值为常量 `50` 和` "I love Go"`。关键字 const 用于指示常量，如 `50` 和 `"I love Go"`。在上面的代码中，尽管没有使用关键字 `const` 修饰 `50` 与` "I love Go"`，但它们在Go的内部表示为常量。   

关键字 const 修饰的名字为常量，不能被重新赋予任何值。 因此下面的程序会报错：`cannot assign to a`。  

```golang
package main

func main() {  
    const a = 55 //allowed
    a = 89 //reassignment not allowed
}
```

常量的值必须在编译期间确定。因此不能将函数的返回值赋给常量，因为函数调用发生在运行期。  
```golang
package main

import (  
    "fmt"
    "math"
)

func main() {  
    fmt.Println("Hello, playground")
    var a = math.Sqrt(4)//allowed
    const b = math.Sqrt(4)//not allowed
}
```

在上面的程序中，`a `是一个变量因此可以被赋值为函数` math.Sqrt(4)` 的调用结果。（我们将在单独的教程中更详细的讨论函数）。而 `b `是一个常量，它的值必须在编译期间确定，但是函数 `math.Sqrt(4) `的调用结果只能在运行时被计算出来，因此在编译 `const b = math.Sqrt(4) `时将会报错：  

```golang
error main.go:11: const initializer math.Sqrt(4) is not a constant.
```

## 字符串常量  

字符串常量最简单的常量，通过了解字符串常量可以更好的理解常量的概念。  

在Go中任何用双引号括起来的值都是字符串常量，比如 `"Hello World"，"Sam" `都是字符串常量。  

字符串常量是什么类型呢？答案是 无类型**untyped**。  

**像 "Hello World" 这样的字符串没有任何类型。**  

```golang
const hello = "Hello World"  
```

上面的代码将 "Hello World" 赋给一个名为 hello 的常量。那么现在常量** hello** 是不是就有了类型？答案是：No。hello 仍然没有类型。 

Go是一种强类型语言。 所有变量都需要显式类型。 变量 `name` 被赋值为一个无类型的常量 `"Sam"`，它是如何工作的呢？  

```golang
package main

import (  
    "fmt"
)

func main() {  
    var name = "Sam"
    fmt.Printf("type %T value %v", name, name)

}
```

答案是无类型常量有一个默认的类型，当且仅当代码中需要无类型常量提供类型时，它才会提供该默认类型。在语句 `var name = "Sam" `中，`name`需要一个类型，因此它从常量 "Sam" 中获取其默认类型：`string`。  

<u>（译者注：这里可以理解为常量是无类型的，但是在需要类型的场合，编译器会根据常量的值和上下文将常量转换为相应的类型。）</u>

有没有办法创建一个有类型（typed）的常量？答案是：Yes。下面的代码创建了一个有类型常量。  
```golang
const typedhello string = "Hello World" 
```

上面的代码中， `typedhello` 是一个字符串类型的常量。  

Go是强类型语言。在赋值时混合使用类型是不允许的。让我们通过以下代码说明这是什么意思。  

```golang
package main

func main() {  
    var defaultName = "Sam" //allowed
    type myString string
    var customName myString = "Sam" //allowed
    customName = defaultName //not allowed
}
```

在上面的代码中，我们首先创建了一个变量 `defaultName` 并且赋值为常量 `"Sam"`。常量 `"Sam"` 的默认类型为 `string`，因此赋值之后，`defaultName` 的类型亦为`string`。  

下一行我们创建了一个新的类型 `myString`，它是 `string` 的别名。  

<u>（译者注：可以通过 type NewType Type 的语法来创建一个新的类型。类似 C 语言的 typedef 。）</u>  

接着我们创建了一个名为 `customName` 的 `myString` 类型的变量，并将常量 `"Sam"` 赋给它。因为常量 `"Sam"` 是无类型的所以可以将它赋值给任何字符串变量。因此这个赋值是合法的，`customName` 的类型是 `myString`。  

现在我们有两个变量：string 类型的 `defaultName` 和 `myString` 类型的 `customName`。尽管我们知道 `myString` 是 string 的一个别名，但是Go的强类型机制不允许将一个类型的变量赋值给另一个类型的变量。<u>**因此，customName = defaultName 这个赋值是不允许的，编译器会报错：main.go:10: cannot use defaultName (type string) as type myString in assignment  **</u>   

## 布尔常量  

布尔常量与字符串常量（在概念上）没有区别。布尔常量只包含两个值：`true` 和 `false`。字符串常量的规则也同样适用于布尔常量，这里不再赘述。下面的代码解释了布尔常量：  

```golang
package main

func main() {  
    const trueConst = true
    type myBool bool
    var defaultBool = trueConst //allowed
    var customBool myBool = trueConst //allowed
    defaultBool = customBool //not allowed
}
```

上面的程序很好理解，这里就不过多解释了。  

## 数值常量  

数值常量（Numeric constants）包括整数，浮点数以及复数常量。数值常量有一些微妙之处。  

让我们看一些例子使事情变得明朗。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    fmt.Println("Hello, playground")
    const a = 5
    var intVar int = a
    var int32Var int32 = a
    var float64Var float64 = a
    var complex64Var complex64 = a
    fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)
}
```

上面的程序中，`const a `是无类型的，值为` 5`。你也许想知道` a` 的默认类型是什么？如果它有默认类型，那么它是怎么赋值给其他类型的变量的？ 答案在于使用 `a` 时的上下文。我们暂时放下这两个问题，先来看下面的程序：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    fmt.Println("Hello, playground")
    var i = 5
    var f = 5.6
    var c = 5 + 6i
    fmt.Printf("i's type %T, f's type %T, c's type %T", i, f, c)
}
```

在上面的程序中，所有变量的类型都是由数值常量决定的。在语法上看，`5` 在是一个整数，`5.6 `是一个浮点数，` 5 + 6i `是一个复数。运行上面的程序，输出为：`i's type int, f's type float64, c's type complex128`  

<u>（译者注：编译器可以根据字面值常量的表现形式来确定它的默认类型，比如 5.6 表现为浮点数，因此它的默认类型为 float64 ，而 "Sam" 表现为字符串，因此它的默认类型为 stirng。）</u>  

现在应该很清楚下面的程序是如何工作的了：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    fmt.Println("Hello, playground")
    const a = 5
    var intVar int = a
    var int32Var int32 = a
    var float64Var float64 = a
    var complex64Var complex64 = a
    fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)
}
```

在这个程序中，`a` 的值是 `5` 并且 `a` 在语法上是泛化的（它既可以表示浮点数 `5.0`，也可以表示整数 `5`，甚至可以表示没有虚部的复数 `5 + 0i`），因此 `a` 可以赋值给任何与之类型兼容的变量。像 `a` 这种数值常量的默认类型可以想象成是通过上下文动态生成的。`var intVar int = a` 要求 `a` 是一个 `int`，那么 `a` 就变成一个` int` 常量。`var complex64Var complex64 = a` 要求 `a` 是一个 `complex64`，那么 `a` 就变成一个 `complex64` 常量。这很容易理解。   

以上程序打印结果：

```golang
Hello, playground
intVar 5
int32Var 5
float64Var 5
complex64Var (5+0i)
```

## 数值表达式  

数值常量可以在表达式中自由的混合和匹配，仅当将它们赋值给变量或者在代码中明确需要类型的时候，才需要他们的类型。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    var a = 5.9/8
    fmt.Printf("a's type %T value %v",a, a)
}
```

在上面的程序中，语法上 `5.9` 是一个浮点数，`8` 是一个整数。因为它们都是数值常量，因此 `5.9/8` 是合法的。相除的结果为 `0.7375`，是一个浮点数。因此变量 `a` 的类型为浮点型。  

上面的程序输出为：     

```golang
a's type float64 value 0.7375
```

本文由 [GCTT](https://github.com/studygolang/GCTT) 原创翻译，[Go语言中文网](https://studygolang.com/)首发。
