package main

import (
    "fmt"
    //"flag"
    "os"
    "io"
)


func main(){
    len := len(os.Args)
    if len < 2{
        l("you need specify a filename");
        os.Exit(0);
    }
    for i:=1;i<len;i++{
        filename  := os.Args[i]
        catFile(filename)
    }
}

func catFile(filename string){
    file, err := os.Open(filename)
    if err != nil {
        l("cat: "+filename+": no such file")
        return
    }
    defer file.Close()
    data := make([]byte, 100)
    for  {
        count, err := file.Read(data)
        fmt.Printf("%s\n",data[:count]);
        if err == io.EOF  {
            return
        }else if err != nil {
            l("cat:read "+filename+" failed");
            return 
        }
    }
}

func l(s string){
    fmt.Println(s);
}
