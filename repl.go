package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("WizWord101 > ")
		reader.Scan()

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

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"start": {
			name:			"start",
			description:	"Start a round of guessing!",
			callback:		commandStart,
		},
		"login": {
			name: 			"login",
			description:	"Bot Login for API references",
			callback:		commandLogin,
		},
		"ref": {
			name: 			"ref",
			description:	"reference wiki for spell info",
			callback:		commandRef,
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