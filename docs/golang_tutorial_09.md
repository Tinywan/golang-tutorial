9 - 循环语句  
========================

上一节：[第八篇 if else 语句](/docs/golang_tutorial_08.md)   
下一节：[第十篇 switch 语句](/docs/golang_tutorial_10.md)  

这是本Golang系列教程的第9篇。  

循环语句用于重复执行一段代码。  

for 语句是 Go 中唯一的循环语句。Go 没有提供其他语言（如 C）中的 while 和 do while 语句。  

## for 语句语法  

for 语句的语法如下：  

```golang
for initialisation; condition; post {  
}
```

其中， `initialisation` 为初始化语句，该语句仅执行一次。`initialisation` 语句结束后，接着对 `condition` 求值，如果 `condition` 求值结果为 `true`，则执行大括号` {} `里面的循环体，然后执行 post 语句，如果 `condition` 求值结果为 false 则退出循环。`post` 语句会在每次循环体执行结束后执行。执行完 `post` 语句之后，`condition` 会被重新求值，如果是`true`，则继续执行循环体，否则退出循环。  

## 例子  

下面的程序使用 for 循环打印 1 到 10 之间的整数。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        fmt.Printf(" %d",i)
    }
}
```

在上面的程序中，`i` 被初始化为 `1`。条件语句判断` i `是否小于等于 `10`，如果是则打印 `i`，否则结束循环。`post` 语句在每次迭代结束时将 `i`递增 `1`。直到`i`的值大于 `10` 循环结束。  

上面的程序打印：`1 2 3 4 5 6 7 8 9 10`  

在 `for` 头部定义的变量仅在 `for` 语句范围内可见，因此 `i `不能在 `for` 循环体外被访问。  

## break  

`break` 语句用于终止 `for` 循环，继续执行 `for` 循环后面的语句。  

下面的程序打印 1 到 5 之间的整数。请注意该程序中 `break` 的用法。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        if i > 5 {
            break //loop is terminated if i > 5
        }
        fmt.Printf("%d ", i)
    }
    fmt.Printf("\nline after for loop")
}
```

在上面的程序中，每次迭代都会检查`i`的值，如果`i`大于 5 则执行 `break` 语句跳出循环，并执行 for 循环后面的那条 fmt.Printf 语句。上面的程序输出如下：  

```golang
1 2 3 4 5  
line after for loop  
```

## continue  

`continue` 语句用于跳过 `for` 循环的当前迭代。循环体中`continue`语句之后的所有语句将被跳过不予执行。循环将继续执行下一次迭代。  

让我们写一个程序利用 `continue` 来打印 `1` 到 `10` 之间的奇数。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Printf("%d ", i)
    }
}
```

在上面的程序中，`if i%2 == 0`检测 `i` 除以 `2` 的余数是否为 `0`，如果为 `0` 则 `i` 是偶数，利用 `continue` 语句跳过当前迭代并继续下一次迭代。因此 `continue` 语句之后的 `fmt.Printf` 语句将不被执行，并且循环进入到下一次迭代。上面的程序输出为：`1 3 5 7 9`。  

## 更多例子  

让我们再写一些代码来演示 for 循环的其它变体。  

下面的程序 打印 0 到 10 之间的所有偶数。  

```golang
package main

import (  
    "fmt"
)

func main() {  
    i := 0
    for ;i <= 10; { // initialisation and post are omitted
        fmt.Printf("%d ", i)
        i += 2
    }
}
```

我们已经知道 `for` 循环头部的三个部分`initialisation`，`condition`，`post` 都是可选的。上面的程序中，忽略了 `initialisation` 和 `post` 部分。`i` 在 `for` 循环之外初始化为 `0`，只要`i <= 10` 循环就一直执行，`i` 在循环体内每次递增 `2`。上面的程序输出为：`0 2 4 6 8 10`。  

上面程序中的分号（`;`）也可以省略。这种形式的 `for` 循环可以视为 `while` 循环的替代品。上面的程序可以被重写如下：   

```golang
package main

import (  
    "fmt"
)

func main() {  
    i := 0
    for i <= 10 { //semicolons are ommitted and only condition is present
        fmt.Printf("%d ", i)
        i += 2
    }
}
```

可以在 `for` 循环中声明和操作多个变量，比如下面的程序：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
        fmt.Printf("%d * %d = %d\n", no, i, no*i)
    }

}
```

上面的程序中，`no` 和` i `被声明并初始化为 `10` 和 `1`。它们在每次迭代结束时递增 `1`。在 `condition` 部分使用 `&&` 操作符来确保 `i` 小于或等于 `10` 并且 `no` 小于或等于 `19`。程序的输出如下：  

```golang
10 * 1 = 10  
11 * 2 = 22  
12 * 3 = 36  
13 * 4 = 52  
14 * 5 = 70  
15 * 6 = 90  
16 * 7 = 112  
17 * 8 = 136  
18 * 9 = 162  
19 * 10 = 190 
```

## 无限循环  

可以用下面的语法实现无限循环：  

```golang
for {  
}
```

下面的程序将一直打印` Hello World `永不终止。  

```golang
package main

import "fmt"

func main() {  
    for {
        fmt.Println("Hello World")
    }
}
```

如果你在 [go playground ](https://play.golang.org/p/kYQZw1AWT4)执行上面的程序，你将得到一个错误：`process took too long`。请尝试在本地系统中运行它以无限打印`"Hello World"`。  

还有一个` range for `可用于遍历数组，我们将在介绍数组时介绍它。  

希望你喜欢阅读。请留下宝贵的意见和反馈:)  
