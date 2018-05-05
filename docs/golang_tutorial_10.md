10 - switch 语句  
========================

上一节：[第九篇 循环语句](/docs/golang_tutorial_09.md)   
下一节：[第十一篇 数组和切片](/docs/golang_tutorial_11.md)  

这是本Golang系列教程的第10篇。  

switch 是一个条件语句，用于将一个表达式的求值结果与可能的值的列表进行匹配，并根据匹配结果执行相应的代码。可以认为 `switch` 语句是编写多个 `if-else `子句的替代方式。  

举例是说明问题最好的方式，让我们写一个简单的程序，输入手指编号，输出对应的手指称：）。例如 0 表示拇指，1 表示食指等。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    finger := 4
    switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")
    }
}
```

在上面的程序中，case 语句依次（从上到下）求值并与 `finger` 进行匹配，直到找到第一个与 `finger` 匹配的 case，并执行其中的代码。在这里 `finger` 的值为 4，因此打印 `Ring`。  

多个有相同值的 case 是不允许的。如果你运行下面的程序，编译将会报错：`duplicate case 4 in switch`。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    finger := 4
    switch finger {
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 4://duplicate case
        fmt.Println("Another Ring")
    case 5:
        fmt.Println("Pinky")
    }
}
```

## default case  

我们每只手只有 5 根手指，但是如果我们输入一个错误的手指序号会发生什么呢？这里就要用到 `default` 语句了。当没有其他 case 匹配时，将执行 default 语句。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    switch finger := 8; finger {//finger is declared in switch
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")
    default: //default case
        fmt.Println("incorrect finger number")
    }
}
```

在上面的程序中，`finger` 的值为 8，它不匹配任何 case，因此打印 `incorrect finger number`。default 语句不必放在 switch 语句的最后，而可以放在 switch 语句的任何位置。  

你也许发现了另外一个小的改变，就是将 finger 声明在了 switch 语句中。switch 语句可以包含一个可选的语句，该语句在表达式求值之前执行。在` switch finger := 8`; `finger` 这一行中， `finger` 首先被声明，然后作为表达式被求值。这种方式（译者注：在 switch 语句中声明变量的方式）声明的 finger 只能在 switch 语句中访问。  

## 包含多个表达式的 case  

可以在一个 case 中包含多个表达式，每个表达式用逗号分隔。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    letter := "i"
    switch letter {
    case "a", "e", "i", "o", "u": //multiple expressions in case
        fmt.Println("vowel")
    default:
        fmt.Println("not a vowel")
    }
}
```

上面的程序检测 letter 是否是元音。`case "a", "e", "i", "o", "u":` 这一行匹配所有的元音。程序的输出为：`vowel`。  

## 没有表达式的 switch  

switch 中的表达式是可选的，可以省略。如果省略表达式，则相当于` switch true`，这种情况下会将对每一个 case 的表达式的求值结果与 true 做比较，如果相等，则执行相应的代码。   

```golang
package main

import (  
    "fmt"
)

func main() {  
    num := 75
    switch { // expression is omitted
    case num >= 0 && num <= 50:
        fmt.Println("num is greater than 0 and less than 50")
    case num >= 51 && num <= 100:
        fmt.Println("num is greater than 51 and less than 100")
    case num >= 101:
        fmt.Println("num is greater than 100")
    }
}
```

在上面的程序中，switch 后面没有表达式因此被认为是` switch true `并对每一个 case 表达式的求值结果与 true 做比较。`case num >= 51 && num <= 100:` 的求值结果为 `true`，因此程序输出：`num is greater than 51 and less than 100`。这种类型的 switch 语句可以替代多重 if else 子句。   

## fallthrough  

在 Go 中执行完一个 case 之后会立即退出 switch 语句。`fallthrough`语句用于标明执行完当前 case 语句之后按顺序执行下一个case 语句。  

让我们写一个程序来了解 `fallthrough`。下面的程序检测 `number` 是否小于 50、100 或 200。列如：如果我们输入75，程序将打印 75 小于 100 和 200，这是通过 `fallthrough` 语句实现的。   

```golang
package main

import (  
    "fmt"
)

func number() int {  
        num := 15 * 5 
        return num
}

func main() {

    switch num := number(); { //num is not a constant
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is lesser than 200", num)
    }
}
```

switch 与 case 中的表达式不必是常量，他们也可以在运行时被求值。在上面的程序中 num 初始化为函数 number() 的返回值。程序首先对 switch 中的表达式求值，然后依次对每一个case 中的表达式求值并与 true 做匹配。匹配到 `case num < 100:` 时结果是 true，因此程序打印：`75 is lesser than 100`，接着程序遇到 `fallthrough` 语句，因此继续对下一个 case 中的表达式求值并与 true 做匹配，结果仍然是 true，因此打印：`75 is lesser than 200`。最后的输出如下：   

```golang
75 is lesser than 100  
75 is lesser than 200 
```

`fallthrough` 必须是 case 语句块中的最后一条语句。如果它出现在语句块的中间，编译器将会报错：`fallthrough statement out of place`。  

还有一种 switch 语句叫做 `type switch`，我们将在学习接口时介绍它。  

希望你喜欢阅读。请留下宝贵的意见和反馈:)  