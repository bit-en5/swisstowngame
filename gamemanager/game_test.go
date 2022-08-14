package gamemanager

import (
	"testing"

	"github.com/bit-en5/swisstowngame/args"
	"github.com/bit-en5/swisstowngame/console"
	"github.com/bit-en5/swisstowngame/suggestions"
)

func TestGetSuggestions(t *testing.T) {
	args := args.New()

	g := New(args, mockSwissPost(), suggestions.New(), console.New())
	list, _ := g.getQuestions()

	// Check number of questions
	if len(list) != g.args.GetNbQuestions() {
		t.Errorf("expected %d questions / has %d", g.args.GetNbQuestions(), len(list))
	}

	// Check number of suggestions
	for _, q := range list {
		if len(q.Suggestions) != g.args.GetNbSuggestions() {
			t.Errorf("expected %d suggestions / has %d", g.args.GetNbSuggestions(), len(q.Suggestions))
		}
	}
}

/* Mock SwissPost service */

type mockswisspost struct{}

func mockSwissPost() mockswisspost {
	return mockswisspost{}
}

func (m mockswisspost) GetTown(bfsNr int) (string, string, error) {
	return "BE", "Zollikofen", nil
}
