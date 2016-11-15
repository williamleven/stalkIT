package main

import (
	"net"
)

// Sends notification in the format input + append
func notifier(arrivals chan *User, departures chan *User)  {

	connections := make(map[net.Conn]bool)

	go connectionAcceptor(&connections)


	messages := make(chan string)
	buildMessages(arrivals, departures, messages)


	// Send message to all clients whenever receiving input
	for {
		message := <-messages
		for conn, value := range connections {
			if value {
				conn.Write([]byte(message + "\n"))
			}
		}
	}
}

func connectionAcceptor(clients *map[net.Conn]bool){
	// Listen on all interfaces.
	ln, _ := net.Listen("tcp", ":7825")
	for {
		conn, _ := ln.Accept()
		(*clients)[conn] = true
	}
}

func buildMessages(arrivals chan *User, departures chan *User, output chan string){
	go func() {
		for {
			user := <-arrivals
			message := StalkITMessage{
				true,
				user,
			}
			output <- string(message.toJson())
		}
	}()
	go func() {
		for {
			user := <-departures
			message := StalkITMessage{
				false,
				user,
			}
			output <- string(message.toJson())
		}
	}()
}
