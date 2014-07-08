package main

import (
    "log"
    "net"
    "flag"
    "io"
    "os"
    "strings"
    "container/list"
)

const SYSTEM_NAME = "SYSTEM"

var clients *list.List //global variable, so that it can be used in all functions

type client struct{
    Conn net.Conn
    Name string
}

func main(){
    //mark_todo : get default configuration from config file
    port := flag.String("p","9002","the port that you want to listen");
    host := flag.String("h","","the host/ip that you want to listen");
    flag.Parse()
    clients = list.New()
    ln,err := net.Listen("tcp",*host+":"+*port)
    log.Println("listenning:"+*host+":"+*port);
    e(err) 
    for {
        conn, err := ln.Accept();
        e(err) //error
        go response(conn);
    }
}

func response(conn net.Conn){
    //mark_todo : check version
    
    //init variable
    var c client;
    var content string
    data := make([]byte,1024);

    //save the client into list
    c.Conn = conn
    c.Name = getName(conn)
    elem := clients.PushBack(c)
    msg := "new conn from "+conn.RemoteAddr().String()+" name:"+c.Name
    broadCast(msg,SYSTEM_NAME);
    defer func(){
        clients.Remove(elem)
        conn.Close()
    }()

    //listen and broadcast
    forBreak:
    for {
        data = make([]byte,1024);//mark review
        n,err := conn.Read(data);
        d := data[:n]
        content = strings.TrimRight(string(d),"\n\r")
        switch err {
            case nil:
                broadCast(content,c.Name);
            case io.EOF:
                broadCast(content,c.Name);
                break forBreak;
            default:
                break forBreak;
        }
        if content == "bye" {
           break forBreak 
        }
    }
    broadCast(c.Name+" from "+conn.RemoteAddr().String()+" left",SYSTEM_NAME);
}
func broadCast(msg,name string) {
    msgFmt := name+":"+msg+"\r\n";
    log.Print(msgFmt);
    for e := clients.Front();e!=nil;e=e.Next(){
        c := e.Value.(client)
        if c.Name != name {
            c.Conn.Write([]byte(msgFmt));
        }
    }
    return 
}

func getName(conn net.Conn) (name string){
    data := make([]byte,1024);
    _,err := conn.Write([]byte{'y','o','u','r',' ','n','a','m','e',':'});
    e(err)
    n,err := conn.Read(data);
    e(err)
    name = string(data[0:n-2])
    for ;existsName(name); {
        _,err = conn.Write([]byte{'n','a','m','e',' ','e','x','i','s','t','s','!','y','o','u','r',' ','n','a','m','e',':'});
        e(err)
        n,err = conn.Read(data);
        e(err)
        name = string(data[0:n-2])
    }

    return name
}

func existsName(name string) (r bool) {
    for e := clients.Front();e!=nil;e=e.Next(){
        c := e.Value.(client)
        if c.Name == name {
            return true
        }
    }
    if name == SYSTEM_NAME {
        return true
    }
    return false
}

func e(err error){
    if err != nil {
        log.Fatal("error:"+err.Error());
        os.Exit(-1);
    }
}
