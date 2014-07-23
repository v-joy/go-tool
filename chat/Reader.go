package chat

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"log"
)

type Reader struct {
	bufio.Reader
}

func NewReader(rd io.Reader) *Reader {
	bufrd := bufio.NewReader(rd)
	cr := Reader{*bufrd}
	return &cr
}

func NewReaderSize(rd io.Reader, size int) *Reader {
	bufrd := bufio.NewReaderSize(rd, size)
	cr := Reader{*bufrd}
	return &cr
}

func (this *Reader) ReadMsg() (data interface{}, err error) {
	//read head
	var bodyLen16 uint16
	err = binary.Read(this, binary.BigEndian, &bodyLen16)
	if err != nil {
		log.Println("error:", err)
		return nil, errors.New("read header error")
	}
	bodyLen := int(bodyLen16)
	//read body
	body := make([]byte, bodyLen)
	_, err = io.ReadFull(this, body)
	if err != nil {
		log.Println("error:", err)
		return nil, errors.New("read body error")
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("error:", err)
		return nil, errors.New("convert body error")
	}
	return data, nil
}
