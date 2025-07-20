package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	fmt.Print("start:")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		sanitizedInput := cleanInput(input)
		command := sanitizedInput[0]
		if val, ok := getCommands()[command]; ok {
			val.callback()
		} else {
			fmt.Println("Unknown command")
		}

	}

}
func cleanInput(text string) []string {
	sanitizedText := strings.Trim(strings.ToLower(text), " ")
	return strings.Fields(sanitizedText)
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
