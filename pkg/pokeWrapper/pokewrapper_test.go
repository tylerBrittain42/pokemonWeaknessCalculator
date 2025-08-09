package pokewrapper

import (
	"testing"
)

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
