package chat

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"io"
	"log"
)

type Writer struct {
	bufio.Writer
}

func NewWriter(rd io.Writer) *Writer {
	bufrd := bufio.NewWriter(rd)
	cr := Writer{*bufrd}
	return &cr
}

func NewWriterSize(rd io.Writer, size int) *Writer {
	bufrd := bufio.NewWriterSize(rd, size)
	cr := Writer{*bufrd}
	return &cr
}

func (this *Writer) WriteMsg(data interface{}) (err error) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Println("error:", err)
		return err
	}
	var bodyLen uint16
	bodyLen = uint16(len(body))
	//write head
	err = binary.Write(this, binary.BigEndian, bodyLen)
	if err != nil {
		log.Println("binary.Write failed:", err)
		return err
	}
	//write body  mark: need handle error
	_, err = this.Write(body)
	if err != nil {
		log.Println("bufio.Write failed:", err)
		return err
	}
	this.Flush()
	return nil
}
