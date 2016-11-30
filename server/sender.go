package main

import (
	"github.com/Gurgy/broadcastPool"
	"net"
	"github.com/spacemonkeygo/openssl"
)

// Sends notification in the format input + append
func sender(messages chan *Message) {
	var CERTIFICATE string = "server.crt"
	var KEY string = "server.key"

	// Creating certificate
	ctx, err := openssl.NewCtxFromFiles(CERTIFICATE, KEY)
	if err != nil {
		panic(err)
	}

	// Creating listener
	l, err := openssl.Listen("tcp", ":7777", ctx)
	if err != nil {
		panic(err)
	}

	// Creates a new pool of connections
	pool := broadcastPool.New(validateConnection)

	// Accept connections to this pool
	pool.Open(l)

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
