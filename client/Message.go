package main

// Message structure to be received from server
type Message struct {
	Action string `json:"action"`
	User   *User  `json:"user"`
}
