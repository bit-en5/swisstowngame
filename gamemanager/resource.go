package gamemanager

import (
	"math/rand"
	"time"
)

// default values
const (
	nbQuestionsByDefault   = 5
	nbSuggestionsByDefault = 3
)

// GameManager handles the swiss town game
type GameManager struct {
	args        arger
	post        swisspostgetter
	suggestions suggestioner
	console     consoler
}

type arger interface {
	GetNbQuestions() int
	GetNbSuggestions() int
	GetTimeoutQuestion() time.Duration
}

type swisspostgetter interface {
	GetTown(bfsNr int) (string, string, error)
}

type suggestioner interface {
	GetRandomList(except string, nb int) []string
}

type consoler interface {
	AskForInt(msg string, timeout time.Duration) (int, error)
}

// New constructs a new GameManager
func New(
	args arger,
	post swisspostgetter,
	suggestions suggestioner,
	console consoler,
) GameManager {

	rand.Seed(time.Now().UnixNano())

	return GameManager{
		args:        args,
		post:        post,
		suggestions: suggestions,
		console:     console,
	}
}
