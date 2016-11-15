package main

import "encoding/json"

// Message to be sent to clients
type Message struct {
	Action string `json:"action"`
	User   *User  `json:"user"`
}

// Converting message to json prior to sending
func (sm *Message) toJson() []byte {
	message, _ := json.Marshal(sm)
	return message
}

// Used to fan in messages from both inputs to the output
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