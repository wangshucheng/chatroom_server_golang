package network

import (
	"encoding/json"
)

type SendMessage struct {
	Sender string `json:sender`
	Context string `json:context`
}

type ReadMessage struct {
	Context string `json:context`
}

func MakeMessage(sender string,context string) ([]byte,error)  {
	msg := SendMessage{
		Sender: sender,
		Context: context,
	}

	jsonStr,err := json.Marshal(msg)
	if err != nil {
		return nil,err
	}

	return []byte(jsonStr),err
}

func ResolveMessage(buf []byte) (string,error) {
	var msg ReadMessage
	err := json.Unmarshal(buf, &msg)
	if err != nil {
		return "",err
	}
	return msg.Context,err
}