package print

import "fmt"

type Color string

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

func Red(msg string) {
	fmt.Println(red + msg + reset)
}

func Green(msg string) {
	fmt.Println(green + msg + reset)
}

func Yellow(msg string) {
	fmt.Println(yellow + msg + reset)
}

func Blue(msg string) {
	fmt.Println(blue + msg + reset)
}

func Purple(msg string) {
	fmt.Println(purple + msg + reset)
}

func Cyan(msg string) {
	fmt.Println(cyan + msg + reset)
}

func Grey(msg string) {
	fmt.Println(gray + msg + reset)
}

func White(msg string) {
	fmt.Println(white + msg + reset)
}
