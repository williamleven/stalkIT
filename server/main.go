package main

import "net"

func main() {
	input := make(chan *Users) 		// Delivering user Lists
	outputArrivals := make(chan *User) 	// Delivering users arriving
	outputDepartures := make(chan *User) 	// Delivering users leaving

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":7825")

	// accept connection on port
	conn, _ := ln.Accept()

	// Handles frontend notifications
	go notifier(outputArrivals, true, conn)
	go notifier(outputDepartures, false, conn)

	// Coordinates data-flow
	go coordinate(input, outputArrivals, outputDepartures)

	// Grabs and times data from hubbit
	go collectorTimer(10, "https://hubbit.chalmers.it/sessions.json", input)

	// Halts at command line interface
	cli()
}