package scratchcard

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Scratchcard struct {
	Id int
	// A list of numbers that will award points
	WinningNumbers []int
	// A list of numbers to compare against the WinningNumbers
	CardNumbers []int
}

// WinningCount returns the count of CardNumbers that are present in WinningNumbers
func (s Scratchcard) WinningCount() int {
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

	return winnersCount
}

// Points uses made up rules to figure out how many points you won on a scratch card.
// This solves puzzle 1, but doesn't follow the true rules of the scratchcard!
func (s Scratchcard) Points() int {
	winnersCount := s.WinningCount()
	return int(math.Pow(2, float64(winnersCount-1)))
}

// WinningCardIds returns a list of ids corresponding to other scratchcards.
// When a scratchcard scores a "point", it instead gives you a copy of another scratchcard
// whose id is equal to the id of the following cards.
//
// i.e. if Card 1 has 4 winners, it gives you the ids []int{2,3,4,5} back, indicating copies of scratchcards won
func (s Scratchcard) WinningCardIds() []int {
	winningCount := s.WinningCount()
	ids := []int{}

	for i := 1; i <= winningCount; i += 1 {
		ids = append(ids, s.Id+i)
	}

	return ids
}

// New creates a Scratchcard based off of input in the following format:
//
// Card Id: WinningNumbers | CardNumbers
//
// i.e. `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53`
func New(cardInput string) (Scratchcard, error) {
	s := Scratchcard{}

	// []{"Card 1"," 41 ... 17 | 83 ... 53"}
	parts := strings.Split(strings.Trim(cardInput, "\n"), ":")

	if len(parts) != 2 {
		return s, fmt.Errorf("invalid input: %s", cardInput)
	}

	id, err := parseCardId(parts[0])

	if err != nil {
		return s, err
	}

	s.Id = id
	s.WinningNumbers, s.CardNumbers = parseNumbers(parts[1])

	return s, nil
}

// parseCardId returns the number present in a string of the following format: Card %d
func parseCardId(cardIdSegment string) (int, error) {
	idParts := filterSlice(strings.Split(cardIdSegment, " "), func(segment string) bool {
		return len(segment) > 0
	})

	if len(idParts) != 2 || idParts[0] != "Card" {
		return 0, fmt.Errorf("cannot parse card id: %s", cardIdSegment)
	}

	id, err := strconv.Atoi(idParts[1])

	if err != nil {
		return 0, fmt.Errorf("cannot parse card id: %w", err)
	}

	return id, nil
}

// parseNumbers returns []int slices from a string of the following format: %d %d ... %d | %d ... %d
//
// If no numbers could be parsed, empty slices may be returned
func parseNumbers(numbersSegment string) ([]int, []int) {
	parts := strings.Split(strings.TrimSpace(numbersSegment), "|")

	if len(parts) != 2 {
		return []int{}, []int{}
	}

	// turn the string of numbers into a []string{} with extra spaces removed from each entry
	winningNumbersStr := strings.TrimSpace(parts[0])
	cardNumbersStr := strings.TrimSpace(parts[1])
	winningNumbers := mapSlice(strings.Split(winningNumbersStr, " "), strings.TrimSpace)
	cardNumbers := mapSlice(strings.Split(cardNumbersStr, " "), strings.TrimSpace)
	// lets remove non-digit elements
	winningNumbers = filterSlice(winningNumbers, onlyDigits)
	cardNumbers = filterSlice(cardNumbers, onlyDigits)

	// and finally convert them into numbers proper
	return mapSlice(winningNumbers, atoi), mapSlice(cardNumbers, atoi)
}

// filterSlice returns a new slice containing only elements that match the predicate
func filterSlice[T any](slice []T, predicate func(T) bool) []T {
	filtered := []T{}

	for _, v := range slice {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// mapSlice returns a new slice with the result of a mapping function applied to each element
func mapSlice[T any, R any](slice []T, mapper func(T) R) []R {
	mapped := []R{}

	for _, v := range slice {
		mapped = append(mapped, mapper(v))
	}

	return mapped
}

// isDigit determines if a character is a digit (0-9)
func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

// onlyDigits determines if a string consists of only digit characters (0-9)
func onlyDigits(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, char := range str {
		if !isDigit(char) {
			return false
		}
	}

	return true
}

// atoi is as strconv.Atoi, but ignores the error
func atoi(str string) int {
	n, _ := strconv.Atoi(str)

	return n
}
