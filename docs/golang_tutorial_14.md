
14 - 字符串  
========================

上一节：[第十篇 if else 语句](/docs/golang_tutorial_10.md)   
下一节：[第十二篇 包](/docs/golang_tutorial_12.md)  

这是本Golang系列教程的第14篇。  

`string` 类型单独提取为一篇教程是因为在 Go 中，`string` 的实现方式同其他语言的不同。  

## 访问字符串中的字节  

因为字符串是字节数组，因此可以访问一个字符串中的字节。  

```golang
package main

import (
	"fmt"
)

// 打印字符串中的字节
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func main() {
	name := "Hello World"
	printBytes(name)
}
```
在上面的程序中，`len(s)` 返回字符串中的字节数，我们用一个 `for` 循环以 16 进制打印这些字节。`%x` 格式化指示符用来以 16 进制打印参数。上面的程序打印：`48 65 6c 6c 6f 20 57 6f 72 6c 64`。它们是 `"Hello World"` 以`UTF-8`方式编码的`Unicode`值。对 `Unicode` 字符集和 `UTF-8` 编码有一个基本的了解会更好的理解 `string` 类型。我（原文作者）建议大家阅读：https://naveenr.net/unicode-character-set-and-utf-8-utf-16-utf-32-encoding/ 来学习什么是 `Unicode` 和 `UTF-8`。  

让我们修改上面的程序以打印字符串中的字符：  

```golang
package main

import (  
    "fmt"
)

func printBytes(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}


func printChars(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%c ",s[i])
    }
}

func main() {  
    name := "Hello World"
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
}
```
在第 16 行的 `printChars` 函数中，`%c` 格式化指示符用来打印字符串中的字符。上面的程序输出为：

```golang
48 65 6c 6c 6f 20 57 6f 72 6c 64
Hello World
```

虽然上面的程序看起来是一种合法的打印字符串中各个字符的方法，但是这里有一个严重的错误。让我们深入这段代码看看究竟是哪里不对。  

```golang
package main

import (  
    "fmt"
)

func printBytes(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%c ",s[i])
    }
}

func main() {  
    name := "Hello World"
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
    fmt.Printf("\n")
    name = "Señor"
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
}
```
上面程序的输出为：  
```golang
48 65 6c 6c 6f 20 57 6f 72 6c 64
Hello World
53 65 c3 b1 6f 72
SeÃ±or
```

## rune  

`rune` 是 Go 中的内置类型，它是 `int32` 的别名。在 Go 中，`rune` 表示一个 `Unicode` 码点。无论一个码点会被编码为多少个字节，它都可以表示为一个 `rune`。让我们修改上面的程序，使用 `rune` 来打印字符串中的字符。  

```golang
package main

import (  
    "fmt"
)

func printBytes(s string) {  
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {  
    runes := []rune(s)
    for i:= 0; i < len(runes); i++ {
        fmt.Printf("%c ",runes[i])
    }
}

func main() {  
    name := "Hello World"
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
    fmt.Printf("\n\n")
    name = "Señor"
    printBytes(name)
    fmt.Printf("\n")
    printChars(name)
}
```

在上面的程序中，第 14 行，字符串被转换为 `tune` 切片。然后我们遍历该切片并打印其中的字符。程序的输出如下：  

```golang
48 65 6c 6c 6f 20 57 6f 72 6c 64
Hello World
53 65 c3 b1 6f 72
Señor
```
上面的输出是正确的。这正是我们想要的结果。  

## 使用 range for 遍历字符串  

上面的程序是遍历字符串中字符的一个正确方式。但是 Go 提供了一种更简单的方式来做到这一点：使用 `range` `for`。  

```golang
package main

import (  
    "fmt"
)

func printCharsAndBytes(s string) {  
    for index, rune := range s {
        fmt.Printf("%c starts at byte %d\n", rune, index)
    }
}

func main() {  
    name := "Señor"
    printCharsAndBytes(name)
}
```

在上面的程序中，第 8 行通过使用 `range for` 遍历字符串。`range` 返回一个 `rune` （在 `byte` 数组中）的位置，以及 `rune` 本身。上面的程序输出为：  

```golang
S starts at byte 0
e starts at byte 1
ñ starts at byte 2
o starts at byte 4
r starts at byte 5
```

从上面的输出可以看到，ñ 占两个字节：）

## 通过 `byte` 切片创建字符串  

```golang
package main

import (  
    "fmt"
)

func main() {  
    byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice)
    fmt.Println(str)
}
```
在上面的程序中，`byteSlice` 是 `"Café"` 经过 `UTF-8` 编码后得到的切片（用 16 进制表示） 。上面的程序输出为：`Café`。  

如果我们换成对应的十进制数程序会正常工作吗？答案是：`Yes`。让我们测试一下：  

```golang
package main

import (  
    "fmt"
)

func main() {  
    byteSlice := []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
    str := string(byteSlice)
    fmt.Println(str)
}
```
上面的程序同样输出：`Café`。  

## 通过 rune 切片创建字符串  
```golang
package main

import (  
    "fmt"
)

func main() {  
    runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str := string(runeSlice)
    fmt.Println(str)
}
```
在上面的程序中，`runeSlice` 包含了字符串 `"Señor"` 的 `Unicode` 码点（以 16 进制表示）。程序的输出为：`Señor`。  

## 字符串的长度  
utf8 包 提供了 `func RuneCountInString(s string) (n int)` 来获取字符串的长度，该方法接受一个字符串作为参数，并返回该字符串中 `rune` 的数量。  

（译者注： `RuneCountInString` 返回字符串中 `Unicode` 字符的个数，而 `len` 返回字符串中 `byte` 的个数，注意两者的区别。 ）  

```golang
package main

import (  
    "fmt"
    "unicode/utf8"
)

func length(s string) {  
    fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}
func main() {  

    word1 := "Señor" 
    length(word1)
    word2 := "Pets"
    length(word2)
}
```

上面程序的输出为：  

```golang
length of Señor is 5  
length of Pets is 4 
```

##  字符串是不可变的   

在 Go 中字符串是不可变的。字符串一旦被创建就无法改变。  

```golang
package main

import (  
    "fmt"
)

func mutate(s string)string {  
    s[0] = 'a'//any valid unicode character within single quote is a rune 
    return s
}
func main() {  
    h := "hello"
    fmt.Println(mutate(h))
}
```

上面的程序中，第 8 行我们试图改变字符串的第一个字符为 `a`。因为字符串是不可变的，因此这是非法的，将会报错：`main.go:8: cannot assign to s[0]`。  

为了改变一个字符串中的字符，我们需要先把字符串转换为 `rune` 切片，然后修改切片中的内容，最后将这个切片转换回字符串。  

```golang
package main

import (  
    "fmt"
)

func mutate(s []rune) string {  
    s[0] = 'a' 
    return string(s)
}
func main() {  
    h := "hello"
    fmt.Println(mutate([]rune(h)))
}
```

在上面的程序中，第 7 行 `mutate` 函数接受一个 `rune` 切片作为参数。然后将该切片的第一个元素改为 `a`，最后再转换回字符串并返回。该函数在程序中的第 13 行被调用。`h` 被转换为一个 `rune` 切片传递给 `mutate`。程序的输出为：`aello`。

字符串的介绍到此为止。感谢阅读。  
