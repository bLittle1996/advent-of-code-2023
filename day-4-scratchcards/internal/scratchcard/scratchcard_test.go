package scratchcard_test

import (
	"testing"

	"github.com/bLittle1996/advent-of-code-2023/scratchcards/internal/scratchcard"
)

func TestThatItSolvesTheExampleScratchcard(t *testing.T) {
	expectedPoints := 8
	scratchcard := scratchcard.Scratchcard{
		Id:             1,
		WinningNumbers: []int{41, 48, 83, 86, 17},
		CardNumbers:    []int{83, 83, 6, 31, 17, 9, 48, 53},
	}

	if scratchcard.Points() != expectedPoints {
		t.Errorf("Expected %d, received %d", expectedPoints, scratchcard.Points())
	}
}
