package main

func main() {
	// Detect system language
	detectLocale()

	// Handles frontend notifications
	go notifier()

	// Halts at command line interface
	cli()
}
