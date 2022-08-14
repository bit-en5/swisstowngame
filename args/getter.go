package args

import "time"

func (a AppArgs) GetNbQuestions() int {
	return a.nbQuestions
}

func (a AppArgs) GetNbSuggestions() int {
	return a.nbSuggestions
}

func (a AppArgs) GetTimeoutApplication() time.Duration {
	return time.Duration(a.timeoutApp) * time.Second
}

func (a AppArgs) GetTimeoutQuestion() time.Duration {
	return time.Duration(a.timeoutQuest) * time.Second
}

func (a AppArgs) GetAPIKey() string {
	return a.apiKey
}
