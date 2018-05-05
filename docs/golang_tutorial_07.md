7 - 包
========================

上一节：[第六篇  函数](/docs/golang_tutorial_06.md)   
下一节：[第八篇  if else 语句](/docs/golang_tutorial_08.md)  

这是本Golang系列教程的第7篇。  

## 什么是包？为什么使用包？  

到目前为止我们见到的 Go 程序都只有一个文件，文件中包含了一个main函数和几个其他函数。在实际中这种将所有代码都放在一个文件里的组织方式是不可行的。这样的组织方式使得代码变得无法重用和维护困难。包`(package)`用于解决这样的问题。  

**包用于组织Go源代码，以获得更好的重用性和可读性**。包提供了代码封装的机制从而使得Go应用程序易于维护。例如，假设我们正在开发一个图像处理应用，它提供了诸如图像裁剪，锐化，模糊和增色等功能。一种组织代码的方式是将所有实现同一功能的代码放在一个独立的包中。例如裁剪功能可以放在一个单独包中，而锐化功能可以放在另一个包中。这种做法的优点是：增色功能可能需要做一些锐化的处理，那么增色代码中可以简单地导入（我们即将讨论导入）锐化包，使用其中提供的功能即可。这种方式使得代码变得更容易重用。  

我们将逐步创建一个计算矩形面积和对角线的应用程序。  

通过构建该程序，我们将更好的理解包。  

## main函数与main包  

每个可执行的Go程序都必须包含一个 main 函数。这个函数是执行程序的入口点。main 函数应该包含在 main 包中。  

**指定一个特定源文件属于一个包的语法为：package packagename，这条语句应该放在源文件的第一行**。  

下面让我们开始创建 main 函数和 main 包。**在 [工作空间目录]/src 目录下新建一个子目录，命名为 `geometry`**。在`geometry`目录下新建 `geometry.go` 文件。  

在 geometry.go 中写入以下代码:  

```golang
//geometry.go
package main 

import "fmt"

func main() {  
    fmt.Println("Geometrical shape properties")
}
```

`package main `这一行指定了该文件属于 `main` 包。`import "packagename" `语句用来导入一个包，这里我们导入 `fmt` 包，该包导出了 `Println` 方法。然后是 main 函数，在这里仅打印 `Geometrical shape properties`。  

执行 `go install geometry` 编译上面的程序。该命令在 geometry 目录下查找包含 main 函数的文件，在这里就是` geometry.go`。找到后编译该文件并在 [工作空间目录]/bin 目录下生成二进制文件`geometry`（在Windows下是 geometry.exe）。 现在的目录结构如下所示：  

```golang
src  
    geometry
            gemometry.go
bin  
    geometry
```

执行 [工作空间目录]/bin/geometry 运行该程序，其中 [工作空间目录] 需要换成自己的实际目录。这条命令运行在 bin 目录下的 geometry 二进制文件。你应该可以看到如下输出：  

```golang
Geometrical shape properties
```

## 创建自定义包  

下面我们将创建一个 `rectangle` 包，将与矩形相关的功能（计算矩形的面积和对角线）都放在这个包里。  

属于同一个包的源文件应该放在独立的文件夹中，按照Go的惯例，该文件夹的名字应该与包名相同。  

因此让我们在 `geometry` 目录下创建一个 `rectangle` 子目录。所有放在该目录下的源文件都应该以 `package rectangle`开头，用以表示这些源文件都属于 `rectangle` 包。  

```golang
//rectprops.go
package rectangle

import "math"

func Area(len, wid float64) float64 {  
    area := len * wid
    return area
}

func Diagonal(len, wid float64) float64 {  
    diagonal := math.Sqrt((len * len) + (wid * wid))
    return diagonal
}
```

在上面的代码中我们实现了两个函数 `Area` 和 `Diagonal` 分别用于计算矩形的面积和对角线。矩形的面积为长与宽的积。矩形的对角线为长与宽的平方和再开根号。这里调用 `math` 包中的 `Sqrt` 函数来计算平方根。  

*注意上面实现的两个函数的函数名 `Area` 和 `Diagonal` 都是以大写字母开头的。这是必要的，我们将很快解释为什么需要这样做。*  

在 `rectangle` 目录下新建`rectangle.go` ，编写如下代码：  

```golang
package rectangle

import "math"

func Area(len, wid float64) float64 {  
    area := len * wid
    return area
}

func Diagonal(len, wid float64) float64 {  
    diagonal := math.Sqrt((len * len) + (wid * wid))
    return diagonal
}
```

## 导入自定义包  

为了使用自定义包我们必须先导入它。用来导入自定义包的语法为：`import path`。我们必须指定 `path` 为相对于 `[工作空间目录]/src `目录的相对路径。我们当前的目录结构如下所示：  
```golang
src/
    geometry/
        geometry.go
        rectangle/
            rectangle.go
```

语句 `import "geometry/rectangle"` 表示我们要导入 rectangle 包。  

在 geometry.go 中添加如下代码：  

```golang
//geometry.go
package main 

import (  
    "fmt"
    "geometry/rectangle" //importing custom package
)

func main() {  
    var rectLen, rectWidth float64 = 6, 7
    fmt.Println("Geometrical shape properties")
        /*Area function of rectangle package used
        */
    fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
        /*Diagonal function of rectangle package used
        */
    fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))
}
```

上面的代码导入了 `rectangle` 包并且使用 `Area` 和 `Diagonal` 函数计算矩形的面积和对角线。在 Printf 中的 %.2f 格式化指示符表示仅保留浮点数的两位小数。程序输出如下：  

```golang
Geometrical shape properties  
area of rectangle 42.00  
diagonal of the rectangle 9.22 
```

**注意**：本教程 `geometry`目录中的包的导入路径为：` "golang-tutorial/geometry/rectangle"`，相对于src的绝对路径。  

## 导出名字  

我们将 `rectangle` 包中的两个函数名称 `Area` 和 `Diagonal` 的首字母大写，这在 Go 中有特殊的意义。**在 Go 中，任何以大写字母开头的变量名、函数名都是被导出的名字（exported name）。只有被导出的名字才能被其它包访问。在这里我们需要在 main 包中访问 Area 和 Diagonal 函数，因此将它们的首字母大写。**  

如果将 `rectprops.go` 中的 `Area(len, wid float64)` 改成 `area(len, wid float64)`，并且将` geometry.go` 中的` rectangle.Area(rectLen, rectWidth) `改成` rectangle.area(rectLen, rectWidth)`，那么运行程序时编译器将会报错：`geometry.go:11: cannot refer to unexported name rectangle.area`。因此，如果想访问包外的函数，必须将其首字母大写。  

## init 函数  

每一个包都可以包含一个 `init` 函数。该函数不应该有任何参数和返回值，并且在我们的代码中不能显式调用它。init 函数形式如下：  

```golang
func init() {  
}
```

**init 函数可用于执行初始化任务，也可用于在执行开始之前验证程序的正确性。**  

一个包的初始化顺序如下：  

* 包级别的变量首先被初始化  
* 接着 init 函数被调用。一个包可以有多个 init 函数（在一个或多个文件中），它们的调用顺序为编译器解析它们的顺序。  

如果一个包导入了另一个包，被导入的包先初始化。  

尽管一个包可能被包含多次，但是它只被初始化一次。  

下面让我们对我们的程序做一些修改来理解 init 函数。  

首先在` rectprops.go` 中添加一个 `init `函数：  

```golang
//rectprops.go
package rectangle

import "math"  
import "fmt"

/*
 * init function added
 */
func init() {  
    fmt.Println("rectangle package initialized")
}
func Area(len, wid float64) float64 {  
    area := len * wid
    return area
}

func Diagonal(len, wid float64) float64 {  
    diagonal := math.Sqrt((len * len) + (wid * wid))
    return diagonal
}
```

我们添加了一个简单的 `init` 函数，它仅打印：`rectangle package initialized`。  

现在我们来修改 `main` 包。我们知道矩形的 `length` 和 `width` 应该大于 0。我们将在 `geometry.go` 中添加 `init` 函数和包级别的变量来做此检查。  

修改 `geometry.go` 如下：  

```golang
//geometry.go
package main 

import (  
    "fmt"
    "geometry/rectangle" //importing custom package
    "log"
)
/*
 * 1. package variables
*/
var rectLen, rectWidth float64 = 6, 7 

/*
*2. init function to check if length and width are greater than zero
*/
func init() {  
    println("main package initialized")
    if rectLen < 0 {
        log.Fatal("length is less than zero")
    }
    if rectWidth < 0 {
        log.Fatal("width is less than zero")
    }
}

func main() {  
    fmt.Println("Geometrical shape properties")
    fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
    fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))
}
```

我们对` geometry.go` 做了如下修改：  

* `rectLen` 和 `rectWidth` **变量从 `main` 函数中移到了外面，成为了包级别的变量**。  
* 添加 `init` 函数。当 `rectLen` 或 `rectWidth` 小于 0 时，该函数利用` log.Fatal `打印一条日志并终止程序。  

main 包的初始化顺序为：  

* 首先初始化被导入的包。因此 rectangle 包先被初始化。   
* 然后初始化包级别的变量： rectLen 和 rectWidth 。  
* init 函数被调用。  

最后 main 函数被调用。  
```golang
rectangle package initialized  
main package initialized  
Geometrical shape properties  
area of rectangle 42.00  
diagonal of the rectangle 9.22 
```

正如预期的那样，`rectangle` 包的 `init` 函数首先被调用，接着是包级别的变量 `rectLen` 和 `rectWidth` 被初始化，接着是 `main` 包的 `init` 函数被调用，该函数检测 `rectLen` 和 `rectWidth` 是否小于 0，如果小于 0，则终止程序。我们将会在单独的教程里介绍 if 语句。现在你可以假定` if rectLen < 0` 将会检测 `rectLen` 是否小于 0，如果是，则终止程序。rectWidth 也是同样的处理。也就是说两个条件都为假程序才继续执行。最后，`main` 函数被调用。  

让我们再次修改程序来学习 `init` 函数的使用。  

将 `geometry.go` 中的` var rectLen, rectWidth float64 = 6, 7 `这一行改为 `var rectLen, rectWidth float64 = -6, 7`。这里将 rectLen 改为负值。  

```golang
rectangle package initialized
main package initialized
2018/03/24 11:37:21 length is less than zero
exit status 1
```

像上面一样， `rectangle` 包首先被初始化，然后是 `main` 包中的包级别变量 `rectLen` 和 `rectWidth` 初始化。接着调用 `main` 包的 `init` 函数，因为 rectLen 是小于 0 的，因此程序打印 `length is less than zero` 后退出。  

代码可以在 [github](https://github.com/golangbot/geometry) 上下载。  

## 使用空指示符  

在 Go 中只导入包却不在代码中使用它是非法的。如果你这么做了，编译器会报错。这样做的原因是为了避免引入过多未使用的包而导致编译时间的显著增加。将 geometry.go 中的代码替换为如下代码：  

```
//geometry.go
package main 

import (   

     "geometry/rectangle" //importing custom package

)
func main() {

}
```

上面的程序将会报错：`geometry.go:6: imported and not used: "geometry/rectangle"`  

但是在开发过程中，导入包却不立即使用它是很常见的。可以用空指示符（`_`）来处理这种情况。  

下面的代码可以避免抛出上面的错误：  

```golang
package main

import (  
    "geometry/rectangle" 
)

var _ = rectangle.Area //error silencer

func main() {

}
```

`var _ = rectangle.Area`这一行屏蔽了错误。我们应该跟踪这些**“错误消音器”**`（error silencer）`， 在开发结束时，我们应该去掉这些**“错误消音器”**，并且如果没有使用相应的包，这些包也应该被一并移除。因此，建议在 import 语句之后的包级别中写“错误消音器”。    

有时我们导入一个包只是为了确保该包初始化的发生，而我们不需要使用包中的任何函数或变量。例如，我们也许需要确保 rectangle 包的 init 函数被调用而不打算在代码中的任何地方使用这个包。空指示符仍然可以处理这种情况，像下面的代码一样：  

```golang
package main 

import (   
     _ "geometry/rectangle" 
)
func main() {

}
```

运行上面的程序，将会得到输出：`rectangle package initialized`。我们成功地初始化了这个包，即使在代码中的任何地方都没有使用它。  

希望你喜欢阅读。请留下宝贵的意见和反馈:)  