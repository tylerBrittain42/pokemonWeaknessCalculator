package pokewrapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Sanitizes the name of a pokemon
func cleanInput(name string) (string, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "", errors.New("empty string was given as name")
	}
	final := strings.ToLower(strings.ReplaceAll(trimmed, " ", "-"))

	return final, nil

}

// Returns the type of the pokemon
//
// Assumes that the name has been sanitized
func getType(name string) ([]string, error) {
	var specPokemon Pokemon
	types := []string{}

	url := "https://" + domain + "pokemon/" + name
	resp, err := http.Get(url)
	if err != nil {
		return []string{}, nil
	}

	if resp.StatusCode != 200 {
		return []string{}, fmt.Errorf("unable to find match for: %s", name)
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode((&specPokemon))
	if err != nil {
		return []string{}, err
	}

	types = append(types, specPokemon.Types[0].Type.Name)
	if len(specPokemon.Types) > 1 {
		types = append(types, specPokemon.Types[1].Type.Name)
	}

	return types, nil

}

func Foo() {
	getType("bulbasaur")
}
