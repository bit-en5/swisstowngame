package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bit-en5/swisstowngame/print"
)

func handleTermination(timeout time.Duration) {
	fmt.Println(timeout)
	handleCancel()
	handleTimeout(timeout)
}

func handleCancel() {
	go func() {
		quit := make(chan os.Signal, 10)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("")
		print.Cyan("Sorry to see you leave the game")
		os.Exit(0)
	}()
}

func handleTimeout(timeout time.Duration) {
	if timeout > 0 {
		// Quit after timeout
		go func() {
			<-time.After(timeout)
			fmt.Println()
			print.Red("Sorry, your game time is up")
			os.Exit(0)
		}()
	}
}
