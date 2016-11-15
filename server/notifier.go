package main

import (
	"net"
)

// Sends notification in the format input + append
func notifier(input chan *User, arriving bool, conn net.Conn)  {

	// Pushing notifications whenever receiving input
	for {
		user := <-input

		conn.Write([]byte(int8(arriving) + user.Nick + "\n"))
	}
}