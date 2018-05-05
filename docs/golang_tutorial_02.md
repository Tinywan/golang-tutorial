2 - Hello World
========================

上一节：[第一篇  介绍和安装](/docs/golang_tutorial_01.md)   
下一节：[第三篇  变量](/docs/golang_tutorial_03.md)   

这是本Golang系列教程的第二篇。  

除了亲自写代码以外没有更好的方式来学习一门新的编程语言。在这篇教程中我们将开始编写我们的第一个程序。  

推荐使用 `Visual Studio Code ` 搭配 Go插件作为 IDE  

## 搭建Go工作空间  

在写代码之前，我们需要搭建Go工作空间。  

如果是 Mac 或者 Linux，Go的工作空间应该在 $HOME/go 目录。因此我们在 $HOME 目录下新建一个 go 子目录。  

如果是 Windows，Go的工作空间应该在 C:\Users\YourNamae\go 目录。因此我们在 C:\Users\YourName 目录下新建一个 go 子目录。  

也可以通过设置环境变量 GOPATH 来指定 Go 的工作空间为其他目录。但是为了简单让我们使用上面建立的目录。  

所有的源代码文件应该放在工作空间目录下的 src 子目录中，因此我们在工作空间目录（也就是上面创建的 go 目录）下创建 src 子目录。  

在 src 目录下，每一个 Go 的项目应该有一个单独的目录。因此我们在 src 目录下创建一个 hello 子目录用来存放我们的 hello world 程序。  

创建完上面的目录，目录结构如下：  

```golang
go/
    src/
        hello/
```

在 hello 目录下创建 helloworld.go，代码如下:  

```golang
package main

import "fmt"

func main() {  
    fmt.Println("Hello World")
}
```

创建完上面的程序，目录结构如下：  

```golang
go/
    src/
        hello/
            helloworld.go
```

## 运行 Go 程序  

可以有多种方式可以运行一个 Go 程序，下面一个一个地介绍。  

####  1、使用 go run 命令。在命令提示符中输入以下命令  

```golang
go run workspacepath/src/hello/helloworld.go
```

其中，workspacepath 应该是你自己的工作空间路径（在Windows下为 C:/Users/YourName/go ，在 Linux 和 Mac 下为 $HOME/go）。  

执行之后应该可以看到在控制台中打印了 Hello World。  

### 2、使用 go install 命令。在命令提示符中输入以下命令：  

```golang
go install hello
```

然后再输入以下命令运行程序：  

```golang
workspacepath/bin/hello
```

以上命令是在 Mac 或者 Linux下的命令，如果是在 Windows 下，则应该为 workspacepath\bin\hello.exe  

其中 workspacepath 应该替换成你自己的工作空间目录（在 Windows 下为 C:/Users/YourName/go ，在 Linux 和 Mac 下为 $HOME/go）。执行之后应该可以看到在控制台中打印了 Hello World。   

当执行 go install hello时，Go 工具会在工作空间中查找 hello 包（hello 称为一个包，我们将在以后的教程中解释什么是包）。接着它将会在 [工作空间]/bin 下创建（译者注：经过编译、链接）一个名为 hello（在Windows下为 hello.exe） 的二进制文件。执行完 go install hello之后的目录结构如下：  

```golang
go/
    bin/
        hello  -- 译者注：在Windows下是 hello.exe
    src/
        hello/
            helloworld.go
```

### 3、使用 Go playground

第三种运行 Go 程序的“酷酷的”方式是使用 Go playground。虽然这种方式有其自身的限制，但是这种方式在我们需要运行一个简单的程序时非常方便。我（原文作者）已经为 hello world 程序创建了一个 playgournd。可以点击这里在线执行它。（译者注：访问Go playground可能被墙。）  

你可以在 Go playground 与其他人分享你的源代码。（译者注：在 playground 页面可以点击 [Share] 按钮创建用于分享的URL）  

## 对 hello world 程序的简要解释

下面是我们刚刚编写的 hello world 程序代码：  

```golang
package main //1

import "fmt" //2

func main() { //3  
    fmt.Println("Hello World") //4
}
```

这里简要解释每一行都做了什么。我们会在以后的教程中更详细地介绍它们。  

`package main`：每个 Go 文件都必须以 package name 语句开头。包（package）提供了代码封装和重用。这里包的名字为 main。  

`import "fmt"`：导入 fmt 包，在 main 函数中将使用这个包打印文本到标准输出。  

`func main()`：main函数是一个特殊的函数，它是 Go 程序的入口点。main 函数必须包含在 main package 中。 { 和 } 表示 main 函数的开始和结束。  

`fmt.Println(“Hello World”)`：fmt 包里的 Println 函数用来打印文本到标准输出。  

本文由 [GCTT](https://github.com/studygolang/GCTT) 原创翻译，[Go语言中文网](https://studygolang.com/)首发。