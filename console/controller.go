package console

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (c Console) Ask(msg string, timeout time.Duration) (answer string, err error) {
	if timeout > 0 {
		answer, err = c.askWithTimeout(msg, timeout)
	} else {
		answer, err = c.ask(msg)
	}
	if err != nil {
		return
	}

	// Remove the delimiter from input string
	answer = strings.TrimSuffix(answer, "\n")

	return
}

func (c Console) ask(msg string) (answer string, err error) {
	fmt.Print(msg)
	answer, err = c.reader.ReadString('\n')
	if err != nil {
		return
	}

	return
}

func (c Console) askWithTimeout(msg string, timeout time.Duration) (answer string, err error) {
	ch := make(chan bool)

	go func() {
		fmt.Print(msg)
		answer, err = c.reader.ReadString('\n')
		if err != nil {
			return
		}

		ch <- true
	}()

	select {
	case <-ch:

	case <-time.After(timeout):
		fmt.Println("time expired")
		err = errors.New("timeout")
	}

	return
}

func (c Console) AskForInt(msg string, timeout time.Duration) (int, error) {
	answer, err := c.Ask(msg, timeout)
	if err != nil {
		return 0, err
	}

	// Check if answer is a valid number
	return strconv.Atoi(answer)
}
