package pokewrapper

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		caseName        string
		inputPokemon    string
		expectedPokemon string
		shouldError     bool
		errorMessage    string
	}{
		{
			caseName:        "base case",
			inputPokemon:    "turtwig",
			expectedPokemon: "turtwig",
			shouldError:     false,
			errorMessage:    "",
		}, {
			caseName:        "assorted caps",
			inputPokemon:    "BuLbASaur",
			expectedPokemon: "bulbasaur",
			shouldError:     false,
			errorMessage:    "",
		}, {
			caseName:        "multiple word name without dashes",
			inputPokemon:    "iron hands",
			expectedPokemon: "iron-hands",
			shouldError:     false,
			errorMessage:    "",
		}, {
			caseName:        "multiple world name using dashes",
			inputPokemon:    "iron-hands",
			expectedPokemon: "iron-hands",
			shouldError:     false,
			errorMessage:    "",
		}, {
			caseName:        "form name with spaces",
			inputPokemon:    "sandslash alola",
			expectedPokemon: "sandslash-alola",
			shouldError:     false,
			errorMessage:    "",
		}, {
			caseName:        "form name with dashes",
			inputPokemon:    "sandslash alola",
			expectedPokemon: "sandslash-alola",
			shouldError:     false,
			errorMessage:    "",
		}, {
			caseName:        "nothing given",
			inputPokemon:    "",
			expectedPokemon: "",
			shouldError:     true,
			errorMessage:    "empty string was given as name",
		}, {
			caseName:        "only whitespace given",
			inputPokemon:    "     ",
			expectedPokemon: "",
			shouldError:     true,
			errorMessage:    "empty string was given as name",
		},
	}

	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {
			output, err := cleanInput(tt.inputPokemon)
			if !tt.shouldError {
				if output != tt.expectedPokemon {
					t.Errorf("Got %s, expected %s", output, tt.expectedPokemon)
				}

			} else {
				if err == nil {
					t.Errorf("%s %s did not error out", tt.caseName, tt.inputPokemon)
				} else if err.Error() != tt.errorMessage {
					t.Errorf("Got %s, expected %s", err.Error(), tt.errorMessage)
				}

			}
		},
		)
	}

}

// getType of pokemon
func TestGetType(t *testing.T) {
	cases := []struct {
		caseName      string
		inputPokemon  string
		expectedTypes []string
		shouldError   bool
		errorMessage  string
	}{
		{
			caseName:      "single type(base case)",
			inputPokemon:  "turtwig",
			expectedTypes: []string{"grass"},
			shouldError:   false,
			errorMessage:  "",
		}, {
			caseName:      "dual type",
			inputPokemon:  "bulbasaur",
			expectedTypes: []string{"grass", "poison"},
			shouldError:   false,
			errorMessage:  "",
		}, {
			caseName:      "multiple word name without dashes(invalid)",
			inputPokemon:  "iron hands",
			expectedTypes: []string{},
			shouldError:   true,
			errorMessage:  "unable to find match for: iron hands",
		}, {
			caseName:      "multiple world name using dashes",
			inputPokemon:  "iron-hands",
			expectedTypes: []string{"fighting", "electric"},
			shouldError:   false,
			errorMessage:  "",
		}, {
			caseName:      "alolan form",
			inputPokemon:  "sandslash-alola",
			expectedTypes: []string{"ice", "steel"},
			shouldError:   false,
		},
		// cases

		// multiple regions, think the snow sandslash <--consider using generation for handling these casses
		//no result found
		// mega
		// type changes depending on generation(Azurill or other fairy)
	}

	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {
			output, err := getType(tt.inputPokemon)
			if !tt.shouldError {
				if len(output) != len(tt.expectedTypes) {
					t.Errorf("for %s, the improper number of types were listed", tt.inputPokemon)
				}
				if output[0] != tt.expectedTypes[0] {
					t.Errorf("for %s, %s != %s", tt.inputPokemon, output[0], tt.expectedTypes[0])
				}

				if len(output) == 2 && output[1] != tt.expectedTypes[1] {
					t.Errorf("for %s, %s-%s != (expected) %s-%s", tt.inputPokemon, output[0], output[1], tt.expectedTypes[0], tt.expectedTypes[1])
				}
			} else {
				if err == nil {
					t.Errorf("%s %s did not error out", tt.caseName, tt.inputPokemon)
				}

			}
		},
		)
	}

}

func TestGetTypeInteraction(t *testing.T) {
	cases := []struct {
		caseName     string
		inputType    string
		interactions PureTypeInteractions
	}{
		{
			caseName:  "Ghost type interactions",
			inputType: "ghost",
			interactions: PureTypeInteractions{
				DoubleDamageFrom: []string{"ghost", "dark"},
				HalfDamageFrom:   []string{"poison", "bug"},
				NoDamageFrom:     []string{"normal", "fighting"},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {
			output, err := getPureTypeInteraction(tt.inputType)
			if err != nil {
				t.Errorf("Encountered error: %v", err)
			}
			if !reflect.DeepEqual(output.DoubleDamageFrom, tt.interactions.DoubleDamageFrom) {
				t.Errorf("Double damage from: got %v expectedc %v", output.DoubleDamageFrom, tt.interactions.DoubleDamageFrom)
			} else if !reflect.DeepEqual(output.HalfDamageFrom, tt.interactions.HalfDamageFrom) {
				t.Errorf("Half damage from: got %v expectedc %v", output.HalfDamageFrom, tt.interactions.HalfDamageFrom)
			} else if !reflect.DeepEqual(output.NoDamageFrom, tt.interactions.NoDamageFrom) {
				t.Errorf("No damage: got %v expectedc %v", output.NoDamageFrom, tt.interactions.NoDamageFrom)
			}

		},
		)
	}

}

// func TestGetPokemonTypeInteractionDouble(t *testing.T) {
// func TestGetPokemonTypeInteractionHalf(t *testing.T) {
// func TestGetPokemonTypeInteractionQuarter(t *testing.T) {
// func TestGetPokemonTypeInteractionNone(t *testing.T) {
// single type pokemon(check all?)
// dual type pokemon(check all?)

// todo add mixed
func TestGetPokemonTypeInteractionFullPokemon(t *testing.T) {
	cases := []struct {
		caseName string
		pokemon  string
		types    TypeInteractions
	}{
		{
			caseName: "single-type",
			pokemon:  "furret",
			types: TypeInteractions{
				DoubleDamageFrom: []string{"fighting"},
				NoDamageFrom:     []string{"ghost"},
			},
		}, {
			caseName: "dual-type",
			pokemon:  "nuzleaf",
			types: TypeInteractions{
				QuadDamageFrom:   []string{"bug"},
				DoubleDamageFrom: []string{"fighting", "flying", "poison", "fire", "ice", "fairy"},
				HalfDamageFrom: []string{
					"ground", "ghost", "water", "grass", "electric", "dark",
				},
				QuarterDamageFrom: []string{},
				NoDamageFrom:      []string{},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {
			output, err := GetPokemonTypeInteraction(tt.pokemon)
			if err != nil {
				t.Errorf("Encountered error: %v", err)
			}
			if !sameStringSlice(output.QuadDamageFrom, tt.types.QuadDamageFrom) {
				t.Errorf("Quad damage from: got %v expected %v", output.QuadDamageFrom, tt.types.QuadDamageFrom)
			}
			if !sameStringSlice(output.DoubleDamageFrom, tt.types.DoubleDamageFrom) {
				t.Errorf("Double damage from: got %v expected %v", output.DoubleDamageFrom, tt.types.DoubleDamageFrom)
			}
			if !sameStringSlice(output.HalfDamageFrom, tt.types.HalfDamageFrom) {
				t.Errorf("Half damage from: got %v expected %v", output.HalfDamageFrom, tt.types.HalfDamageFrom)
			}
			if !sameStringSlice(output.QuarterDamageFrom, tt.types.QuarterDamageFrom) {
				t.Errorf("Quarter damage from: got %v expected %v", output.QuarterDamageFrom, tt.types.QuarterDamageFrom)
			}
			if !sameStringSlice(output.NoDamageFrom, tt.types.NoDamageFrom) {
				t.Errorf("No damage from: got %v expected %v", output.NoDamageFrom, tt.types.NoDamageFrom)
			}
		},
		)
	}

}
func TestGetPokemonTypeInteractionQuad(t *testing.T) {
	cases := []struct {
		caseName string
		pokemon  string
		quad     []string
	}{
		{
			caseName: "case A",
			pokemon:  "nuzleaf",
			quad: []string{
				"bug",
			},
		},
		{
			caseName: "case B",
			pokemon:  "scizor",
			quad: []string{
				"fire",
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {
			output, err := GetPokemonTypeInteraction(tt.pokemon)
			if err != nil {
				t.Errorf("Encountered error: %v", err)
			}
			if !sameStringSlice(output.QuadDamageFrom, tt.quad) {
				t.Errorf("Double damage from: got %v expected %v", output.DoubleDamageFrom, tt.quad)
			}
		},
		)
	}

}

/*

	{
		{caseName: "single-type"},
		{caseName: "dual-type(has 4x)"},
		{caseName: "dual-type(has 1/4x)"},
		{caseName: "dual-type(cancel each other out)"},
		{caseName: "dual-type(has none)"},

		{caseName: "1"},
		{caseName: "1/2"},
		{caseName: "1/4"},
		{caseName: "2"},
		{caseName: "4"},
		{caseName: "0"},
	}
	for _, tt := range cases {

	}
}
*/

// returns true if the contents are identical
// sorted does not matter
func sameStringSlice(x, y []string) bool {
	// obv if the lengths are different they can't have identical contents
	if len(x) != len(y) {
		return false
	}

	// think of diff like this:
	// having an instance in x increments,
	// having a matching y instance decrements
	// if y attempts to decrement a non-existent entry,
	// we know they do not match bc that is a value that y has that x does not
	// if there are any values in the map at the end,
	// those are values that x has, but y is missing
	// NOTE: This holds true for any x and y, provided x goes first
	diff := make(map[string]int, len(x))
	for _, val := range x {
		diff[val]++
	}
	for _, val := range y {
		// this means that x never had it
		if _, ok := diff[val]; !ok {
			return false
		}
		diff[val]--
		if diff[val] == 0 {
			delete(diff, val)
		}
	}
	return len(diff) == 0
}
