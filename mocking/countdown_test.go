package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {

		buffer := &bytes.Buffer{}

		Countdown(buffer, &SpyCoundtdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCoundtdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

type SpyCoundtdownOperations struct {
	Calls []string
}

func (s *SpyCoundtdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCoundtdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const (
	write = "wirte"
	sleep = "sleep"
)
