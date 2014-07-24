package main

import (
	"container/list"
	"flag"
	. "go-tool/chat"
	//"io"
	"log"
	"net"
	"sync"
	//"os"
	//"strings"
)

var clients *list.List //global variable, so that it can be used in all functions

type Client struct {
	Conn net.Conn
	Wt   *Writer
	Rd   *Reader
	Name string
}

type server struct {
	Port string
	Host string
}

func main() {
	port := flag.String("p", SERVER_PORT, "the port that you want to listen")
	host := flag.String("h", SERVER_IP, "the host/ip that you want to listen")
	flag.Parse()
	clients = list.New()
	ln, err := net.Listen("tcp", *host+":"+*port)
	l("listenning:" + *host + ":" + *port)
	e(err)
	wg := new(sync.WaitGroup)
	for {
		conn, err := ln.Accept()
		wg.Add(1)
		e(err) //error
		clt := getClient(conn)
		ele := clients.PushBack(clt)
		go func() {
			defer func() {
				name := clt.Name
				clients.Remove(ele)
				clt.distroy()
				broadCast(name+" has left the chat room!", SYSTEM_NAME)
				wg.Done()
			}()
			clt.response()
		}()
	}
	wg.Wait()
}

func getClient(conn net.Conn) (clt *Client) {
	wt := NewWriter(conn)
	rd := NewReader(conn)
	clt = &Client{
		Wt:   wt,
		Rd:   rd,
		Conn: conn,
		Name: CLIENT_INIT_NAME,
	}
	return clt
}

func (this *Client) distroy() {
	//mark todo
	l(this.Name + " is dead")
}

func (c *Client) response() {
	//listen and broadcast
	reply := map[string]interface{}{}
	for {
		msg, err := c.Rd.ReadMsg()
		if err != nil {
			log.Println("error:" + err.Error())
			return
		}
		action, data, err := MsgInfo(msg)
		if err != nil {
			//reply error and close the connection
			l("error:" + err.Error())
			return
		}
		reply = map[string]interface{}{
			"action": "ack",
			"data": map[string]interface{}{
				"status": 200,
				"msg":    "ok",
			},
		}
		switch action {
		case "ack":
			continue
		case "broadcast":
			broadCast(data["msg"].(string), c.Name)
		case "register":
			name, ok := data["name"].(string)
			if !ok || existsName(name) {
				reply = map[string]interface{}{
					"action": "ack",
					"data": map[string]interface{}{
						"status": 400,
						"msg":    "name " + name + " exists",
					},
				}
				l("name " + name + " exists")
			} else {
				c.Name = name
				broadCast("--- new member is comming: "+c.Name+" from "+c.Conn.RemoteAddr().String()+" --- ", name)
				l("new user :" + c.Name + " from " + c.Conn.RemoteAddr().String())
			}
		default:
			//show error
			l("invalid action:" + action)
			//and close the connection
			return
		}
		err = c.Wt.WriteMsg(reply)
		e(err)
	}
}
func broadCast(info, name string) {
	msg := map[string]interface{}{
		"action": "chat",
		"data": map[string]interface{}{
			"name": name,
			"msg":  info,
		},
	}
	for e := clients.Front(); e != nil; e = e.Next() {
		c := e.Value.(*Client)
		if c.Name != name {
			c.Wt.WriteMsg(msg)
		}
	}
	return
}

func existsName(name string) (r bool) {
	for e := clients.Front(); e != nil; e = e.Next() {
		c := e.Value.(*Client)
		if c.Name == name {
			return true
		}
	}
	if name == SYSTEM_NAME || name == CLIENT_INIT_NAME {
		return true
	}
	return false
}

func l(msg string) {
	log.Println(msg)
}

func e(err error) {
	if err != nil {
		log.Fatal("error:" + err.Error())
		//os.Exit(-1)
	}
}
