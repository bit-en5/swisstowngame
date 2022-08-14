package args

import "flag"

func (a *AppArgs) readArgs() (appArgs AppArgs) {
	flag.IntVar(&a.nbQuestions, "questions", 5, "set number of questions")
	flag.IntVar(&a.nbSuggestions, "suggestions", 3, "set number of suggestions")
	flag.IntVar(&a.timeoutApp, "timeoutApp", 0, "max game duration in second")
	flag.IntVar(&a.timeoutQuest, "timeoutQuest", 0, "max game duration in second")
	flag.StringVar(&a.apiKey, "apiKey", "", "set swisspost apiKey (optional)")

	flag.Parse()

	return
}
