package suggestions_test

import (
	"strings"
	"testing"

	"github.com/bit-en5/swisstowngame/suggestions"
)

func TestMultipleGetRandomList(t *testing.T) {
	for i := 0; i < 1000; i++ {
		TestGetRandomList(t)
	}
}

func TestGetRandomList(t *testing.T) {
	want := "BE"
	nb := 5

	suggestions := suggestions.New()
	has := suggestions.GetRandomList(want, nb)

	// Check the number of element in the list
	if len(has) != nb {
		t.Errorf("expect %d elements / has %d", nb, len(has))
	}

	// Check if the given canton is in the list
	flat := strings.Join(has, ",") + ","
	if !strings.Contains(flat, want+",") {
		t.Errorf("missing %s in the list", want)
	}

	// Each canton must be different
	var nbOccurence int

	for _, canton := range has {
		nbOccurence = strings.Count(flat, canton+",")
		if nbOccurence > 1 {
			t.Errorf("canton %s in the list", want)
		}
	}
}
