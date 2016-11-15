package main

import "fmt"

// Command line interface, halts here until exit
func cli()  {

	var command string	// Last command line command

	for {
		fmt.Println("Action:")
		if _, err := fmt.Scanf("%s", &command); err != nil {
			fmt.Printf("ERROR(cli): %s\n", err)
		}else{
			if "exit" == command {	// Terminate command
				fmt.Println("Terminating...")
				return
			} else {			// Default case
				fmt.Printf("Sorry, %q is not a command. ", command)
				fmt.Printf("Type <%q> to terminate.", "exit")
				fmt.Println();
			}
		}
	}
}
