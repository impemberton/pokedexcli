package main

import (
	"fmt"
	"bufio"
	"os"
)
func init() {
registry = map[string]cliCommand {
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
}
}

type cliCommand struct {
		name string
		description string
		callback func() error
	}
var registry map[string]cliCommand


func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")	
	fmt.Println("Usage:\n")	
	for _, command := range registry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}



func main() {

	

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		user_input := cleanInput(scanner.Text())
		found := false
		for _, command := range registry {
			if user_input[0] == command.name {
				found = true
				command.callback()
			}
		}
		if !found {
			fmt.Println("Unknown command")
		}
	}
}

