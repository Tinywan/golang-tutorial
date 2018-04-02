package main

import (
    "fmt"
    "os"
    "os/exec"
    "github.com/Tinywan/golang-tutorial/stringutil"
)    

func main() {
    //fmt.Println("Hello golang-tutorial!")
    fmt.Printf(stringutil.Reverse("!oG ,olleH"))
    cmd := exec.Command("pwd")
    stdin,err :=cmd.StdinPipe();
    if err != nil {
        fmt.Println(err)
    }
    _,err = stdin.Write([]byte("tmp.txt"))
    if err != nil {
        fmt.Println(err)
    }
    stdin.Close()
    cmd.Stdout = os.Stdout  //终端标准输出tmp.txt
    cmd.Start()
}