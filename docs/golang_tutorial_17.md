
17 - 方法  
========================

上一节：[第十六篇 结构体](/docs/golang_tutorial_16.md)  
下一节：[第十八篇 接口一](/docs/golang_tutorial_18.md)  

这是本Golang系列教程的第17篇。   

## 什么是方法？  

一个方法只是一个函数，它有一个特殊的接收者（`receiver`）类型，该接收者放在 `func` 关键字和函数名之间。接收者可以是结构体类型或非结构体类型。可以在方法内部访问接收者。  

通过下面的语法创建一个方法：  

```golang
func (t Type) methodName(parameter list) {  
}
```

上面的代码片段创建了一个名为 `methodName` 的方法，该方法有一个类型为 Type 的接收者。  

## 案例  

让我们编写一个简单的程序，它创建一个结构体类型的方法并调用它。  

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    name     string
    salary   int
    currency string
}

/*
 displaySalary() method has Employee as the receiver type
*/
func (e Employee) displaySalary() {  
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {  
    emp1 := Employee {
        name:     "Sam Adolf",
        salary:   5000,
        currency: "$",
    }
    emp1.displaySalary() //Calling displaySalary() method of Employee type
}
```
[在 Playground 中运行](https://play.golang.org/p/3khFtFJdbee)  

上面程序的第 6 行，我们为 `Employee` 创建了一个名为 `displaySalary` 的方法。在 `displaySalary()` 方法内部可以访问它的接收者 `e` （类型为 `Employee`）。在第 17 行，我们使用接收者 `e`，并打印它的 `name`，`currency` 以及 `salary`。  

程序的输出为:   

```golang
Salary of Sam Adolf is $5000 
```

## 为什么使用方法而不是函数？  

上面的程序可以使用函数而不是方法重写如下：  

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    name     string
    salary   int
    currency string
}

/*
 displaySalary() method converted to function with Employee as parameter
*/
func displaySalary(e Employee) {  
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {  
    emp1 := Employee{
        name:     "Sam Adolf",
        salary:   5000,
        currency: "$",
    }
    displaySalary(emp1)
}
```

[在 Playground 中运行](https://play.golang.org/p/xKqdal-DqHT)   

在上面的程序中，我们使用 `displaySalary` 函数替换了方法，并将 `Employee` 结构体作为参数传给它。该程序的输出与上面的程序输出一样：`Salary of Sam Adolf is $5000`。  

那么为什么我们可以用函数完成同样的工作，却还要使用方法呢？这里有几个原因，我们一个一个地看。  

* [Go 不是一个纯面向对象的编程语言](https://golang.org/doc/faq#Is_Go_an_object-oriented_language)，它不支持 `class` 类型。因此通过在一个类型上建立方法来实现与 `class` 相似的行为。  
* 同名方法可以定义在不同的类型上，但是 Go 不允许同名函数。假设我们有一个 `Square` 和 `Circle` 两个结构体。在 `Square` 和 `Circle` 上定义同名的方法是合法的，比如下面的程序：     

```golang
package main

import (  
    "fmt"
    "math"
)

type Rectangle struct {  
    length int
    width  int
}

type Circle struct {  
    radius float64
}

func (r Rectangle) Area() int {  
    return r.length * r.width
}

func (c Circle) Area() float64 {  
    return math.Pi * c.radius * c.radius
}

func main() {  
    r := Rectangle{
        length: 10,
        width:  5,
    }
    fmt.Printf("Area of rectangle %d\n", r.Area())
    c := Circle{
        radius: 12,
    }
    fmt.Printf("Area of circle %f", c.Area())
}
```

程序的输出为：  
```golang
Area of rectangle 50  
Area of circle 452.389342  
```

接口正是应用了这一点（译者注：*相同的方法名可以用在不同的接收者类型上*）。我们将在下面的教程中讨论接口的细节。  

##  指针接收者 vs. 值接收者  

目前为止我们看到的都是以值作为接收者。以指针为接收者也是可以的。两者的区别在于，以指针作为接收者时，方法内部对其的修改对于调用者是可见的，但是以值作为接收者却不是。让我们通过下面的程序帮助理解。   

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    name string
    age  int
}

/*
Method with value receiver  
*/
func (e Employee) changeName(newName string) {  
    e.name = newName
}

/*
Method with pointer receiver  
*/
func (e *Employee) changeAge(newAge int) {  
    e.age = newAge
}

func main() {  
    e := Employee{
        name: "Mark Andrew",
        age:  50,
    }
    fmt.Printf("Employee name before change: %s", e.name)
    e.changeName("Michael Andrew")
    fmt.Printf("\nEmployee name after change: %s", e.name)

    fmt.Printf("\n\nEmployee age before change: %d", e.age)
    (&e).changeAge(51)
    fmt.Printf("\nEmployee age after change: %d", e.age)
}
```

上面的程序中， `changeName` 方法有一个值接收者 (`e Employee`)，而 `changeAge` 方法有一个指针接收者 (`e *Employee`)。在 `changeName` 中改变 `Employee` 的字段 `name` 的值对调用者而言是不可见的，因此程序在调用 `e.changeName("Michael Andrew")` 方法之前和之后，打印的 `name` 是一致的。而因为 `changeAge` 的接受者是一个指针 (`e *Employee`)，因此通过调用方法 `(&e).changeAge(51)` 来修改 `age` 对于调用者是可见的。 程序的输出如下：  

```golang
Employee name before change: Mark Andrew  
Employee name after change: Mark Andrew

Employee age before change: 50  
Employee age after change: 51
```

在上面的程序第36行，我们用 `(&e).changeAge(51)` 来调用 `changeAge` 方法。因为 `changeAge` 有一个指针类型的接收者我们必须使用 `(&e)` 来调用。这不是必须的，Go允许我们省略 `&` 符号，因此可以只写为 `e.changeAge(51)`。Go 将 `e.changeAge(51)` 解析为 `(&e).changeAge(51)`。  

下面的程序使用 `e.changeAge(51)` 来替代 `(&e).changeAge(51)`。它与上面的程序的打印结果是一样的。  

程序的输出为：  

```golang
package main

import (  
    "fmt"
)

type Employee struct {  
    name string
    age  int
}

/*
Method with value receiver  
*/
func (e Employee) changeName(newName string) {  
    e.name = newName
}

/*
Method with pointer receiver  
*/
func (e *Employee) changeAge(newAge int) {  
    e.age = newAge
}

func main() {  
    e := Employee{
        name: "Mark Andrew",
        age:  50,
    }
    fmt.Printf("Employee name before change: %s", e.name)
    e.changeName("Michael Andrew")
    fmt.Printf("\nEmployee name after change: %s", e.name)

    fmt.Printf("\n\nEmployee age before change: %d", e.age)
    e.changeAge(51)
    fmt.Printf("\nEmployee age after change: %d", e.age)
}  
```

## 何时使用指针接收者，何时使用值接收者？  

一般来讲，指针接收者可用于对接收者的修改应该对调用者可以见的场合。  

指针接收者也可用于拷贝结构体代价较大的场合。考虑一个包含较多字段的结构体，若使用值作为接收者则必须拷贝整个结构体，这样的代价很大。这种情况下使用指针接收者将避免结构体的拷贝，而仅仅是指向结构体指针的拷贝。  

其他情况下可以使用值接收者。  

## 匿名字段函数  

匿名字段的方法可以被包含该匿名字段的结构体的变量调用，就好像该匿名字段的方法属于包含该字段的结构体一样。  

```golang
package main

import (  
    "fmt"
)

type address struct {  
    city  string
    state string
}

func (a address) fullAddress() {  
    fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {  
    firstName string
    lastName  string
    address
}

func main() {  
    p := person{
        firstName: "Elon",
        lastName:  "Musk",
        address: address {
            city:  "Los Angeles",
            state: "California",
        },
    }
    p.fullAddress() //accessing fullAddress method of address struct
}
```

在上面的程序中，第32行，我们通过 `p.fullAddress()` 调用了 `address` 的方法 `fullAddress()`。像 `p.address.fullAddress()` 这样的直接调用是不必要的。程序的输出为：  

```golang
Full address: Los Angeles, California 
```

##  方法的值接收者 vs. 函数的值参数  

这是很多新手遇到的问题。我们将尽可能把它说明白。  

当一个函数有一个值参数时，它只接受一个值参数。  

当一个方法有一个值接收者时，它可以接受值和指针接收者。  

让我们通过程序说明这一点。  

```golang
package main

import (  
    "fmt"
)

type rectangle struct {  
    length int
    width  int
}

func area(r rectangle) {  
    fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {  
    fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func main() {  
    r := rectangle{
        length: 10,
        width:  5,
    }
    area(r)
    r.area()

    p := &r
    /*
       compilation error, cannot use p (type *rectangle) as type rectangle 
       in argument to area  
    */
    //area(p)

    p.area()//calling value receiver with a pointer
}
```

第12行，函数 `func area(r rectangle)` 接受一个值参数，而方法 `func (r rectangle) area()` 接受一个值接收者。

在第25行，我们传递了一个值来调用 `area` 函数 `area(r)`，它将工作。同样地，我们通过值接收者调用 `area` 方法 `r.area()` 它也可以工作。

在第28行，我们创建了一个指向 `r` 的指针 `p`。如果我们试图将这个指针传递给只接受值的 area 函数那么编译器将报错：`compilation error, cannot use p (type *rectangle) as type rectangle in argument to area`.。这是我们预期的。

现在来到了微妙的地方，第35行 `p.area()` 使用指针接收者 `p` 调用了接受一个值接收者的方法 `area` 。这是完全合法的。原因是对于 `p.area()`，Go 将其解析为 `(&p).area()`，因为 `area` 方法必须接受一个值接收者。这是为了方便 Go 给我们提供的语法糖。

程序的输出为：  

```golang
Area Function result: 50  
Area Method result: 50  
Area Method result: 50 
```

## 方法的指针接收者 vs. 函数的指针参数  

与值参数相似，一个接受指针参数的函数只能接受指针，而一个接收者为指针的方法可以接受值接收者和指针接收者。   

```golang
package main

import (  
    "fmt"
)

type rectangle struct {  
    length int
    width  int
}

func perimeter(r *rectangle) {  
    fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {  
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func main() {  
    r := rectangle{
        length: 10,
        width:  5,
    }
    p := &r //pointer to r
    perimeter(p)
    p.perimeter()

    /*
        cannot use r (type rectangle) as type *rectangle in argument to perimeter
    */
    //perimeter(r)

    r.perimeter()//calling pointer receiver with a value
}
```

在上面的程序中，第12行定义了一个函数 `perimeter`，该函数接受一个指针作为参数，而17行定义了一个方法，有一个指针接收者。  

27行我们通过指针参数调用 `perimeter` 函数，在第28行我们通过一个指针接收者调用 `perimeter` 方法。一切都好。  

在被注释掉的第33行，我们试图通以一个值 `r` 调用 `perimeter` 函数。这是非法的，因为一个接受指针为参数的函数不能接受一个值作为参数。如果去掉注释运行程序，则编译将报错：`main.go:33: cannot use r (type rectangle) as type *rectangle in argument to perimeter.`。  

在第35行我们通过一个值接收者 `r` 调用接受一个指针接收者的 `perimeter` 方法。这是合法的，`r.perimeter()` 这一行将被 Go 解析为 `(&r).perimeter()`。这是为了方便 Go 给我们提供的语法糖。程序的输出为：  

```golang
perimeter function output: 30  
perimeter method output: 30  
perimeter method output: 30  
```

## 定义非结构体类型的方法  

现在我们定义的都是结构体类型的方法。同样可以定义非结构体类型的方法，不过这里需要注意一点。为了定义某个类型的方法，接收者类型的定义与方法的定义必须在同一个包中。目前为止，我们定义的结构体和相应的方法都是在`main`包中的，因此没有任何问题。   

```golang
package main

func (a int) add(b int) {  
}

func main() {

}
```

在上面的程序中，第3行我们试图添加一个方法 `add` 给内置类型 `int`。这是不允许的，因为定义方法 add 所在的包和定义类型 `int` 的包不是同一个包。这个程序将会报编译错误：`cannot define new methods on non-local type int`。  

使其工作的方法为定义内置类型的别名，然后以这个新类型为接收者创建方法。   

```golang
package main

import "fmt"

type myInt int

func (a myInt) add(b myInt) myInt {  
    return a + b
}

func main() {  
    num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)
}
```

上面的程序中，第5行，我们创建了新的类型，一个 `int` 的别名 `myInt`。在第7行，我们定义了一个方法 add，以 `myInt` 作为接收者。  

程序的输出为： `Sum is 15`。  

我已经创建了一个程序，其中包含了目前为止所讨论的所有概念，可以在 github 上找到它。  

希望你喜欢阅读。请留下宝贵的意见和反馈:)  