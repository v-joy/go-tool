package main

import (
	"flag"
	"fmt"
	. "go-tool/chat"
	"log"
	"net"
	"os"
	"sync"
)

type Client struct {
	Conn net.Conn
	Wt   *Writer
	Rd   *Reader
	Lwt  *Writer
	Lrd  *Reader
	Name string
}

func main() {
	port := flag.String("p", SERVER_PORT, "the port that you want to connect")
	host := flag.String("h", SERVER_IP, "the host/ip that you want to connect")
	flag.Parse()
	wg := new(sync.WaitGroup)
	wg.Add(2)
	clt, err := NewClient(*host, *port)
	e(err)
	go func() {
		defer wg.Done()
		clt.handleListen()
	}()
	go func() {
		defer wg.Done()
		clt.handleWrite()
	}()
	wg.Wait()
}

func NewClient(host, port string) (clt *Client, err error) {
	conn, err := net.Dial("tcp", host+":"+port)
	e(err)
	wt := NewWriter(conn)
	rd := NewReader(conn)
	lwt := NewWriter(os.Stdout)
	lrd := NewReader(os.Stdin)
	clt = &Client{
		Wt:   wt,
		Rd:   rd,
		Lwt:  lwt,
		Lrd:  lrd,
		Name: CLIENT_INIT_NAME,
		Conn: conn,
	}
	err = clt.Register()
	return clt, err
}

func (this *Client) Register() (err error) {
	if this.Name != CLIENT_INIT_NAME {
		return nil
	}
	this.output("please input your name")
	for {
		name, _, _ := this.Lrd.ReadLine()
		msg := map[string]interface{}{
			"action": "register",
			"data": map[string]string{
				"name": string(name),
			},
		}
		err := this.Wt.WriteMsg(msg)
		e(err)
		ret, err := this.Rd.ReadMsg()
		e(err)
		action, data, err := MsgInfo(ret)
		e(err)
		switch action {
		case "ack":
			if data["status"].(float64) == 200 {
				this.Name = string(name)
				this.output("set name success :" + string(name) + "! now , let's chat!")
				return nil
			}
		case "chat":
			// this is never supposed to happen
			this.output("this is never supposed to happen")
		default:
			this.output("action type error:" + action)
			os.Exit(1)
		}
		this.output(" name exists: please input your name again ")
	}
	return nil
}

func (this *Client) handleListen() {
	for {
		msg, err := this.Rd.ReadMsg()
		e(err)
		action, data, err := MsgInfo(msg)
		e(err)
		switch action {
		case "ack":
		case "chat":
			this.output(data["msg"].(string), data["name"].(string))
		default:
			this.output("action type error:" + action)
			os.Exit(1)
		}
	}

}

func (this *Client) handleWrite() {
	for {
		line, _, _ := this.Lrd.ReadLine()
		msg := map[string]interface{}{
			"action": "broadcast",
			"data": map[string]string{
				"msg": string(line),
			},
		}
		err := this.Wt.WriteMsg(msg)
		e(err)
	}
}

func (this *Client) output(str ...string) {
	var name string
	if len(str) == 1 {
		name = SYSTEM_NAME
	} else {
		name = str[1]
	}
	this.Lwt.WriteString(name + " : " + str[0] + "\n")
	this.Lwt.Flush()
}

func e(err error) {
	if err != nil {
		fmt.Println("error:" + err.Error())
		os.Exit(-1)
	}
}

func l(msg string) {
	log.Println(msg)
}
