package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tylerBrittain42/pokemonWeaknessCalculator/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func startRepl(cfg *config) {

	fmt.Print("start:")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		sanitizedInput := cleanInput(input)
		command := sanitizedInput[0]
		if val, ok := getCommands()[command]; ok {
			err := val.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
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
	callback    func(*config) error
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
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandMapb,
		},
	}
}
