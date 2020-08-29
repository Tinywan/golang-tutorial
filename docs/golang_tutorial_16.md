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

下面的程序说明了如何定义一个具名结构体 Employee 的变量。  

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

在上面的程序中，通过 `emp6.firstName` 访问 emp6 中的字段 `firstName`。随后我们修改了该员工的薪水，程序的输出为：  

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

