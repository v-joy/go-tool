package chat

import (
	"bufio"
	"encoding/json"
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

//witch type should be returned
func (this *Reader) ReadMsg() (data interface{}, err error) {
	//read head
	//read body
	body := []byte(`{“action”:”register”, “data”:{“name”:”xxx”}}`)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("error:", err)
		return 0, err
	}
	return data, nil
}
