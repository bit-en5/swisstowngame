package gamemanager

type Question struct {
	BFSNr          int
	City           string
	CantonCode     string
	CantonName     string
	Suggestions    []string
	ExpectedAnswer int
	Result         bool
}

func (q Question) getAnswerNr(cantonCode string) int {
	for i, code := range q.Suggestions {
		if code == cantonCode {
			return i + 1
		}
	}

	return 0
}
