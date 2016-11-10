package main

import "fmt"

func cli()  {

	var command string	//Last command line command

	for {
		fmt.Println("Action:")
		if _, err := fmt.Scanf("%s", &command); err != nil {
			fmt.Printf("ERROR(cli): %s\n", err)
		}
		if "exit" == command {
			fmt.Println("Terminating...")
			return
		}
		fmt.Printf("Sorry, %q is not a command.  (exit) to terminate.\n", command)
	}
}