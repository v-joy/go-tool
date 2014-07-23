package main

import (
    "fmt"
    "net"
    "flag"
    "os"
    "log"
	. "go-tool/chat"
)

type Client struct {
	Conn net.Conn
	Wt *Writer
    Rd *Reader
    Lwt *Writer
    Lrd *Reader
    Name string
}

func NewClient(host,port string)( clt *Client,err error) {
    conn, err := net.Dial("tcp", host+":"+port)
    e(err) 
    wt := NewWriter(conn)
    rd := NewReader(conn)
    lwt := NewWriter(os.Stdout)
    lrd := NewReader(os.Stdin)
    clt = &Client{
        Wt:wt,
        Rd:rd,
        Lwt:lwt,
        Lrd:lrd,
        Name:CLIENT_INIT_NAME,
        Conn:conn,
    }
    err = clt.Register()
    return clt,err
}


func (this *Client)Register() (err error){
    if this.Name != CLIENT_INIT_NAME {
        return nil
    }
    log.Println("register begin")
    this.Lwt.WriteString("please input your name");
    for{
        name,_,_:= this.Lrd.ReadLine()
        data := map[string]string{
            "name":string(name),
        }
        msg := Msg{
            Action:"register",
            Data:data,
        }
        this.Wt.WriteMsg(msg)
        ret,_ := this.Rd.ReadMsg()

        switch ret.Data["status"]{
            case "200":
                this.Name = string(name)
                return nil
            default:
                continue
        }
    }
    return nil
}




func main(){
    port := flag.String("p",SERVER_PORT,"the port that you want to connect");
    host := flag.String("h",SERVER_IP,"the host/ip that you want to connect");
    flag.Parse()
    clt,err := NewClient(*host,*port)
    e(err)
    go clt.handleListen()
    go clt.handleWrite()
}


func (this *Client)handleListen(){
    
}

func (this *Client)handleWrite(){
    
}

func (this *Client)output(){
    
}


func e(err error){
    if err != nil {
        fmt.Println("error:"+err.Error());
        os.Exit(-1);
    }
}

