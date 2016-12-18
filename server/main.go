package main

func main() {

	// Grabs data from HubbIT and sends on input chanel
	input := make(chan *Users) // Delivering user Lists
	go smurfGetter("https://hubbit.chalmers.it/sessions.json", input)

	// Analyzes data from input chanel and send arriving/ departures on respective chanel
	newUsers := make(chan *User)   // Delivering new users
	missingUsers := make(chan *User) // Delivering missing users
	go analyze(input, newUsers, missingUsers)

	arrivals := make(chan *User)   // Delivering users arriving
	departures := make(chan *User) // Delivering users leaving
	go departureHandler(missingUsers, departures, newUsers, arrivals)

	// Packs users from the arrivals chanel and send them on messages1
	messages1 := make(chan *Message)
	go packer(arrivals, "Arrived", messages1)

	// Packs users from the departures chanel and send them on messages2
	messages2 := make(chan *Message)
	go packer(departures, "Departed", messages2)

	// Forward all messages from both messages1 and messages2 to messages
	messages := make(chan *Message)
	fanInMessages(messages1, messages2, messages)

	// Send messages to clients
	go sender(messages)

	// Halts at command line interface
	cli()
}
