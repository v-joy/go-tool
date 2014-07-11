package main

import (
    "fmt"
    "net"
    "flag"
    "io"
    "os"
    "strings"
)

const (
    SERVER_IP = "127.0.0.1"
    SERVER_PORT = "9002"
)

func main(){
    port := flag.String("p",SERVER_PORT,"the port that you want to connect");
    host := flag.String("h",SERVER_IP,"the host/ip that you want to connect");
    flag.Parse()
    conn, err := net.Dial("tcp", *host+":"+*port)
    e(err) 
    register(&conn)
    go handleListen(&conn)
    go handelWtire(&conn)
}

func register(*net.Conn conn){
    
}

func handleListen(*net.Conn conn){
    
}

func handleWrite(*net.Conn conn){
    
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

