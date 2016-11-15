package main

import "encoding/json"

type Message struct {
	Action string `json:"action"`
	User   *User  `json:"user"`
}

func (sm *Message) toJson() []byte {
	message, _ := json.Marshal(sm)
	return message
}
