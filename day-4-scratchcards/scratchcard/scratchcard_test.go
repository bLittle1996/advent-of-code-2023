package scratchcard_test

import (
	"reflect"
	"testing"

	"github.com/bLittle1996/advent-of-code-2023/scratchcards/data"
	"github.com/bLittle1996/advent-of-code-2023/scratchcards/scratchcard"
)

func TestThatItSolvesTheExampleScratchcard(t *testing.T) {
	expectedPoints := 8
	scratchcard := scratchcard.Scratchcard{
		Id:             1,
		WinningNumbers: []int{41, 48, 83, 86, 17},
		CardNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
	}

	if scratchcard.Points() != expectedPoints {
		t.Errorf("Expected %d, received %d", expectedPoints, scratchcard.Points())
	}

	if !reflect.DeepEqual(scratchcard.WinningCardIds(), []int{2, 3, 4, 5}) {
		t.Errorf("Expected winning card ids 2, 3, 4,and 5, got %v", scratchcard.WinningCardIds())
	}
}

func Test_New_ReturnsScratchcardForValidInput(t *testing.T) {
	type testCase struct {
		input                  string
		expectedId             int
		expectedWinningNumbers []int
		expectedCardNumbers    []int
		expectedPoints         int
	}

	cases := []testCase{
		{"Card      1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 1, []int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}, 8},
		{"Card   2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2, []int{13, 32, 20, 16, 61}, []int{61, 30, 68, 82, 17, 32, 24, 19}, 2},
		{"Card     3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 3, []int{1, 21, 53, 59, 44}, []int{69, 82, 63, 72, 16, 21, 14, 1}, 2},
		{"Card  4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 4, []int{41, 92, 73, 84, 69}, []int{59, 84, 76, 51, 58, 5, 54, 83}, 1},
		{"Card     5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 5, []int{87, 83, 26, 28, 32}, []int{88, 30, 70, 12, 93, 22, 82, 36}, 0},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 6, []int{31, 18, 13, 56, 72}, []int{74, 77, 10, 23, 35, 67, 36, 11}, 0},
	}

	for i, tc := range cases {
		s, err := scratchcard.New(tc.input)

		if err != nil {
			t.Errorf("Test Case %d New: %s", i, err.Error())
		}

		if s.Id != tc.expectedId {
			t.Errorf("Test Case %d: expected id %d, got %d", i, tc.expectedId, s.Id)
		}

		if !reflect.DeepEqual(s.CardNumbers, tc.expectedCardNumbers) {
			t.Errorf("Test Case %d: expected CardNumbers %v, got %v", i, tc.expectedCardNumbers, s.CardNumbers)
		}

		if !reflect.DeepEqual(s.WinningNumbers, tc.expectedWinningNumbers) {
			t.Errorf("Test Case %d: expected CardNumbers %v, got %v", i, tc.expectedWinningNumbers, s.WinningNumbers)
		}

		if s.Points() != tc.expectedPoints {
			t.Errorf("Test Case %d: expected points %d, got %d", i, tc.expectedPoints, s.Points())
		}
	}
}

func Test_NewPile_IsAbleToSolveTheExampleInput(t *testing.T) {
	cards := scratchcard.NewPile(data.TestData)
	pointSum := 0
	expectedPoints := 13

	for _, s := range cards {
		pointSum += s.Points()
	}

	if pointSum != expectedPoints {
		t.Errorf("expected %d, got %d", expectedPoints, pointSum)
	}
}
