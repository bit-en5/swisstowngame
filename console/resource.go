package console

import (
	"bufio"
	"os"
)

type Console struct {
	reader *bufio.Reader
}

func New() Console {
	return Console{
		reader: bufio.NewReader(os.Stdin),
	}
}
