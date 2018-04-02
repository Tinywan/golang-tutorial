如何编写Go代码
========================

简介：介绍如何使用go命令获取、构建和安装包，命令和运行测试。  

## 概述  

> 目录  

* 介绍  
* 代码组织  
  * 概观  
  * 工作区  
  * GOPATH环境变量  
  * 导入路径  
  * 你的第一个程序  
  * 你的第一个库  
* 包名称  
* 测试  
* 远程软件包  

> 介绍  

本文档演示了一个简单Go包的开发过程，并介绍了go工具，这是获取、构建和安装Go包和命令的标准方法。  

该go工具要求您以特定方式组织您的代码。  

## 代码组织  

> 概观  

请注意，这与其他编程环境不同，在这些编程环境中，每个项目都有一个单独的工作区，工作区与版本控制存储库紧密相关。  

* Go程序员通常将他们所有的Go代码保存在一个工作区中。  
* 工作区包含许多版本控制存储库 （例如，由Git管理）。  
* 每个存储库都包含一个或多个包。  
* 每个软件包由一个目录中的一个或多个Go源文件组成。  
* 包的目录的路径决定了它的导入路径。  

> 工作区  

工作空间是一个目录层次结构，其根目录包含三个目录：  

* src 包含Go源文件，  
* pkg 包含包对象  
* bin 包含可执行命令。  

该go工具构建源包并将生成的二进制文件安装到pkg和bin目录。  

该`src`子目录通常包含多个版本控制存储库（例如Git或Mercurial），用于跟踪一个或多个源包的开发。  

为了让您了解工作空间在实践中的外观，下面是一个例子：  

```golang
bin/
    hello                          # 命令可执行文件
    outyet                         # 命令可执行文件
pkg/
    linux_amd64/
        github.com/golang/example/
            stringutil.a           # 包对象
src/
    github.com/golang/example/
        .git/                      # Git存储库元数据
	hello/
	    hello.go               # 命令源
	outyet/
	    main.go                # command source
	    main_test.go           # test source
	stringutil/
	    reverse.go             # 包源码
	    reverse_test.go        # 测试源码
    golang.org/x/image/
        .git/                      # Git存储库元数据
	bmp/
	    reader.go              # package source
	    writer.go              # package source
    ... (many more repositories and packages omitted) ...
```

上面的树显示了一个包含两个存储库（example和image）的工作空间。该example库包含两个命令（hello 和outyet）和一个库（`stringutil`）。该image存储库包含该bmp包和其他几个包。  

典型的工作空间包含许多包含许多包和命令的源代码库。大多数Go程序员将他们所有的Go源代码和依赖关系保存在一个工作区中。  

命令和库由不同类型的源代码包构建而成。  

> 第一个程序  

要编译并运行一个简单的程序，首先选择一个包路径（我们将使用 `github.com/user/hello`）并在工作区内创建一个相应的包目录：  

```bash
mkdir $GOPATH/src/github.com/user/hello
```

接下来，创建一个名为hello.go该目录内的文件，其中包含以下Go代码。  

```golang
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}
```

现在，您可以使用该go工具构建和安装该程序  

```golang
$ go install github.com/user/hello // 编译包文件并且编译整个程序
```

请注意，您可以从系统上的任何位置运行此命令。该 go工具通过查找 `github.com/user/hello ` 指定的工作区内的程序包来查找源代码 `GOPATH`。  

如果`go install`从软件包目录运行，也可以省略软件包路径：  

```golang
$ cd $GOPATH/src/github.com/user/hello
$ go install
```

该命令构建`hello`命令，生成可执行的二进制文件。然后它将该二进制文件安装到工作空间的`bin`目录`hello`（或者在Windows下`hello.exe`）。在我们的例子中，那将是`$GOPATH/bin/hello`，这是`$HOME/go/bin/hello`。  

该`go`工具只会在发生错误时打印输出，因此如果这些命令不产生任何输出，它们将成功执行。  

您现在可以通过在命令行键入完整路径来运行该程序：  

```bash
$ $GOPATH/bin/hello
Hello, world.
```
或者，如您添加`$GOPATH/bin`到您的`PATH`，只需键入二进制名称：  

```golang
$ hello
Hello, world.
```

如果您正在使用源代码管理系统，现在应该是初始化存储库，添加文件并提交第一个更改的好时机。再一次，这一步是可选的：你不需要使用源代码控制来编写Go代码。  

```bash
$ cd $GOPATH/src/github.com/user/hello
$ git init
Initialized empty Git repository in /home/user/work/src/github.com/user/hello/.git/
$ git add hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 1 insertion(+)
  create mode 100644 hello.go
$ git remote add origin https://github.com/Tinywan/hello.git
$ git push -u origin master
```
将代码推送到远程存储库作为自己的练习。  

> 你的第一库  

我们来编写一个库并从`hello`程序中使用它。  

再次，第一步是选择一个包路径（我们将使用 `github.com/user/stringutil`）并创建包目录：  

```bash
$ mkdir $GOPATH/src/github.com/user/stringutil
```

接下来，使用以下内容创建一个名称在该目录中的文件。  

```golang
// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

现在，测试该软件包是用下面用`go build`命令进行代码编译：  

```bash
$ go build github.com/user/stringutil  // 测试编译，检查编译是否有编译错误
```

或者，如果您在包的源目录中工作，只需：  

```bash
$ go build 
```

这不会产生输出文件。为此，必须使用`go install`将包对象放置在工作空间的`pkg`目录中的方法。  

确认stringutil软件包构建完成后，修改原始文件`hello.go`（位于 `$GOPATH/src/github.com/user/hello`）以使用它：  

```golang
package main

import (
	"fmt"
	"github.com/user/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
```

无论何时该go工具**安装包或二进制文件**，**它也会安装它所具有的任何依赖关系**。所以当你安装hello程序  

```bash
$ go install github.com/user/hello
```

该`stringutil`包也将自动安装。  

运行该程序的新版本，您应该看到一条新的反转消息：  

```bash
$ hello
Hello, Go!
```

完成上述步骤后，您的工作空间应如下所示：  

```bash
bin/
    hello                 # command executable
pkg/
    linux_amd64/          # this will reflect your OS and architecture
        github.com/user/
            stringutil.a  # package object
src/
    github.com/user/
        hello/
            hello.go      # command source
        stringutil/
            reverse.go    # package source
```

请注意，`go install`将`stringutil.a`对象放置在`pkg/linux_amd64`镜像源目录中的目录中。这样未来的`go`工具调用可以找到包对象并避免不必要地重新编译包。该`linux_amd64`部分有助于交叉编译，并将反映您的系统的操作系统和体系结构。  

**Go命令可执行文件是静态链接的**，包对象不需要存在来运行Go程序。  

> 包名称  

Go源文件中的第一条语句必须是  

```bash
package name
```

这里`name`是对进口的包的默认名称。（包中的所有文件都必须使用相同的文件`name`）

Go的惯例是包名称是导入路径的最后一个元素：`crypto/rot13`应该命名导入为“ ” 的包rot13。

可执行命令必须始终使用`package main`。  

没有要求软件包名称在链接到一个二进制文件的所有软件包中唯一，**只需要导入路径（它们的完整文件名）是唯一的**。  

请参阅[Effective Go](http://127.0.0.1:8080/doc/effective_go.html#names)详细了解Go的命名约定。  

> 测试  

Go有一个由`go test `命令和`testing`包组成的轻量级测试框架。  

您通过创建一个名称以文件名结尾的文件来编写一个测试`_test.go` ，其中包含以TestXXX签名 命名的函数`func (t *testing.T)`。测试框架运行每个这样的功能; 如果函数调用失败函数（如t.Error或）` t.Fail`，则认为测试失败。  

`stringutil`通过创建`$GOPATH/src/github.com/user/stringutil/reverse_test.go`包含以下Go代码的文件 向测试包 添加测试。  

```golang
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

然后运行测试`go test`：  

```bash
$ go test github.com/user/stringutil
ok  	github.com/user/stringutil 0.165s
```

与往常一样，如果您正在go从软件包目录运行该工具，则可以省略软件包路径：

```bash
$ go test
ok  	github.com/user/stringutil 0.165s
```
运行`go help test`并查看 测试包文档以获取更多详细信息。  

> 远程软件包  

导入路径可以描述如何使用Git或Mercurial等版本控制系统获取软件包源代码。该go工具使用此属性从远程存储库自动获取软件包。例如，本文档中描述的示例也保存在GitHub托管的Git存储库中` github.com/golang/example`。如果您在存储库的导入路径中包含存储库URL，` go get`将自动获取，构建和安装它：  

```bash
$ go get github.com/golang/example/hello
$ $GOPATH/bin/hello
Hello, Go examples!
```

如果指定的软件包不在工作区中，`go get `则将其放入由指定的第一个工作区内`GOPATH`。（如果包已经存在，则`go get`跳过远程抓取并且行为与`go install`相同）  

发出上述`go get`命令后，工作空间目录树现在应该如下所示：  

```golang
bin/
    hello                           # command executable
pkg/
    linux_amd64/
        github.com/golang/example/
            stringutil.a            # package object
        github.com/user/
            stringutil.a            # package object
src/
    github.com/golang/example/
	.git/                       # Git repository metadata
        hello/
            hello.go                # command source
        stringutil/
            reverse.go              # package source
            reverse_test.go         # test source
    github.com/user/
        hello/
            hello.go                # command source
        stringutil/
            reverse.go              # package source
            reverse_test.go         # test source
```

hello在GitHub上托管的命令取决于`stringutil`同一个存储库中的软件包。在导入`hello.go`文件中使用相同的导入路径约定，所以` go get`命令能够找到并安装相关的包了。  

导入`“github.com/golang/example/stringutil”`这个约定是让你的Go软件包可供其他人使用的最简单的方法。在转到维基 和`godoc.org` 提供外部围棋项目清单。

有关在go工具中使用远程存储库的更多信息，请参阅 [go help importpath](http://127.0.0.1:8080/cmd/go/#hdr-Remote_import_paths)。