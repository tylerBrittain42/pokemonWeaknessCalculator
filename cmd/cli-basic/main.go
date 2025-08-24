package main

import (
	"fmt"
	"os"

	pokewrapper "github.com/tylerBrittain42/pokemonWeaknessCalculator/pkg/pokeWrapper"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No name entered")
		return
	}
	name := os.Args[1]
	interactions, err := pokewrapper.GetPokemonTypeInteraction(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(interactions)
}
