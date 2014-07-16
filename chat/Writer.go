package chat

import (
    "bufio"
    "io"
    "encoding/json"
    "log"
)

type Writer struct{
   bufio.Writer 
}

func NewWriter(rd io.Writer) *Writer {
    bufrd := bufio.NewWriter(rd)
    cr := Writer{*bufrd}
   return &cr
}


func NewWriterSize(rd io.Writer,size int) *Writer {
    bufrd := bufio.NewWriterSize(rd,size)
    cr := Writer{*bufrd}
   return &cr
}

//witch type should be returned
func ( this *Writer) WriteMsg( data interface{})(n int,err error){
    body,err := json.Marshal(data)
    if err != nil{
        log.Println("error:",err)
        return 0,err
    }
    n = len(body)
    //write head
    //write body
    return  n,nil
}


