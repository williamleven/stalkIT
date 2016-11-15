package main

func main() {
	input := make(chan *Users) // Delivering user Lists

	// Grabs and times data from hubbit
	go smurfGetter("https://hubbit.chalmers.it/sessions.json", input)

	arrivals := make(chan *User)   // Delivering users arriving
	departures := make(chan *User) // Delivering users leaving

	// Coordinates data-flow
	go analyze(input, arrivals, departures)

	messages1 := make(chan *Message)
	go packer(arrivals, "Arrived", messages1)

	messages2 := make(chan *Message)
	go packer(departures, "Departed", messages2)

	messages := make(chan *Message)

	fanInMessages(messages1, messages2, messages)

	// Handles frontend notifications
	go sender(messages)

	// Halts at command line interface
	cli()
}
