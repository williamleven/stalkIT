package main

import "encoding/json"

type StalkITMessage struct {
	IsInTheHub	bool	`json:"is_in_the_hub"`
	User		*User	`json:"user"`
}

func (sm *StalkITMessage) toJson() ([]byte)  {
	message, _ := json.Marshal(sm)
	return message
}