package main

import (
    "fmt"
    "os"
)

func main(){
    len := len(os.Args)
    for i:=1;i<len;i++{
        fmt.Printf("%s ",os.Args[i])
    }
    fmt.Println("")
}
