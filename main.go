package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello")
}

func cleanInput(text string) []string {
	sanitizedText := strings.Trim(strings.ToLower(text), " ")
	return strings.Fields(sanitizedText)
}
