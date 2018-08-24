
18-1 - 接口1  
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
