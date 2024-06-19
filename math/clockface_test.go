package clockfac_test

import (
	clockfac "LearnGoWithTests/math"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockfac.Point{X: 150, Y: 150 - 90}
	got := clockfac.SecondHand(tm)

	if got != want {
		t.Errorf("Got: %v, Wanted: %v", got, want)
	}
}
