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
