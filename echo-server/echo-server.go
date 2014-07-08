package main

import (
    "fmt"
    "net"
    "flag"
    "io"
    "os"
    "strings"
)

func main(){
    port := flag.String("p","9002","the port that you want to listen");
    host := flag.String("h","","the host/ip that you want to listen");
    flag.Parse()
    ln,err := net.Listen("tcp",*host+":"+*port)
    e(err) // error
    for {
        conn, err := ln.Accept();
        e(err) //error
        go response(conn);
    }
}

func response(conn net.Conn){
    data := make([]byte,1024);
    var content string
    fmt.Println("new conn!")
    forBreak:
    for {
        data = make([]byte,1024);//mark review
        n,err := conn.Read(data);
        switch err {
            case nil:
                fmt.Printf("recived:%s",string(data))
                _,_ = conn.Write(data);
            case io.EOF:
                fmt.Printf("recived:%s --end",string(data))
                _,_ = conn.Write(data);
                break forBreak;
            default:
                fmt.Println("rev failed:"+err.Error())
                break forBreak;
        }
        d := data[:n]
        content = strings.TrimRight(string(d)," \n\r")
        if content == "bye" {
           break forBreak 
        }
    }
    //_,_ = conn.Write("bye")
    fmt.Println("conn closed")
    conn.Close();
}
func e(err error){
    if err != nil {
        fmt.Println("error:"+err.Error());
        os.Exit(-1);
    }
}

