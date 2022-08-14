package swisspost

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// getSecretAPIKey asks for the apiKey without echo
func getSecretAPIKey() string {
	fmt.Printf("Please enter your valid swisspost.opendatasoft.com API-Key: ")
	apiKeyInputValue, err := term.ReadPassword(0)
	if err != nil {
		fmt.Println("Could not read api key")
		os.Exit(1)
	}
	fmt.Println()

	return string(apiKeyInputValue)
}
