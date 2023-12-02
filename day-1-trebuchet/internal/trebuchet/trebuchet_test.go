package trebuchet_test

import (
	"fmt"
	"testing"

	"github.com/bLittle1996/advent-of-code-2023/day-1-trebuchet/internal/trebuchet"
)

const exampleInput string = "1abc2" +
	"\npqr3stu8vwx" +
	"\na1b2c3d4e5f" +
	"\ntreb7uchet\n"
const exampleAnswer int = 142

const exampleHardInput string = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
const exampleHardAnswer = 281

func TestThatItSolvesTheExampleProblem(t *testing.T) {
	value := trebuchet.GetCalibrationValueForDocument(exampleInput, false)

	if value != exampleAnswer {
		t.Error(fmt.Sprintf("Expected %d, received %d", exampleAnswer, value))
	}
}

func TestThatItCanDoTheHardVersionOfTheProblem(t *testing.T) {
	value := trebuchet.GetCalibrationValueForDocument(exampleHardInput, true)

	if value != exampleHardAnswer {
		t.Error(fmt.Sprintf("Expected %d, received %d", exampleHardAnswer, value))
	}
}

func TestThatItCanSolveIndividualLines(t *testing.T) {
	type testCase struct {
		line                  string
		answer                int
		replaceTextWithDigits bool
	}

	cases := []testCase{
		{"", 0, false},
		{"1", 11, false},
		{"2", 22, false},
		{"3", 33, false},
		{"0", 0, false},
		{"12", 12, false},
		{"69", 69, false},
		{"jklweaf3weopajkfjiwaefo89weuipafhwaefh", 39, false},
		{"111111111111111111111111111111", 11, false},
		{"1234456767897890", 10, false},
		{"wioafheuinvweaijbnweavuihweaf8y923r789yweauibof89q23f7g832wweg89ua", 89, false},
		{"", 0, true},
		{"one", 11, true},
		{"two", 22, true},
		{"three", 33, true},
		{"four", 44, true},
		{"five", 55, true},
		{"six", 66, true},
		{"seven", 77, true},
		{"eight", 88, true},
		{"nine", 99, true},
		{"sixnine", 69, true},
		{"sixtynine", 69, true},
		{"sixty", 66, true}, // has six in it...
		{"seventy", 77, true},
		{"fourtwenty", 44, true},
		{"1fourtwenty3", 13, true},
		{"1fourtwenty", 14, true},
		{"7fourtwenty", 74, true},
		{"threeandseven", 37, true},
		// Tricky gotchas
		{"oneighthree", 13, true},
		{"eightwothree", 83, true},
		{"twoneight", 28, true},
		// Lines from input, manually solved
		{"1rdtwofjvdllht5eightsixfourbl", 14, true},
		{"oneeightmbbklndlztwo9nine18", 18, true},
		{"4cmhfccrttfive", 45, true},
		{"fivetptfpone89ponefourjxmdrjkrleightwoh", 52, true},
		{"3199", 39, true},
		{"cbtxjqqdqc56fourjhgtrjsxnbxnineeightwov", 52, true},
		{"kkbjcptltjsdjrlhzzg4drkffivezkxl", 45, true},
		{"7mptczpscnq4vdfbveightfourjkhnhlkrkgch", 74, true},
		{"pnineonetwo2", 92, true},
		{"5qrhonetflqdnsztwonine9vnctxjnine", 59, true},
		{"kgoneightqqxlrhtpx58threethree7vvqq", 17, true},
		{"sixjt76fiveninedjzpceight3", 63, true},
		{"twotwo5", 25, true},
		{"hzbfour63nttfktqjzjhponeightcz", 48, true},
		{"ksqkttvninesevensix4pbnjsfznch5dlxfq", 95, true},
		{"5212j912", 52, true},
		{"xeightwopbgt7two", 82, true},
		{"fiveqdrgeightvrvtmtvvmnl2kbctwofive", 55, true},
		// Now I see the problem with my implementation, here's a test case to fix it
		{"oneight", 18, true},
	}

	for i, testCase := range cases {
		value := trebuchet.GetCalibrationValueForDocument(testCase.line, testCase.replaceTextWithDigits)

		if value != testCase.answer {
			t.Errorf("Test #%d: Given %s, expected %d. Received %d.", i, testCase.line, testCase.answer, value)
		}
	}
}
