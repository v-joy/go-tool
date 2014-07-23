package chat

import (
	"errors"
	"log"
)

func MsgInfo(msg interface{}) (action string, data map[string]interface{}, err error) {
	ele, ok := msg.(map[string]interface{})
	if !ok {
		log.Println("return type err")
		return action, data, errors.New("return type err")
	}
	action, ok = ele["action"].(string)
	if !ok {
		log.Println("return type err")
		return action, data, errors.New("return type err")
	}
	data, ok = ele["data"].(map[string]interface{})
	if !ok {
		log.Println("return type err")
		return action, data, errors.New("return type err")
	}
	return action, data, nil
}
