package main

import (
	"strings"

	pokewrapper "github.com/tylerBrittain42/pokemonWeaknessCalculator/pkg/pokeWrapper"
)

func main() {
	pokewrapper.Foo()
}

func cleanInput(text string) []string {
	sanitizedText := strings.Trim(strings.ToLower(text), " ")
	return strings.Fields(sanitizedText)
}
