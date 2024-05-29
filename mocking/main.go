package main

import (
	"fmt"
	"io"
	"os"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func main() {
	Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprintf(out, finalWord)
}
