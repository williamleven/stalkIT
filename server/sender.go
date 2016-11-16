package main

import (
	"fmt"
	"net"
)

// Sends notification in the format input + append
func sender(messages chan *Message) {

	// Handles connections
	connections := make(map[net.Conn]bool)
	go connectionAcceptor(&connections)

	// Send message to all clients whenever receiving input
	var message *Message
	for {
		message = <-messages
		for conn, value := range connections {
			if value {
				conn.Write(message.toJson())
				conn.Write([]byte("\n"))
			}
		}
	}
}

// Handles connections to all clients
func connectionAcceptor(clients *map[net.Conn]bool) {
	// Listen on all interfaces.
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("FATAL(connectionAcceptor): " + err.Error())
		panic(err)
	}

	// Accept all incoming connections and add them to the map
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERROR(connectionAcceptor): " + err.Error())
		} else {
			(*clients)[conn] = true
		}
	}
}
