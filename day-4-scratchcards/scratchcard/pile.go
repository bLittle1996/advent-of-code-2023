package scratchcard

import "strings"

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
