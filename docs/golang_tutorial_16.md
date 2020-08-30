16 - 结构体  
========================

上一节：[第十五篇 指针](/docs/golang_tutorial_15.md)   
下一节：[第十七篇 方法](/docs/golang_tutorial_17.md)  

这是本Golang系列教程的第16篇。

## 什么是结构体？  

结构体（struct）是用户自定义的类型，它代表若干字段的集合。有些时候将多个数据看做一个整体要比单独使用这些数据更有意义，这种情况下就适合使用结构体。  

比如，我们可以将一个员工的 firstName, lastName 和 age 三个属性打包在一起成为一个 `employee` 结构体。  

## 结构体的声明  

```golang
type Employee struct {  
    firstName string
    lastName  string
    age       int
}
```

上面的代码片段声明了一个名为 `Employee` 的结构体类型，它拥有 `firstName`，`lastName` 和 `age` 三个字段。同一类型的多个字段可以合并到一行（用逗号分隔），并将类型放在后面。上面的结构体中 `firstName` 与 `lastName` 都是 `string` 类型，因此可以将它们写在一起。

```golang
type Employee struct {  
    firstName, lastName string
    age, salary         int
}
```

*(尽管以上语句节省了代码行数，但是这会导致字段定义不够清晰，请尽量避免使用以上语句的定义方式)*

上面的结构体 `Employee` 是一个**具名结构体 `（named structs）`** ，因为它创建了一个具有名字的结构体类型：`Employee`。我们可以定义具名结构体类型的变量。  

我们也可以定义一个没有类型名称的结构体，这种结构体叫做**匿名结构体 `（anonymous structs）`**。  

```golang
var employee struct {  
        firstName, lastName string
        age int
}
```  

上面的代码片段声明了一个匿名结构体变量 `employee`。  

## 定义具名结构体变量  

下面的程序说明了如何定义一个具名结构体 `Employee` 的变量。  

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {

    //creating structure using field names
    emp1 := Employee{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Anderson",
    }

    //creating structure without using field names
    emp2 := Employee{"Thomas", "Paul", 29, 800}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)
}
```  

在上面的程序中，我们定义了一个名为 `Employee` 的结构体类型。我们可以通过指定字段的名称和对应的值来创建一个结构体变量，比如在第 15 行，我们就是通过这种方式定义了 `Employee` 类型的一个结构体变量 `emp1`。这里字段名称的顺序没必要和声明结构体类型时的一致。例如这里我们将 `lastName` 放在了最后，程序同样正确运行。  

在定义结构体变量时也可以忽略字段名称，例如在第 23 行，我们定义 `emp2` 时没有指定字段名称。但是通过这种方式定义的结构体变量时，字段值的顺序必须与声明结构体类型时字段的顺序保持一致。  

上面的程序输出如下：  

```golang
Employee 1 {Sam Anderson 25 500}  
Employee 2 {Thomas Paul 29 800} 
```  

## 定义匿名结构体变量  

```golang
package main

import (  
    "fmt"
)

func main() {  
    emp3 := struct {
        firstName, lastName string
        age, salary         int
    }{
        firstName: "Andreah",
        lastName:  "Nikola",
        age:       31,
        salary:    5000,
    }

    fmt.Println("Employee 3", emp3)
}
```  

在上面的程序中，第 8 行定义了一个**匿名结构体变量** `emp3`。正如我们之前提到的，这种结构体叫做匿名结构体，因为它只创建了一个新的结构体变量 `emp3`，而没有定义新的结构体类型。

程序的输出为：  

```golang
Employee 3 {Andreah Nikola 31 5000}  
```  

## 结构体变量的零值  

当定义一个结构体变量，但是没有提供初始值时，结构体中的字段会被赋予它们各自类型的零值。  

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    var emp4 Employee //zero valued structure
    fmt.Println("Employee 4", emp4)
}
```  

上面的程序定义了 `emp4` 但是没有赋予任何初值。因此 `firstName` 和 `lastName` 被赋值为 `string` 类型的零值，也就是空字符串（`""`）。`age` 和 `salary` 被赋值为 `int` 类型的零值，也就是 0。程序的输出为：  

```golang
Employee 4 {  0 0} 
```  

我们也可以指定某些字段的值而忽略其它字段。在这种情况下，被忽略的字段被赋予相应类型的零值。  

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp5 := Employee{
        firstName: "John",
        lastName:  "Paul",
    }
    fmt.Println("Employee 5", emp5)
}
```  

在上面的程序中，第14和15行，`firstName` 和 `lastName` 被提供了初始值，而 `age` 和 `salary` 没有。因此 `age` 和 `salary` 被指定为相应类型的零值。程序的输出为：  

```golang
Employee 5 {John Paul 0 0} 
```  

## 访问结构体中的字段  

使用点 `.` 操作符来访问结构体中的字段。  

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

func main() {  
    emp6 := Employee{"Sam", "Anderson", 55, 6000}
    fmt.Println("First Name:", emp6.firstName)
    fmt.Println("Last Name:", emp6.lastName)
    fmt.Println("Age:", emp6.age)
    fmt.Printf("Salary: $%d", emp6.salary)
    emp6.salary = 6500
    fmt.Printf("New Salary: $%d", emp6.salary)
}
```  

在上面的程序中，通过 `emp6.firstName` 访问 `emp6` 中的字段 `firstName`。随后我们修改了该员工的薪水，程序的输出为：  

```golang
First Name: Sam  
Last Name: Anderson  
Age: 55  
Salary: $6000
New Salary: $6500
```  

## 指向结构体的指针

我们也可以创建指向结构体的指针：

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {  
    emp8 := &Employee{
        firstName: "Sam",
        lastName:  "Anderson",
        age:       55,
        salary:    6000,
    }
    fmt.Println("First Name:", (*emp8).firstName)
    fmt.Println("Age:", (*emp8).age)
}
```

上面的程序中，`emp8` 是一个指向 `Employee` 结构体的指针。使用`(\*emp8).firstName` 来访问 `emp8` 中的 `firstName` 字段，程序输出：

```golang
First Name: Sam  
Age: 55 
```

**Go 语言提供了一种简写方法：用 `emp8.firstName` 来代替 `(\*emp8).firstName` 的间接引用。**

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {  
    emp8 := &Employee{
        firstName: "Sam",
        lastName:  "Anderson",
        age:       55,
        salary:    6000,
    }
    fmt.Println("First Name:", emp8.firstName)
    fmt.Println("Age:", emp8.age)
}
```

我们用 `emp8.firstName` 来获取 `firstName` 字段，程序输出：

```golang
First Name: Sam  
Age: 55 
```

## 匿名字段

在创建结构体时，我们可以只定义字段的类型，不包含字段的名称。这种字段叫做匿名字段。

以下代码创建了一个叫 `Person` 的结构体，包含两个匿名字段：`string` 和 `int`。

```golang
type Person struct {  
    string
    int
}
```

**尽管匿名字段没有名称，但是它的默认名称就是它的类型。** 例如，上面例子中的 `Person` 结构体，它的两个字段的默认名称是 `string` 和 `int`。

```golang
package main

import (  
    "fmt"
)

type Person struct {  
    string
    int
}

func main() {  
    p1 := Person{
        string: "Ricky",
        int:    10,
    }
    fmt.Println(p1.string)
    fmt.Println(p1.int)
}
```

在第 17 和 18 行，我们利用匿名字段的类型：`string` 和 `int`，访问了 `Person` 中的匿名字段。程序输出：

```golang
Ricky
10
```

## 嵌套结构体

我们也可以创建一个包含结构体作为字段的结构体，这叫做嵌套结构体。例如：

```golang
package main

import (  
    "fmt"
)

type Address struct {  
    city  string
    state string
}

type Person struct {  
    name    string
    age     int
    address Address
}

func main() {  
    p := Person{
        name: "Ricky",
        age:  10,
        address: Address{
            city:  "Chicago",
            state: "Illinois",
        },
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.address.city)
    fmt.Println("State:", p.address.state)
}
```

这里，名叫 `Person` 的结构体包含一个叫做 `address` 的字段，这个字段本身也是一个结构体。以上程序输出：

```golang
Name: Ricky 
Age: 10
City: Chicago
State: Illinois
```

## 提阶字段

如果结构体中的匿名字段也是一个结构体，那么这个匿名结构体字段叫做提阶字段，因为我们可以通过外部结构体变量直接访问匿名结构体中的字段，就像这些字段原本属于外部结构体一样。这个定义有些复杂，所以我们来看一些代码来更好地理解这一概念：

```golang
type Address struct {  
    city string
    state string
}
type Person struct {  
    name string
    age  int
    Address
}
```

在上面的代码中，名叫 `Person` 的结构体包含了一个匿名结构体字段 `Address`。我们将 `Address` 的字段：`city` 和 `state`，叫做提阶字段，因为我们可以在 `Person` 结构体里直接访问这两个字段，就如同这两个字段在 `Person` 结构体里被声明了一样。

```golang
package main

import (  
    "fmt"
)

type Address struct {  
    city  string
    state string
}
type Person struct {  
    name string
    age  int
    Address
}

func main() {  
    p := Person{
        name: "Ricky",
        age:  10,
        Address: Address{
            city:  "Chicago",
            state: "Illinois",
        },
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.city)   //city is promoted field
    fmt.Println("State:", p.state) //state is promoted field
}
```

在第 29 和 30 行，提阶字段 `city` 和 `state` 可以直接在结构体 `p` 中被访问（利用 `p.city` 和 `p.state` 语句），就好像这两个字段是在 `p` 中被声明的。程序输出：

```golang
Name: Ricky
Age: 10
City: Chicago
State: Illinois
```

## 导出结构体和字段

如果一个结构体类型的名称以大写字母开头，那么它就是一个**导出类型**，可以被其他的包访问。同理，如果一个结构体的字段名称以大写字母开头，那么该字段也可以被其他包访问。

我们写一段程序，来理解这一概念：

在 `Documents` 目录下创建一个叫做 `structs` 的文件夹（译者注：本文作者在 `Documents` 目录下创建了文件夹，你也可以根据个人喜好将文件夹创建在任何目录下）：

```
mkdir ~/Documents/structs
```

我们再来创建一个叫做 `structs` 的 Go 模块：

```
cd ~/Documents/structs/  
go mod init structs  
```

在 `structs` 文件夹内创建另一个文件夹：`computer`：

```
mkdir computer
```

在 `computer` 文件夹内，创建一个叫做 `spec.go` 的文件，写入以下内容：

```golang
package computer

type Spec struct { //exported struct  
    Maker string //exported field
    Price int //exported field
    model string //unexported field
    
}
```

上面的代码创建了一个叫做 `computer` 的包，里面包含了导出类型的结构体 `Spec`，以及两个导出类型的字段：`Maker` 和 `Price`，还有一个非导出类型的字段：`model`。我们来尝试导入这个包，并使用 `Spec` 结构体。

在 `structs` 文件夹内创建 `main.go` 文件，并写入以下内容：

```golang
package main

import (  
    "structs/computer"
    "fmt"
)

func main() {  
    spec := computer.Spec {
            Maker: "apple",
            Price: 50000,
        }
    fmt.Println("Maker:", spec.Maker)
    fmt.Println("Price:", spec.Price)
}
```

现在，`structs` 文件夹的结构如下：

```
├── structs
│   ├── computer
│   │   └── spec.go
│   ├── go.mod
│   └── main.go
```

在程序的第 4 行中，我们导入了 `computer` 包。在 13 和 14 行，我们访问了 `Spec` 结构体中的 `Maker` 和 `Price` 两个导出类型字段，会得到输出：

```golang
Maker: apple  
Price: 50000  
```

如果我们试图访问非导出类型字段 `model`，编译器就会报错。我们在 `main.go` 中重新写入以下代码：

```golang
package main

import (  
    "structs/computer"
    "fmt"
)

func main() {  
    spec := computer.Spec {
            Maker: "apple",
            Price: 50000,
            model: "Mac Mini",
        }
    fmt.Println("Maker:", spec.Maker)
    fmt.Println("Price:", spec.Price)
}
```

在第 12 行中，我们尝试访问非导出类型字段 `model`，这会导致编译错误：

```
# structs
./main.go:12:13: unknown field 'model' in struct literal of type computer.Spec
```

## 结构体的比较

**结构体属于值类型，如果两个结构体的字段都是可比较的，那么这两个结构体就是可比较的。如果两个结构体变量对应的字段都是相等的，那么这两个结构体就是相等的。**

```golang
package main

import (  
    "fmt"
)

type name struct {  
    firstName string
    lastName  string
}

func main() {  
    name1 := name{
        firstName: "Steve",
        lastName:  "Jobs",
    }
    name2 := name{
        firstName: "Steve",
        lastName:  "Jobs",
    }
    if name1 == name2 {
        fmt.Println("name1 and name2 are equal")
    } else {
        fmt.Println("name1 and name2 are not equal")
    }

    name3 := name{
        firstName: "Steve",
        lastName:  "Jobs",
    }
    name4 := name{
        firstName: "Steve",
    }

    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
    }
}
```

上面的程序中，结构体 `name` 包含两个字符串类型的字段。由于字符串是可以比较的，我们因此可以比较两个 `name` 类型的结构体变量。程序输出：

```golang
name1 and name2 are equal  
name3 and name4 are not equal  
```

**但是，如果两个结构体变量包含的字段是不可比较的，那么这两个结构体变量也不可比较。**

```golang
package main

import (  
    "fmt"
)

type image struct {  
    data map[int]int
}

func main() {  
    image1 := image{
        data: map[int]int{
            0: 155,
        }}
    image2 := image{
        data: map[int]int{
            0: 155,
        }}
    if image1 == image2 {
        fmt.Println("image1 and image2 are equal")
    }
}
```

在上面的程序中，结构体类型 `image` 包含了 `map` 类型的字段：`data`。由于 `map` 类型是不可比较的，因此 `image1` 和 `image2` 也不可以进行比较。运行以上程序会发生编译错误：

```
./prog.go:20:12: invalid operation: image1 == image2 (struct containing map[int]int cannot be compared)
```

结构体的介绍到此结束，感谢阅读。
