
22 - 信道  
========================

上一节：[第十六篇 结构体](/docs/golang_tutorial_16.md)  
下一节：[第十八篇 接口一](/docs/golang_tutorial_18.md)  

这是本Golang系列教程的第22篇。   

在[上一篇教程](http://blog.csdn.net/u011304970/article/details/75096323)中，我们讨论了如何使用协程实现并发。在这篇教程中，我们将讨论信道以及如何使用信道实现协程间通信。   

## 什么是信道？  

信道`（Channel）`可以被认为是协程之间通信的管道。与水流从管道的一端流向另一端一样，数据可以从信道的一端发送并在另一端接收。  

## 声明信道  

每个信道都有一个与之关联的类型。此类型是允许信道传输的数据类型，除此类型外不能通过信道传输其他类型。  

`chan T` 是一个 `T` 类型的信道。  

信道的 `0` 值为 `nil`。值为 `nil` 的信道变量没有任何用处，我们需要通过内置函数 `make` 来创建一个信道，就像创建 `map` 和 `slice` 一样。  

下面的代码声明了一个信道：  

```golang
package main

import "fmt"

func main() {  
    var a chan int
    if a == nil {
        fmt.Println("channel a is nil, going to define it")
        a = make(chan int)
        fmt.Printf("Type of a is %T", a)
    }
}
```
因为信道的 `0` 值为 `nil`，因此第 6 行声明的信道 `a` 的值为 `nil`。因此执行 `if` 里面的语句创建信道。上面的程序中 `a` 是一个 `int` 类型的信道。程序的输出为：  

```golang
channel a is nil, going to define it  
Type of a is chan int 
```
像往常一样，速记声明也是定义信道的一种有效而简洁的方式：  
```golang
a := make(chan int) 
```
上面的这行代码同样定义了一个 `int` 型的**信道**。  

## 通过信道发送和接收数据  

通过信道发送和接收数据的语法如下：  

```golang
data := <- a // read from channel a  
a <- data // write to channel a  
```
箭头的指向说明了数据是发送还是接收。  

在第一行，**箭头**的方向是从 `a` 向外指，因此我们正在从信道 `a` 中读取数据并将读取的值赋值给变量 `data` 。  

在第二行，**箭头**的方式是指向 `a` ，因此我们正在向信道 `a` 中写入数据。  

## 发送和接收默认是阻塞的  

通过信道发送和接收数据默认是阻塞的。这是什么意思呢？当数据发送给信道后，程序流程在发送语句处阻塞，直到其他协程从该信道中读取数据。  

同样地，当从信道读取数据时，程序在读取语句处阻塞，直到其他协程发送数据给该信道。  

信道的这种特性使得协程间通信变得高效，而不是向其他编程语言一样，显式的使用锁和条件变量来达到此目的。  

## 信道的一个例子  

理论到此为止：） 让我们通过一个程序来理解协程之间如何使用信道进行通信。  

我们将用信道来重写在上一篇教程协程中的一个例子。  

如下是那篇教程中的一个例子：  

```golang
package main

import (  
    "fmt"
    "time"
)

func hello() {  
    fmt.Println("Hello world goroutine")
}
func main() {  
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("main function")
}
```

这是上一篇教程中的例子，我们通过使用 `Sleep` 来使主协程休眠，以等待 `hello` 协程执行结束。如果你不明白这是为什么，请阅读上一篇教程：`协程` 。  

我们用信道重写上面的程序，如下：   

```golang
package main

import (  
    "fmt"
)

func hello(done chan bool) {  
    fmt.Println("Hello world goroutine")
    done <- true
}
func main() {  
    done := make(chan bool)
    go hello(done)
    <-done
    fmt.Println("main function")
}
```

在上面的程序中，我们在第 12 行定义了一个 `bool` 类型的信道 `done`，然后将它作为参数传递给 `hello` 协程。在第 14 行，我们从信道 `done` 中读取数据。程序将在这一行被阻塞直到其他协程向信道 `done` 里写入数据，在未读取到数据之前程序将在这一行一直等待而不会执行下一行语句。因此这里消除了在原程序中使用 `time.Sleep` 来阻止主协程退出的必要。  

`<-done` 这一行从信道 `done` 中读取数据，但是没有使用该数据，也没有将它赋值给其他变量，这是完全合法的。  

现在我们的 `main` 协程被阻塞，等待从信道 `done` 中读取数据。`hello` 协程接受信道 `done` 作为参数，打印 `Hello world goroutine` 然后将数据写入信道 `done` 中。当写入完毕后，`main` 协程从信道 `done` 中接收到数据，`main` 协程解除阻塞，继续执行下一条语句，打印：`main function`。  

程序的输出为：  

```golang
Hello world goroutine  
main function 
```

让我们修改上面程序，在 `hello` 协程中加入一个休眠，来更好的理解阻塞的概念。  

```golang
package main

import (  
    "fmt"
    "time"
)

func hello(done chan bool) {  
    fmt.Println("hello go routine is going to sleep")
    time.Sleep(4 * time.Second)
    fmt.Println("hello go routine awake and going to write to done")
    done <- true
}
func main() {  
    done := make(chan bool)
    fmt.Println("Main going to call hello go goroutine")
    go hello(done)
    <-done
    fmt.Println("Main received data")
}
```

在上面程序中的第 10 行，我们在 `hello` 函数中增加了`4` 秒钟的休眠。  

该程序首先打印 `Main going to call hello go goroutine` 。然后 `hello` 协程开始执行，它将打印 `hello go routine is going to sleep`，然后 `hello` 协程休眠 `4` 秒，在这期间， `main` 协程由于在等待从信道 `done` 中读取数据而始终阻塞（在 `<-done` 这一行）。`4` 秒中之后， `hello` 协程打印：`hello go routine awake and going to write to done`，接着 `main` 协程打印：`Main received data` 。  

## 信道的另一个例子  
让我们再写一个例子来更好的理解信道。该程序打印一个数字的每一位的**平方和**与**立方和**，并将平方和与立方和相加得出最后的结果。  

例如，输入`123` ，程序将做如下计算以得出最后结果：

```golang
squares = (1 * 1) + (2 * 2) + (3 * 3) 
cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3) 
output = squares + cubes = 49
```

我们将平方和的计算与立方和的计算分别放在一个协程中执行，最后在主协程中将它们的计算结果求和。  

```golang
package main

import (  
    "fmt"
)

func calcSquares(number int, squareop chan int) {  
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit
        number /= 10
    }
    squareop <- sum
}

func calcCubes(number int, cubeop chan int) {  
    sum := 0 
    for number != 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
    }
    cubeop <- sum
} 

func main() {  
    number := 589
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares(number, sqrch)
    go calcCubes(number, cubech)
    squares, cubes := <-sqrch, <-cubech
    fmt.Println("Final output", squares + cubes)
}
```

在第 7 行，函数 `calcSquares` 计算 `number` 每一位的平方和，并将结果发送给信道 `squareop`。

同样地，在第 17 行，函数 `calcCubes` 计算 `number` 每一位的立方和，并将结果发送给信道 `cubeop`。  

这两个函数接受不同的信道作为参数，并分别运行在各自的协程中（第31行和32行），最后将结果写入各自的信道。主协程在第 33 行同时等待这两个信道中的数据。一旦从这两个信道中接收到数据，它们分别被存放在变量 `squares` 和 `cubes` 中，最后将它们的和打印出来。程序的输出为：  

```golang
Final output 1536  
```
## 死锁
使用信道是要考虑的一个重要因素是死锁`（Deadlock`）。如果一个协程发送数据给一个信道，而没有其他的协程从该信道中接收数据，那么程序在运行时会遇到死锁，并触发 `panic` 。  

同样地，如果一个协程从一个信道中接收数据，而没有其他的协程向这个信道发送数据，那么程序同样造成死锁，触发 `panic` 。  

```golang
package main

func main() {  
    ch := make(chan int)
    ch <- 5
}
```

上面的程序中，创建了一个信道 `ch`，并通过 `ch <- 5` 向其中写入 `5` 。这个程序中没有其他协程从 `ch` 中接收数据，因此程序在运行时触发 `panic`，错误如下：  

```golang
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:  
main.main()  
    /tmp/sandbox249677995/main.go:6 +0x80
```

## 单向信道  
目前我们讨论的信道都是双向信道，数据既可以发送到双向信道，也可以从双向信道中读取。同样也可以创建单向信道，即只能发送数据或只能接收数据的信道。  

```golang
package main

import "fmt"

func sendData(sendch chan<- int) {  
    sendch <- 10
}

func main() {  
    sendch := make(chan<- int)
    go sendData(sendch)
    fmt.Println(<-sendch)
}
```

在上面程序中的第 10 行，我们创建了一个只写`（send only）`信道 `sendch` 。`chan<- int` 表示只能发送数据，因为箭头的方向指向 `chan`。在第 12 行，我们试图从一个只写信道中接收数据，这是非法的，程序将无法通过编译，编译器报错如下：  

```golang
main.go:11: invalid operation: <-sendch (receive from send-only type chan<- int)
```

一切都很好，但是如果无法读取，创建一个只写通道有什么用呢？  

这就是信道转型的用途。可以将双向信道转换为只写或只读信道，但是反过来却不行。  

```golang
package main

import "fmt"

func sendData(sendch chan<- int) {  
    sendch <- 10
}

func main() {  
    chnl := make(chan int)
    go sendData(chnl)
    fmt.Println(<-chnl)
}
```

在上面程序中的第 10 行，创建了一个双向信道 `chnl`。在第 11 行它被作为参数传递给协程 `sendData`。第 5 行，`sendData` 函数通过形参 `sendch chan<- int` 将该信道转换成只写信道。因此在 `sendData` 中该信道为只写信道，而在主协程中该信道为双向信道。程序将打印：`10`。  

## 关闭信道以及使用 range for 遍历信道  
发送者可以关闭信道以通知接收者将不会再发送数据给信道。  

在从信道接收数据时，接收者可以通过一个额外的变量来检测信道是否已经被关闭。  

```golang
v, ok := <- ch
```
上面的语句中 `ok` 返回 `true` 表示成功的接收到了发送的数据，如果 `ok` 返回 `false` 则表示信道已经被关闭。从已关闭的信道中读取到的数据为信道类型的 `0` 值。比如从一个被关闭的 `int` 信道中读取数据，那么将得到 `0` 。  

```golang
package main

import (  
    "fmt"
)

func producer(chnl chan int) {  
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl)
}
func main() {  
    ch := make(chan int)
    go producer(ch)
    for {
        v, ok := <-ch
        if ok == false {
            break
        }
        fmt.Println("Received ", v, ok)
    }
}
```

上面的程序中，协程 `producer` 向信道 `chnl` 中写入 `0` 到 `9` 并关闭该信道。主协程在第 16 行进行无限 `for` 循环，并在循环中检测 `ok` 的值判断信道是否已经被关闭（第 18 行）。如果 `ok` 是 `false` 表示信道已经被关闭，则通过 `break` 退出循环。否则接收数据并打印 `ok` 的值。程序的输出为：  

```golang
Received  0 true  
Received  1 true  
Received  2 true  
Received  3 true  
Received  4 true  
Received  5 true  
Received  6 true  
Received  7 true  
Received  8 true  
Received  9 true  
```
`range for` 可以用来接收一个信道中的数据，直到该信道关闭。  

让我们用 `range for` 重写上面的程序：  

```golang
package main

import (  
    "fmt"
)

func producer(chnl chan int) {  
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl)
}
func main() {  
    ch := make(chan int)
    go producer(ch)
    for v := range ch {
        fmt.Println("Received ",v)
    }
}
```
在第6 行，`for range` 不断从信道 `ch` 中接收数据，直到该信道关闭。一旦 `ch` 关闭，循环自动退出。程序的输出如下：  
```golang
Received  0  
Received  1  
Received  2  
Received  3  
Received  4  
Received  5  
Received  6  
Received  7  
Received  8  
Received  9  
```

如果仔细观察该程序，你可以注意到，在 `calcSquares` 和 `calcCubes` 函数中查找一个数的每一位的代码重复了。我们将这段代码提取到一个单独的函数，并异步调用它。  

```golang
package main

import (  
    "fmt"
)

func digits(number int, dchnl chan int) {  
    for number != 0 {
        digit := number % 10
        dchnl <- digit
        number /= 10
    }
    close(dchnl)
}
func calcSquares(number int, squareop chan int) {  
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit
    }
    squareop <- sum
}

func calcCubes(number int, cubeop chan int) {  
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit * digit
    }
    cubeop <- sum
}

func main() {  
    number := 589
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares(number, sqrch)
    go calcCubes(number, cubech)
    squares, cubes := <-sqrch, <-cubech
    fmt.Println("Final output", squares+cubes)
}
```

在上面的程序中，函数 `digits` 包含查找一个数的每一位的逻辑，该函数在 `calcSquares` 和 `calcCubes` 中异步调用。在第 13 行，当数字中没有更多的位数时，信道被关闭。`calcSquares` 和 `calcCubes` 分别在各自的信道上使用 `range for` 直到信道被关闭。程序中的其他部分都是一样的。程序的输出依然是：  

```golang
Final output 1536  
```

这就来到了本教程的最后。信道中还有更多的概念，比如缓冲信道，工作池和 `select` 。我们将在单独的教程中讨论它们。感谢阅读，祝你有美好的一天。  

希望你喜欢阅读。请留下宝贵的意见和反馈:)  