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
	}{
		{
			caseName:        "base case",
			inputPokemon:    "turtwig",
			expectedPokemon: "turtwig",
			shouldError:     false,
		}, {
			caseName:        "assorted caps",
			inputPokemon:    "BuLbASaur",
			expectedPokemon: "bulbasaur",
			shouldError:     false,
		}, {
			caseName:        "multiple word name without dashes",
			inputPokemon:    "iron hands",
			expectedPokemon: "iron-hands",
			shouldError:     false,
		}, {
			caseName:        "multiple world name using dashes",
			inputPokemon:    "iron-hands",
			expectedPokemon: "iron-hands",
			shouldError:     false,
		}, {
			caseName:        "form name with spaces",
			inputPokemon:    "sandslash alola",
			expectedPokemon: "sandslash-alola",
			shouldError:     false,
		}, {
			caseName:        "form name with dashes",
			inputPokemon:    "sandslash alola",
			expectedPokemon: "sandslash-alola",
			shouldError:     false,
		}, {
			caseName:        "nothing given",
			inputPokemon:    "",
			expectedPokemon: "",
			shouldError:     true,
		}, {
			caseName:        "only whitespace given",
			inputPokemon:    "     ",
			expectedPokemon: "",
			shouldError:     true,
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
		expectedTypes [2]string
		shouldError   bool
	}{
		{
			caseName:      "single type(base case)",
			inputPokemon:  "turtwig",
			expectedTypes: [2]string{"grass", "none"},
			shouldError:   false,
		},
		// cases
		// dual type
		// multiple regions, think the snow sandslash
		//no result found
		//multiple words in name
		// mega
		// type changes depending on generation(Azurill or other fairy)
	}

	for _, tt := range cases {
		t.Run(tt.caseName, func(t *testing.T) {
			output, err := getType(tt.inputPokemon)
			if !tt.shouldError {
				if output[0] != tt.expectedTypes[0] && output[1] != tt.expectedTypes[1] {
					t.Errorf("for %s, %s-%s != %s-%s", tt.inputPokemon, output[0], output[1], tt.expectedTypes[0], tt.expectedTypes[1])
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
