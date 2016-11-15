package main

func main() {
	input := make(chan *Users) // Delivering user Lists
	messages := make(chan *Message)

	// Grabs and times data from hubbit
	go collectorTimer(10, "https://hubbit.chalmers.it/sessions.json", input)

	// Coordinates data-flow
	go analyze(input, messages)

	// Handles frontend notifications
	go sender(messages)

	// Halts at command line interface
	cli()
}
