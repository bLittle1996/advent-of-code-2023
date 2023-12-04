package scratchcard

import "math"

type Scratchcard struct {
	Id int
	// A list of numbers that will award points
	WinningNumbers []int
	// A list of numbers to compare against the WinningNumbers
	CardNumbers []int
}

func (s Scratchcard) Points() int {
	winningNumbers := map[int]bool{}
	winnersCount := 0 // how many of the scratch cards numbers were winning numbers

	// Assign to a map for O(1) lookup rather than having to iterate over the numbers for each number in the card
	for _, n := range s.WinningNumbers {
		winningNumbers[n] = true
	}

	for _, n := range s.CardNumbers {
		if isWinning := winningNumbers[n]; isWinning {
			winnersCount += 1
		}
	}

	if winnersCount == 0 {
		return 0
	}

	return int(math.Pow(2, float64(winnersCount-1)))
}

// New creates a Scratchcard based off of input in the following format:
//
// Card Id: WinningNumbers | CardNumbers
//
// i.e. `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53`
func New(cardInput string) Scratchcard {
	s := Scratchcard{}

	return s
}

// NewPile functions as New but processes multiple lines of input at once.
func NewPile(cardsInput string) []Scratchcard {
	return []Scratchcard{}
}
