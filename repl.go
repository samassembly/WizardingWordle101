package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	//Fun Banner to indicate new cli has been entered
	printAsciiBanner()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("WizWord101 > ")
		reader.Scan()

		//input validation
		words := CleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		
		// for longer commands
		// input := ""
		// if len(words) > 1 {
		// 	input = words[1]
		// }

		//check if command exists
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

//formats text input for ease of comparisons
func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

//command structure within repl loop
type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

//prints commands available within repl
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"start": {
			name:			"start",
			description:	"Start a round of guessing!",
			callback:		commandStart,
		},
		"help": {
			name:	 		"help",
			description:	"Displays a help message",
			callback: 		commandHelp,
		},
		"exit": {
			name:			"exit",
			description:	"Exit WizWord101",
			callback:		commandExit,
		},
	}
}