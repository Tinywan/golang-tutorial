
35 - 读取文件  
========================

上一节：[第十六篇 结构体](/docs/golang_tutorial_16.md)  
下一节：[第十八篇 接口一](/docs/golang_tutorial_18.md)  

这是本Golang系列教程的第35篇。   

文件读取是在任何编程语言中执行的最常见操作之一。在本教程中，我们将了解如何使用Go读取文件。  

本教程包含以下部分  

* 将整个文件读入内存  
  *  使用绝对文件路径  
  *  将文件路径作为命令行标志传递  
  *  将文件捆绑在二进制文件中   
* 以小块读取文件  
* 逐行读取文件  

## 将整个文件读入内存  

最基本的文件操作之一是将整个文件读入内存。这是在[ioutil](https://golang.org/pkg/io/ioutil/)包的[ReadFile](https://golang.org/pkg/io/ioutil/#ReadFile)函数的帮助下完成的。  

让我们从go程序所在的目录中读取一个文件。在`GOROOT`内部创建了一个文件夹，里面有一个文本文件`test.txt`，可以从我们的Go程序中读取`filehandling.go`。`test.txt`包含文本`“Hello World. Welcome to file handling in Go“`。这是我的文件夹结构。   

```golang
src  
    filehandling
        filehandling.go
        test.txt
```

让我们马上看看代码吧。   

```golang
package main

import (  
    "fmt"
    "io/ioutil"
)

func main() {  
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println("Contents of file:", string(data))
}
```

请从当地环境运行此程序，因为无法在操场上阅读文件。  

上面的程序中的9行，读取文件并返回存储在其中的字节片`data`。排队号码 14我们转换`data`为a string并显示文件的内容。  

请从test.txt所在的位置运行该程序。  

例如，在`linux / mac`的情况下，如果`test.txt`位于`home /naveen/go/src/filehandling`，则使用以下步骤运行该程序，   

```golang
$]cd /home/naveen/go/src/filehandling/
$]go install filehandling
$]workspacepath/bin/filehandling
```

对于`windows`，如果`test.txt`位于`C：\ Users \ naveen.r \ go \ src \ filehandling`，则使用以下步骤运行此程序。  

```golang
> cd C:\Users\naveen.r\go\src\filehandling
> go install filehandling
> workspacepath\bin\filehandling.exe 
```

程序的输出为：  

```golang
Contents of file: Hello World. Welcome to file handling in Go.  
```

如果此程序从任何其他位置运行，例如尝试运行该程序`/home/userdirectory`，它将打印以下错误。  
```golang
File reading error open test.txt: The system cannot find the file specified.    
```

原因是Go是一种编译语言。`go install`它是什么，它从源代码创建二进制文件。二进制文件独立于源代码，可以从任何位置运行。由于`test.txt`在运行二进制文件的位置找不到，程序会抱怨它无法找到指定的文件。  

有三种方法可以解决这个问题，  

* 1、使用绝对文件路径  
* 2、将文件路径作为命令行标志传递  
* 3、将文本文件与二进制文件捆绑在一起  

让我们一个一个讨论。  

Go协程的介绍就到这里。祝你有美好的一天！  
希望你喜欢阅读。请留下宝贵的意见和反馈:)  