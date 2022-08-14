package gamemanager

import (
	"fmt"
	"math/rand"

	"github.com/bit-en5/swisstowngame/cantondata"
	"github.com/bit-en5/swisstowngame/print"
)

const (
	maxBFSNr = 6820
	maxRetry = 3
)

// List of cases to handle by answer check
const (
	valid = iota
	invalid
	timeout
	tooManyRetry
)

func (g GameManager) Start() {
	// Get questions
	_, nbCorrect := g.getQuestions()

	fmt.Println("*******************************************************")
	fmt.Println("* This game is over. Thank you for playing!           *")
	fmt.Println("*                                                     *")
	fmt.Printf("* You've got %d out of %d questions right!              *\n", nbCorrect, g.args.GetNbQuestions())
	fmt.Println("*******************************************************")
}

func (g GameManager) getQuestions() (list []Question, nbCorrect int) {
	var q Question

	// While loop
	for len(list) < g.args.GetNbQuestions() {
		// Generate a question
		q = g.getQuestion()

		if q.CantonCode != "" && q.City != "" {
			// Display the question and interract with the user
			q.Result = g.display(q)
			if q.Result {
				nbCorrect++
			}
			list = append(list, q)
		}
	}

	return
}

func (g GameManager) getQuestion() Question {
	bfsNr := g.getRandomBFSNr()

	// Get the city & canton data
	cantonCode, city, err := g.post.GetTown(bfsNr)
	if err != nil {
		return Question{}
	}

	// Prepare the question
	q := Question{
		BFSNr:       bfsNr,
		City:        city,
		CantonCode:  cantonCode,
		CantonName:  cantondata.GetName(cantonCode),
		Suggestions: g.suggestions.GetRandomList(cantonCode, g.args.GetNbSuggestions()),
	}

	q.ExpectedAnswer = q.getAnswerNr(cantonCode)

	return q
}

func (g GameManager) getRandomBFSNr() int {
	return rand.Intn(maxBFSNr)
}

func (g GameManager) display(q Question) (isCorrect bool) {
	var (
		totalAnswers int
		answer       int
		err          error
	)

	// Print question
	fmt.Printf("In which canton is %s?\n", q.City)

	// Print possible answers
	for i, cantonCode := range q.Suggestions {
		fmt.Printf("Type %d for %s\n", (i + 1), cantonCode)
	}

	// Initial message
	msg := "Your answer: "

	for {
		// Interract with the user
		answer, err = g.console.AskForInt(msg, g.args.GetTimeoutQuestion())
		switch g.selectCase(answer, totalAnswers, err) {
		case invalid:
			totalAnswers++
			// Change message for retry
			msg = fmt.Sprintf("Answer must be between 1 and %d: ", g.args.GetNbSuggestions())
			continue

		case timeout:
			print.Red("Time is over")
			return

		case tooManyRetry:
			print.Red("Too many attempts")
			return
		}

		// Check if answer is correct
		isCorrect = answer == q.ExpectedAnswer

		// Display result
		if isCorrect {
			print.Green("CORRECT!")
		} else {
			print.Red("WRONG")
		}

		return
	}
}

func (g GameManager) selectCase(answer, totalAnswers int, err error) int {
	if totalAnswers >= maxRetry {
		return tooManyRetry
	}

	if err != nil {
		if err.Error() == "timeout" {
			return timeout
		}

		return invalid
	}

	if answer < 1 || answer > g.args.GetNbSuggestions() {
		return invalid
	}

	return valid
}
