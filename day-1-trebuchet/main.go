package main

import (
	"fmt"

	"github.com/bLittle1996/advent-of-code-2023/day-1-trebuchet/internal/trebuchet"
)

const input string = "1abc2" +
	"\npqr3stu8vwx" +
	"\na1b2c3d4e5f" +
	"\ntreb7uchet"

// ways to solve:
// 1. Iterate over each string, keeping track of the first and last numbers you run across
// 2. Iterate from the start until you hit a number, that's the first number. Repeat for the end by iterating from the end -> start, the first number is the last number.
// 3. ???

// For this problem, we're not gonna handle the error case of there being _no_ numbers - we'll just treat it as 0

func main() {

	fmt.Println("--- COMPUTING CALIBRATION VALUE FOR DOCUMENT ---")
	fmt.Println(fmt.Sprintf("%s", input))
	fmt.Println("")

	fmt.Println("--- CALCULATIONS COMPLETE ---")
	fmt.Println(fmt.Sprintf("Calibration Value: %d", trebuchet.GetCalibrationValueForDocument(input)))

}
