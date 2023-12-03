package gears_test

import (
	"testing"

	"github.com/bLittle1996/advent-of-code-2023/gear-ratios/data"
	"github.com/bLittle1996/advent-of-code-2023/gear-ratios/internal/gears"
)

const testDataAnswer int = 4361

func TestThatItSolvesTheExampleInput(t *testing.T) {
	engine := gears.NewEngine(data.TestData)

	if gears.PartNumberSum(engine) != testDataAnswer {
		t.Errorf("Expected %d, received %d", testDataAnswer, gears.PartNumberSum(engine))
	}
}

func TestThatItSolvesMyArbitraryExamples(t *testing.T) {
	type testCase struct {
		engineSchematic string
		expectedSum     int
	}

	cases := []testCase{
		{"...", 0},
		{"...\n...", 0},
		{"...\n...\n...", 0},
		{"1..", 0},
		{".1.", 0},
		{"..1", 0},
		{"1.*", 0},
		{".1*", 1},
		{".*1", 1},
		{"1..\n*..", 1},
		{"1.&\n111\n111", 111},
		{"1..\n420\n6*9", 420 + 6 + 9},
	}

	for i, testCase := range cases {
		engine := gears.NewEngine(testCase.engineSchematic)

		if gears.PartNumberSum(engine) != testCase.expectedSum {
			t.Errorf("Test #%d: Expected %d, received %d", i, testCase.expectedSum, gears.PartNumberSum(engine))
		}
	}
}
