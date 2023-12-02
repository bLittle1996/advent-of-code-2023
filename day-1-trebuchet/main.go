package main

import (
	"fmt"
	"os"

	"github.com/bLittle1996/advent-of-code-2023/day-1-trebuchet/internal/trebuchet"
)

// ways to solve:
// 1. Iterate over each string, keeping track of the first and last numbers you run across
// 2. Iterate from the start until you hit a number, that's the first number. Repeat for the end by iterating from the end -> start, the first number is the last number.
// 3. ???

// There is a hard mode version of this question where the digits can be spelt out as words
// In this case, we convert the words to a digit and treat it as a number.
// 1. Replace all instances of a number word with the digit equivalent
// 2. Have a map of words -> digits and do above?
// note that only one, two, three, ..., nine count as valid digits. 10, teens, 20, 30, hundred, etc do not.

// For this problem, we're not gonna handle the error case of there being _no_ numbers - we'll just treat it as 0

func main() {
	contents, err := os.ReadFile("./data/data.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	input := string(contents)

	fmt.Println("--- COMPUTING CALIBRATION VALUE FOR DOCUMENT ---")
	fmt.Println(fmt.Sprintf("%s", input))
	fmt.Println("")

	fmt.Println("--- CALCULATIONS COMPLETE ---")
	fmt.Println(fmt.Sprintf("Calibration Value: %d", trebuchet.GetCalibrationValueForDocument(input, false)))
	fmt.Println(fmt.Sprintf("Advanced Calibration Value: %d", trebuchet.GetCalibrationValueForDocument(input, true)))

}
