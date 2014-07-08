package main

import (
    "fmt"
    "os"
    "bufio"
)

func main(){
    reader := bufio.NewReader(os.Stdin);
    len := len(os.Args)
    /*if len < 2{
        l("you need specify at least one filename");
        os.Exit(0);
    }*/
    fds := make([]*os.File,len)
    for i:=1;i<len;i++{
        filename  := os.Args[i]
        file, err := os.OpenFile(filename,os.O_CREATE | os.O_WRONLY,0666)
        if err != nil {
            l("tee: can not open:"+err.Error())
            fds[i-1] = nil
            continue 
        }else{
            fds[i-1] = file
        }
    }
    defer func(){
        for i:=0;i<len-1;i++{
            if fds[i] != nil {
                fds[i].Close()
            }
        }
    }()
    for  {
        str,err := reader.ReadString('\n');
        if err != nil {
            break
        }
        fmt.Printf("%s",str);
        for i:=0;i<len-1;i++{
            if fds[i] != nil {
                _,err := fds[i].WriteString(str)
                if err != nil {
                    l("tee: fail to write:"+err.Error())
                }
            }
        }

    }
}

func l(s string){
    fmt.Println(s);
}
