package scratchcard

import (
	"strings"
)

// A Pile is a bunch of Scratchcards
type Pile []Scratchcard

// Points returns the sum of each call to Scratchcard.Points()
func (p Pile) Points() int {
	pointsSum := 0

	for _, s := range p {
		pointsSum += s.Points()
	}

	return pointsSum
}

// Process calculates the new pile of scratchcards you will have received based off the the copy winning rules outlined in part 2 of the puzzle.
//
// # This version is way faster than OldProcess
//
// According the puzzle, there can never be a case where a card can win you a copy of a non-existant card.
func (p Pile) Process() int {
	totalCardCount := 0
	scWinCounts := map[int]int{}

	// Iterate over the pile in reverse
	// Because the last card can never be a winner, and cards can only
	// award bonus cards that appear after them - we can figure out
	// how many cards a given id will award in advance. We can then step back
	// over the pile, and continue to add cards based off of the number of winners it has
	for i := len(p) - 1; i >= 0; i -= 1 {
		sc := p[i]
		// This is the number of cards this scratch card will add to the deck as bonus cards
		scWinCount := sc.WinningCount()
		bonusCardsAwarded := scWinCount

		// For each winning card, we can then check how many each copy it awards
		// should add to the total number of cards
		for j := 1; j <= scWinCount; j += 1 {
			// If we've calculated the cards that come after this one before, we can add how many cards they add to our total
			// These should always be set, because we are going _backwards_ through the pile
			if n, found := scWinCounts[sc.Id+j]; found {
				bonusCardsAwarded += n
			}
		}

		scWinCounts[sc.Id] = bonusCardsAwarded
		totalCardCount += 1 + bonusCardsAwarded // always account for this one

	}

	return totalCardCount
}

// OldProcess calculates the new pile of scratchcards you will have received based off the the copy winning rules outlined in part 2 of the puzzle.
//
// This version took seconds to run on my PC (Ryzen 7800X3D + 32GB RAM)
//
// According the puzzle, there can never be a case where a card can win you a copy of a non-existant card.
func (p Pile) OldProcess() Pile {
	pMap := p.ToMap()
	newPile := Pile{}
	newPile = append(newPile, p...) // copy original cards in

	// Iterate over new pile so that we can keep processing until we have exhausted cards to add
	// Not using a range so that it always checks the updated length with each new iteration
	for i := 0; i < len(newPile); i += 1 {
		s := newPile[i]
		for _, id := range s.WinningCardIds() {
			// Even though this condition can't occur in the problem, it helps me sleep at night
			if s, found := pMap[id]; found {
				newPile = append(newPile, s)
			}
		}

	}

	return newPile
}

// ToMap creates a map whose keys are the ids of each scratchcard in the pile. The value of each key is the corresponding scratchcard. This assumes no two cards have the same id and differing numbers.
func (p Pile) ToMap() map[int]Scratchcard {
	pMap := map[int]Scratchcard{}

	for _, s := range p {
		pMap[s.Id] = s
	}

	return pMap

}

// NewPile functions as New but processes multiple lines of input at once.
//
// If an error is encountered when parsing a specific line, it is ignored.
// This means it is possible to have an empty slice returned, if every line is invalid.
func NewPile(cardsInput string) Pile {
	cardPile := Pile{}
	lines := strings.Split(strings.Trim(cardsInput, "\n"), "\n")

	for _, line := range lines {
		if s, err := New(line); err == nil { // if no error
			cardPile = append(cardPile, s)
		}
	}

	return cardPile
}
