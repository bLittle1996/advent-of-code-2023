package trebuchet

import (
	"strconv"
	"strings"
)

var textDigitReplacer = strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
var reversedTextDigitReplacer = strings.NewReplacer("eno", "1", "owt", "2", "eerht", "3", "ruof", "4", "evif", "5", "xis", "6", "neves", "7", "thgie", "8", "enin", "9")

// O(n)
// Start from the front, the first gigit we see - return it and cease
func getFirstDigit(str string, includeWrittenDigits bool) int {
	// For left -> right scanning, we can simply replace every text instance of a number
	// with the appropriate number
	if includeWrittenDigits {
		str = textDigitReplacer.Replace(str)
	}

	for _, char := range str {
		digit, err := strconv.ParseInt(string(char), 10, 0)

		if err != nil {
			continue
		}

		return int(digit)
	}

	return 0
}

// O(n)
// Start from the back, the first digit we see - we return it and stop
func getLastDigit(str string, includeWrittenDigits bool) int {
	reversedInput := reverse(str)

	if includeWrittenDigits {
		reversedInput = reversedTextDigitReplacer.Replace(reversedInput)
	}

	for _, char := range reversedInput {
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
func GetCalibrationValueForDocument(calibrationDocument string, includeWrittenDigits bool) int {
	lines := strings.Split(calibrationDocument, "\n")

	sumOfCalibrationValues := 0

	for _, line := range lines {
		value := getCalibrationValue(getFirstDigit(line, includeWrittenDigits), getLastDigit(line, includeWrittenDigits))

		sumOfCalibrationValues += value
	}

	return sumOfCalibrationValues
}

func reverse(str string) string {
	if len(str) <= 1 {
		return str
	}

	strSlice := strings.Split(str, "")

	for startIndex, endIndex := 0, len(strSlice)-1; startIndex < endIndex; startIndex, endIndex = startIndex+1, endIndex-1 {
		endIndexStr := strSlice[endIndex]
		strSlice[endIndex] = strSlice[startIndex]
		strSlice[startIndex] = endIndexStr
	}

	return strings.Join(strSlice, "")
}
