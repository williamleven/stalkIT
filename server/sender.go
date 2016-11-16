package main

import (
	"github.com/Gurgy/broadcastPool"
	"net"
)

// Sends notification in the format input + append
func sender(messages chan *Message) {

	// Creates a new pool of connections
	pool := broadcastPool.New(validateConnection)

	// Accept connections to this pool
	pool.Open(port)

	// Send message to all clients whenever receiving input
	var message *Message
	for {
		message = <-messages

		// Broadcast to all clients in the connection pool
		pool.Broadcast(message.toJson())
		pool.Broadcast([]byte("\n"))
	}
}

func validateConnection(c net.Conn) bool  {
	return true
}
