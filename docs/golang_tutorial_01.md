1 - 介绍和安装
========================

简介：本课程非常重要。  

下一节：[第二篇 Hello World](/docs/golang_tutorial_02.md)   

## 什么是Golang  

Go也被称为Golang，它是一种开源的，由Google创建的编译和静态类型编程语言。  

Golang的主要重点是使高度可用和可扩展的网络应用程序的开发变得简单和容易。  

## 为什么选择Golang   

为什么你会选择Golang作为你的服务端编程语言，当有大量的其他语言，如python，ruby，nodejs ......可以完成同样的工作。  

## 以下是我在选择Go时找到的一些优点  

并发是语言的固有部分。因此编写多线程程序是一块蛋糕。这是通过Goroutines和渠道实现的，我们将在稍后的教程中讨论。  

Golang是一种编译语言。源代码被编译为本机二进制文件。解释型语言（如用于nodejs的JavaScript）缺少这一点。  

语言规范非常简单。整个规范适合一个页面，你甚至可以用它来编写你自己的编译器:)    

go编译器支持静态链接。所有go代码都可以静态链接到一个大型的二进制文件中，并且可以轻松部署到云服务器中，而无需担心依赖关系。  

## 安装  

Golang在所有三种平台的Mac，Windows和Linux上均受支持。您可以从https://golang.org/dl/下载相应平台的二进制文件  

## 苹果系统  

从https://golang.org/dl/下载OS X安装程序。双击即可开始安装。按照提示进行操作，这应该将Golang安装在 `/usr/local/go`中，并且还将文件夹`/usr/local/go/bin`添加到PATH环境变量中。    

## Windows 系统   

从https://golang.org/dl/下载MSI安装程序。双击开始安装并按照提示进行操作。这将在位置c：\ Go中安装Golang，并将目录c：\ Go \ bin添加到您的路径环境变量中。   

## Linux 系统  

从https://golang.org/dl/下载tar文件并将其解压缩到/ `/usr/local/` 。将`/usr/local/go/bin`添加到PATH环境变量中。这应该安装在Linux中。    

本文由 [GCTT](https://github.com/studygolang/GCTT) 原创翻译，[Go语言中文网](https://studygolang.com/)首发。