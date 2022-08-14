package args

type AppArgs struct {
	nbQuestions   int
	nbSuggestions int
	timeoutApp    int
	timeoutQuest  int
	apiKey        string
}

func New() (appArgs AppArgs) {
	appArgs = AppArgs{}
	appArgs.readArgs()

	return
}
