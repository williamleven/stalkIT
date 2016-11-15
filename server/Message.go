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

func fanInMessages(input1 chan *Message, input2 chan *Message, output chan *Message) {
	go func() {
		for {
			output <- <-input1
		}
	}()
	go func() {
		for {
			output <- <-input2
		}
	}()
}