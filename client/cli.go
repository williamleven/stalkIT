package main

import "fmt"

// Command line interface, halts here until exit
func cli()  {

	var command string	// Last command line command

	for {
		fmt.Println(*translate("cli_prompt"))
		if _, err := fmt.Scanf("%s", &command); err != nil {
			fmt.Printf(*translate("cli_error") + "\n", err)
		}else{
			if "exit" == command {	// Terminate command
				fmt.Println(*translate("quit_message"))
				return
			} else if "locale" == command {
				fmt.Println(language.Code)
			} else {			// Default case
				fmt.Printf(*translate("cli_invalid_command") + " ", command)
				fmt.Printf(*translate("cli_termination_guide"), "exit")
				fmt.Println();
			}
		}
	}
}
