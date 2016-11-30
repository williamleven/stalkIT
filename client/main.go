package main

var language *Language

func main() {
	// Detect system language
	var err error
	language, err = NewLanguage(DetectLocale())
	if err != nil {
		panic("Failed to set locale.")
	}

	// Handles frontend notifications
	go notifier()

	// Halts at command line interface
	cli()
}
