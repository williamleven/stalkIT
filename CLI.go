package main

import "fmt"

func cli()  {
	var command string
	for {
		fmt.Println("Action:")
		if _, err := fmt.Scanf("%s", &command); err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		if "exit" == command {
			fmt.Print("Terminating...")
			return
		}
		fmt.Printf("Sorry, %q is not a command.  (exit) to terminate.\n", command)
	}
}
