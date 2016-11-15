package main

import (
	"net"
)

// Sends notification in the format input + append
func notifier(arrivals chan *User, departures chan *User)  {

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":7825")

	// accept connection on port
	//conn, _ := ln.Accept()

	connections := make(map[net.Conn]bool)

	go func(clients *map[net.Conn]bool) {
		for {
			conn, _ := ln.Accept()
			(*clients)[conn] = true
		}
	}(&connections)


	c := make(chan string)
	go func() {
		for {
			user := <-arrivals
			c <- "A"+user.Nick
		}
	}()
	go func() {
		for {
			user := <-departures
			c <- "D"+user.Nick
		}
	}()

	// Pushing notifications whenever receiving input
	for {
		message := <-c
		for conn, value := range connections {
			if value {
				conn.Write([]byte(message + "\n"))
			}
		}
	}
}

