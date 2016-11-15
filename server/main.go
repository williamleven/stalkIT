package main

func main() {
	input := make(chan *Users) 		// Delivering user Lists
	outputArrivals := make(chan *User) 	// Delivering users arriving
	outputDepartures := make(chan *User) 	// Delivering users leaving

	// Handles frontend notifications
	go notifier(outputArrivals, " has arrived at the Hubb")
	go notifier(outputDepartures, " has left the Hubb")

	// Coordinates data-flow
	go coordinate(input, outputArrivals, outputDepartures)

	// Grabs and times data from hubbit
	go collectorTimer(10, "https://hubbit.chalmers.it/sessions.json", input)

	// Halts at command line interface
	cli()
}