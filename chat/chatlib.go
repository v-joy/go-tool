package chatlib

import (
    "fmt"
    "os"
    "net"
)

type Msg struct{
   action string
   data string
}

func (this *Msg) Send(){
    fmt.Println("chatlib.Send");
}

func (this *Msg) Recive(){

}

