package main

func main() {
	// Handles frontend notifications
	go notifier()

	// Halts at command line interface
	cli()
}
