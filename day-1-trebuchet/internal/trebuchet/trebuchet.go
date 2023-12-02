package trebuchet

import (
	"strconv"
	"strings"
)

// Start from the front, the first gigit we see - return it and cease
func getFirstDigit(str string) int {
	for _, char := range str {
		digit, err := strconv.ParseInt(string(char), 10, 0)

		if err != nil {
			continue
		}

		return int(digit)
	}

	return 0
}

// Start from the back, the first digit we see - we return it and stop
func getLastDigit(str string) int {
	for i := len(str) - 1; i > 0; i -= 1 {
		char := str[i]
		digit, err := strconv.ParseInt(string(char), 10, 0)

		if err != nil {
			continue
		}

		return int(digit)
	}

	return 0
}

func getCalibrationValue(firstDigit int, lastDigit int) int {
	// Multiply the first digit by 10 to shift it into the 10s column
	// i.e instead of 1 + 2 (3), we get 10 + 2 (12).
	return firstDigit*10 + lastDigit
}

// Returns the sums of the first and last digit of each line in a given string (seperated by \n)
func GetCalibrationValueForDocument(calibrationDocument string) int {
	lines := strings.Split(calibrationDocument, "\n")

	sumOfCalibrationValues := 0

	for _, line := range lines {
		value := getCalibrationValue(getFirstDigit(line), getLastDigit(line))

		sumOfCalibrationValues += value
	}

	return sumOfCalibrationValues
}
