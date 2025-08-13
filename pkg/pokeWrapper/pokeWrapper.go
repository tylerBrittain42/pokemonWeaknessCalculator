package pokewrapper

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func cleanInput(name string) (string, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "", errors.New("empty string was given as name")
	}
	final := strings.ToLower(strings.ReplaceAll(trimmed, " ", "-"))

	return final, nil

}

}

func Foo() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))
}
