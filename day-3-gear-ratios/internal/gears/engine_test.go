package gears_test

import (
	"testing"

	"github.com/bLittle1996/advent-of-code-2023/gear-ratios/data"
	"github.com/bLittle1996/advent-of-code-2023/gear-ratios/internal/gears"
)

const testDataAnswer int = 4361
const testGearRatioAnswer int = 467835

func TestThatItSolvesTheExampleInput(t *testing.T) {
	engine := gears.NewEngine(data.TestData)

	if engine.GetPartNumberSum() != testDataAnswer {
		t.Errorf("part num sum: Expected %d, received %d", testDataAnswer, engine.GetPartNumberSum())
	}

	if engine.GetGearRatioSum() != testGearRatioAnswer {
		t.Errorf("gear ratio: Expected %d, received %d", testGearRatioAnswer, engine.GetGearRatioSum())
	}

}

func TestThatItSolvesMyArbitraryExamples(t *testing.T) {
	type testCase struct {
		engineSchematic string
		expectedSum     int
		expectedRatio   int
	}

	cases := []testCase{
		{"...", 0, 0},
		{"...\n...", 0, 0},
		{"...\n...\n...", 0, 0},
		{"1..", 0, 0},
		{".1.", 0, 0},
		{"..1", 0, 0},
		{"1.*", 0, 0},
		{".1*", 1, 0},
		{".*1", 1, 0},
		{"1..\n*..", 1, 0},
		{"1.&\n111\n111", 111, 0},
		{"1..\n420\n6*9", 420 + 6 + 9, 0},
		// Gear ratio ones
		{"...\n...\n...", 0, 0},
		{"...\n.*.\n...", 0, 0},
		{"...\n***\n...", 0, 0},
		{"***\n***\n***", 0, 0},
		{"1..\n.*.\n..1", 2, 1},
		{"3..\n.*.\n..2", 5, 6},
		{"3..\n.*.\n222", 222 + 3, 666},
		// invalid, 3 adjacent parts
		{"3..\n.*.\n2.2", 2 + 2 + 3, 0},
	}

	for i, testCase := range cases {
		engine := gears.NewEngine(testCase.engineSchematic)

		if engine.GetPartNumberSum() != testCase.expectedSum {
			t.Errorf("Test #%d Part Numbers: Expected %d, received %d", i, testCase.expectedSum, engine.GetPartNumberSum())
		}

		if engine.GetGearRatioSum() != testCase.expectedRatio {
			t.Errorf("Test #%d Gear Ratios: Expected %d, received %d", i, testCase.expectedSum, engine.GetGearRatioSum())
		}
	}
}
