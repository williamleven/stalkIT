package main

import (
	"github.com/Gurgy/broadcastPool"
	"net"
	"crypto/tls"
	"time"
	"crypto/rand"
)

// Sends notification in the format input + append
func sender(messages chan *Message) {
	var CERTIFICATE string = "device.cer"
	var KEY string = "device.key"

	// Creating certificate
	cert, err := tls.LoadX509KeyPair(CERTIFICATE, KEY)
	if err != nil {
		panic(err)
	}

	config := tls.Config{
		Certificates: []tls.Certificate{cert}}

	now := time.Now()

	config.Time = func() time.Time { return now }

	config.Rand = rand.Reader

	// Creating listener
	l, err := tls.Listen("tcp", ":7825", &config)

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
