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

var clients *list.List //global variable, so that it can be used in all functions

type client struct{
    conn net.Conn
    name string
}

func main(){
    //mark_todo : get default configuration from config file
    port := flag.String("p","9002","the port that you want to listen");
    host := flag.String("h","","the host/ip that you want to listen");
    flag.Parse()
    clients = list.New()
    ln,err := net.Listen("tcp",*host+":"+*port)
    e(err) 
    for {
        conn, err := ln.Accept();
        e(err) //error
        go response(conn);
    }
}

func response(conn net.Conn){

    //init variable
    var c client;
    var content string
    data := make([]byte,1024);

    //save the client into list
    c.conn = conn
    c.name = getName(conn)
    elem := clients.PushBack(c)
    msg := "new conn from "+conn.RemoteAddr().String()+" name:"+c.name

    defer func(){
        clients.Remove(elem)
        conn.Close()
    }()

    //listen and broadcast
    forBreak:
    for {
        data = make([]byte,1024);//mark review
        n,err := conn.Read(data);
        switch err {
            case nil:
                broadCast(data);
            case io.EOF:
                broadCast(data);
                break forBreak;
            default:
                break forBreak;
        }
        d := data[:n]
        content = strings.TrimRight(string(d)," \n\r")
        if content == "bye" {
           break forBreak 
        }
    }
    broadCast(c.name" from "+conn.RemoteAddr().String()+" left");
}
func broadCast(msg []byte) {
    log.Println(string(msg))
    return 
}

func getName(conn net.Conn) (name string){
    data := make([]byte,1024);
    _,err := conn.Write([]byte{'y','o','u','r',' ','n','a','m','e',':'});
    e(err)
    n,err := conn.Read(data);
    e(err)
    name = string(data[0:n])
    for ;existsName(name); {
        _,err = conn.Write([]byte{'n','a','m','e',' ','e','x','i','s','t','s','!','y','o','u','r',' ','n','a','m','e',':'});
        e(err)
        n,err = conn.Read(data);
        e(err)
        name = string(data[0:n])
    }

    return name
}

func existsName(name string) (r bool) {
    //mark_todo 
    //return true
    return false
}

func e(err error){
    if err != nil {
        log.Fatal("error:"+err.Error());
        os.Exit(-1);
    }
}
