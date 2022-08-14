package main

import (
	"fmt"

	"github.com/bit-en5/swisstowngame/args"
	"github.com/bit-en5/swisstowngame/console"
	"github.com/bit-en5/swisstowngame/gamemanager"
	"github.com/bit-en5/swisstowngame/suggestions"
	"github.com/bit-en5/swisstowngame/swisspost"
)

func main() {
	fmt.Println("**********************************************************************************************")
	fmt.Println("* Welcome to the Swiss Town Guess Game!                                                      *")
	fmt.Println("*                                                                                            *")
	fmt.Println("* Thanks to post.ch / swisspost.opendatasoft.com / BFS for providing free data for this quiz *")
	fmt.Println("**********************************************************************************************")

	// Read flags
	appArgs := args.New()

	// Manage termination
	handleTermination(appArgs.GetTimeoutApplication())

	// Inject dependencies and start game
	gamemanager.New(
		appArgs,
		swisspost.New(appArgs.GetAPIKey()), // this object is responsible to get data by swisspost
		suggestions.New(),                  // this object is responsible to get canton data
		console.New(),                      // this object is responsible to interract with the user
	).Start()
}
