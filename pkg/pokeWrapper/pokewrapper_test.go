package pokewrapper

import (
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

// get strength and weakness of tyhpe
// get type info of specific pokemon
