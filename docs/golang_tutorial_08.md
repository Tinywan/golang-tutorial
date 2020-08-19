8 - if else 语句  
========================

上一节：[第七篇  包](/docs/golang_tutorial_07.md)   
下一节：[第九篇  循环语句](/docs/golang_tutorial_09.md)  

这是本Golang系列教程的第8篇。  

if 是一个条件语句。if 语句的语法为：  

```golang
if condition {  
}
```

如果 condition 为 true，那么就执行 { 和 } 之间的代码。与其它语言（如C）不同，即使 {} 之间只有一条语句，{} 也是必需的。  

if 语句后面可以接可选的 else if 和 else 语句：  

```golang
if condition {  
} else if condition {
} else {
}
```

if 后面可以接任意数量的` else if `语句。`condition` 的求值由上到下依次进行，直到某个 `if `或者 `else if `中的 `condition` 为 `true` 时，执行相应的代码块。如果没有一个 `conditon` 为 `true`，则执行 `else` 中的代码块。  

让我们写一个简单的程序来判断一个数是奇数还是偶数：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    num := 10
    if num % 2 == 0 { //checks if number is even
        fmt.Println("the number is even") 
    }  else {
        fmt.Println("the number is odd")
    }
}
```

`if num % 2 == 0 `语句检测一个数除以 2 的余数是否为 0，如果是则输出：`"the number is even"`，否则输出：`"the number is odd"`。以上程序输出：`the number is even`。  

if 语句还有如下的变体。这种形式的 if 语句先执行 `statement`，然后再判断 `conditon` 。  

```golang
if statement; condition {  
}
```

让我们用这种形式的 if 改写上面的程序：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    if num := 10; num % 2 == 0 { //checks if number is even
        fmt.Println(num,"is even") 
    }  else {
        fmt.Println(num,"is odd")
    }
}
```

在上面的程序中， `num` 在 `if `语句中初始化。需要注意的一点是，`num` 只能在 if 和 else 里面进行访问，即 `num` 的范围仅限于 `if else` 块中。如果我们试图在 `if` 或 `else` 之外访问 `num`，编译器将报错。  

让我们用 else if 再写一个程序：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    num := 99
    if num <= 50 {
        fmt.Println("number is less than or equal to 50")
    } else if num >= 51 && num <= 100 {
        fmt.Println("number is between 51 and 100")
    } else {
        fmt.Println("number is greater than 100")
    }
}
```

在上面的程序 `else if num >= 51 && num <= 100` 为 `true`，因此程序将输出：`number is between 51 and 100`  

## 小陷阱

`else` 语句应该在 `if` 语句的大括号 `}` 之后的同一行中开始，否则编译器会报错。  

让我们通过一个程序来理解这一点。

```golang
package main

import (  
    "fmt"
)

func main() {  
    num := 10
    if num % 2 == 0 { //checks if number is even
        fmt.Println("the number is even") 
    }  
    else {
        fmt.Println("the number is odd")
    }
}
```

在上面的程序中，`else` 语句没有在 `if` 语句 `}` 结束后的同一行中开始，而是从下一行开始的。Go中不允许这样做。如果你运行这个程序，编译器会输出错误：  

```golang
main.go:12:5: syntax error: unexpected else, expecting }  
```

***这是由于Go会自动插入分号，点击链接了解分号插入规则：https://golang.org/ref/spec#Semicolons***  

在分号插入规则中，如果 `}` 是当前一行的最后一个字符，那么 `}` 后面会被插入一个分号。所以在上面的例子中，if语句的 `}` 之后会被自动插入一个分号。

所以实际上我们的程序变成了：  

```golang
if num%2 == 0 {  
      fmt.Println("the number is even") 
};  //semicolon inserted by Go
else {  
      fmt.Println("the number is odd")
}
```

你可以看到第三行中插入了一个分号。

由于 `if{...} else {...}` 是一个单独的语句，分号不应该出现在语句的中间。因此，`else` 需要放在 `}` 之后的同一行中。  

我重新改写了程序，将 `else` 移动到if语句结束之后的 `}` 的后面，以防止自动分号插入。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    num := 10
    if num%2 == 0 { //checks if number is even
        fmt.Println("the number is even") 
    } else {
        fmt.Println("the number is odd")
    }
}
```

现在编译器就可以正常编译我们的程序了。

希望你喜欢阅读。请留下宝贵的意见和反馈:)  
