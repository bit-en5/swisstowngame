package suggestions

import (
	"math/rand"

	"github.com/bit-en5/swisstowngame/cantondata"
)

// GetRandom returns a list of randomized cantons + the canton expected as the response
func (s Suggestions) GetRandomList(cantonResponse string, nb int) (res []string) {
	// Create a list of cantons to return
	list := make(map[string]bool)

	// Add the canton response in the list
	list[cantonResponse] = true

	// Complete the list with other cantons until the expected number has been reached
	var canton string
	for len(list) < nb {
		canton = s.getRandom()
		if canton != "" {
			list[canton] = true
		}
	}

	// Convert the map into []string
	for canton := range list {
		res = append(res, canton)
	}

	// Randomize the list
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})

	return
}

// getRandom returns a canton according to a randomize choice
func (s Suggestions) getRandom() string {
	k := rand.Intn(len(cantondata.GetList()))

	for key := range cantondata.GetList() {
		if k == 0 {
			return key
		}
	}

	return ""
}
