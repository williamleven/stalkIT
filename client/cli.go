package main

import "fmt"

// Command line interface, halts here until exit
func cli()  {

	var command string	// Last command line command

	for {
		fmt.Println(language.getPhrase("cli_prompt"))
		if _, err := fmt.Scanf("%s", &command); err != nil {
			fmt.Printf(language.getPhrase("cli_error") + "\n", err)
		}else{
			if "exit" == command {	// Terminate command
				fmt.Println(language.getPhrase("quit_message"))
				return
			} else if "locale" == command {
				fmt.Println(language.Code)
			} else {			// Default case
				fmt.Printf(language.getPhrase("cli_invalid_command") + " ", command)
				fmt.Printf(language.getPhrase("cli_termination_guide"), "exit")
				fmt.Println();
			}
		}
	}
}
