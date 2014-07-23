package main

import (
	"container/list"
	"flag"
	. "go-tool/chat"
	"io"
	"log"
	"net"
	"os"
	"strings"
)


var clients *list.List //global variable, so that it can be used in all functions

type Client struct {
	Conn net.Conn
	Wt *Writer
    Rd *Reader
    Name string
}

type server struct {
	Port string
	Host string
}

func main() {
	//mark_todo : get default configuration from config file
	port := flag.String("p", SERVER_PORT, "the port that you want to listen")
	host := flag.String("h", SERVER_IP, "the host/ip that you want to listen")
	flag.Parse()
	clients = list.New()
	ln, err := net.Listen("tcp", *host+":"+*port)
	log.Println("listenning:" + *host + ":" + *port)
	e(err)
	for {
		conn, err := ln.Accept()
		e(err) //error
		clt := getClient(conn)
        ele := clients.PushBack(clt)
	    defer func() {
		    clients.Remove(ele)
            clt.distroy()
	    }()
        go clt.response()
	}
}

func getClient(conn net.Conn)(clt *Client){
    wt := NewWriter(conn)
    rd := NewReader(conn)
    clt = &Client{
        Wt:wt,
        Rd:rd,
        Conn:conn,
        Name:CLIENT_INIT_NAME,
    }
    l("new conn from " + conn.RemoteAddr().String())
    return clt
}

func (this *Client)distroy(){
    //mark todo 
}


func (c *Client)response() {
	//mark_todo : check version
	//broadCast(msg, SYSTEM_NAME)

	//listen and broadcast
/*
forBreak:
	for {
		msg, err := c.Rd.ReadMsg()
		d := data[:n]
		content = strings.TrimRight(string(d), "\n\r")
		switch err {
		case nil:
			broadCast(content, c.Name)
		case io.EOF:
			broadCast(content, c.Name)
			break forBreak
		default:
			break forBreak
		}
		if content == "bye" {
			break forBreak
		}
	}
	broadCast(c.Name+" from "+conn.RemoteAddr().String()+" left", SYSTEM_NAME)
*/
}
func broadCast(msg, name string) {
	msgFmt := name + ":" + msg + "\r\n"
	log.Print(msgFmt)
	for e := clients.Front(); e != nil; e = e.Next() {
		c := e.Value.(Client)
		if c.Name != name {
			c.Conn.Write([]byte(msgFmt))
		}
	}
	return
}


func getName(conn net.Conn) (name string) {
	data := make([]byte, 1024)
	_, err := conn.Write([]byte{'y', 'o', 'u', 'r', ' ', 'n', 'a', 'm', 'e', ':'})
	e(err)
	n, err := conn.Read(data)
	e(err)
	name = string(data[0 : n-2])
	for existsName(name) {
		_, err = conn.Write([]byte{'n', 'a', 'm', 'e', ' ', 'e', 'x', 'i', 's', 't', 's', '!', 'y', 'o', 'u', 'r', ' ', 'n', 'a', 'm', 'e', ':'})
		e(err)
		n, err = conn.Read(data)
		e(err)
		name = string(data[0 : n-2])
	}

	return name
}

func existsName(name string) (r bool) {
	for e := clients.Front(); e != nil; e = e.Next() {
		c := e.Value.(Client)
		if c.Name == name {
			return true
		}
	}
	if name == SYSTEM_NAME {
		return true
	}
	return false
}

func l(mes string){
    log.Println(msg)
}

func e(err error) {
	if err != nil {
		log.Fatal("error:" + err.Error())
		os.Exit(-1)
	}
}

