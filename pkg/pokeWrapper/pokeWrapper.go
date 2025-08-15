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

	firstTypeInteractions, err := getPureTypeInteraction(types[0])
	if err != nil {
		return TypeInteractions{}, err
	}

	if len(types) == 1 {
		return TypeInteractions{HalfDamageFrom: firstTypeInteractions.HalfDamageFrom, DoubleDamageFrom: firstTypeInteractions.DoubleDamageFrom, NoDamageFrom: firstTypeInteractions.NoDamageFrom}, nil
	}

	secondTypeInteractions, err := getPureTypeInteraction(types[1])
	if err != nil {
		return TypeInteractions{}, err
	}

	// todo: consider doing similar to the slice checker but use
	// nonexistent as normal
	// 1 as double
	// 2 as quad
	// -1 as half
	// -2 as quarter
	// check for immune at end
	combinedTypeInteractionsMap := make(map[string]float32)
	updateTypeMap(combinedTypeInteractionsMap, firstTypeInteractions)
	updateTypeMap(combinedTypeInteractionsMap, secondTypeInteractions)

	var finalTypeInteractions TypeInteractions
	for key, val := range combinedTypeInteractionsMap {
		switch val {
		case 0.25:
			finalTypeInteractions.QuarterDamageFrom = append(finalTypeInteractions.QuarterDamageFrom, key)
		case 0.5:
			finalTypeInteractions.HalfDamageFrom = append(finalTypeInteractions.HalfDamageFrom, key)
		case 1:
			continue
		case 2:
			finalTypeInteractions.DoubleDamageFrom = append(finalTypeInteractions.DoubleDamageFrom, key)
		case 4:
			finalTypeInteractions.QuadDamageFrom = append(finalTypeInteractions.QuadDamageFrom, key)
		}
	}

	return finalTypeInteractions, nil

}

// m = typeMap
// i = pureTypeInteractions
func updateTypeMap(m map[string]float32, i PureTypeInteractions) {
	// in cases where the type does not have a value, we set it,
	// otherwise we treat it as a multiplier
	for _, v := range i.DoubleDamageFrom {
		if _, ok := m[v]; !ok {
			m[v] = 2
		} else {
			m[v] *= 2
		}

	}
	for _, v := range i.HalfDamageFrom {
		if _, ok := m[v]; !ok {
			m[v] = 0.5
		} else {
			m[v] *= 0.5
		}

	}
	for _, v := range i.NoDamageFrom {
		if _, ok := m[v]; !ok {
			m[v] = 0
		} else {
			m[v] *= 0
		}

	}

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
