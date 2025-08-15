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
		return []string{}, err
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

func getPureTypeInteraction(elementType string) (PureTypeInteractions, error) {
	// One reusable type for name objects from PokeAPI

	// Your main type for decoding the GET request
	type respType struct {
		DamageRelations struct {
			DoubleDamageFrom []DamageTypeInfo `json:"double_damage_from"`
			HalfDamageFrom   []DamageTypeInfo `json:"half_damage_from"`
			NoDamageFrom     []DamageTypeInfo `json:"no_damage_from"`
		} `json:"damage_relations"`
	}
	var typeChart respType

	url := "https://" + domain + "type/" + elementType
	resp, err := http.Get(url)
	if err != nil {
		return PureTypeInteractions{}, err
	}

	if resp.StatusCode != 200 {
		return PureTypeInteractions{}, fmt.Errorf("unable to find match for: %s", elementType)
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&typeChart)
	if err != nil {
		return PureTypeInteractions{}, err
	}
	interactions := PureTypeInteractions{
		DoubleDamageFrom: stripKey(typeChart.DamageRelations.DoubleDamageFrom),
		HalfDamageFrom:   stripKey(typeChart.DamageRelations.HalfDamageFrom),
		NoDamageFrom:     stripKey(typeChart.DamageRelations.NoDamageFrom),
	}
	return interactions, nil

}

func getPokemonTypeInteraction(name string) (TypeInteractions, error) {
	sanitizedName, err := cleanInput(name)
	if err != nil {
		return TypeInteractions{}, err
	}

	types, err := getType(sanitizedName)
	if err != nil {
		return TypeInteractions{}, err
	}

	if len(types) == 1 {
		interactions, err := getPureTypeInteraction(types[0])
		if err != nil {
			return TypeInteractions{}, err
		}
		return TypeInteractions{HalfDamageFrom: interactions.HalfDamageFrom, DoubleDamageFrom: interactions.DoubleDamageFrom, NoDamageFrom: interactions.NoDamageFrom}, nil
	}

	return TypeInteractions{}, nil

}

func stripKey(items []DamageTypeInfo) []string {
	stripped := []string{}
	for _, v := range items {
		stripped = append(stripped, v.Name)
	}
	return stripped
}

func Foo() {
	getType("bulbasaur")
}
