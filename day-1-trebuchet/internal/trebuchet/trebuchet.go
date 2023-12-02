package trebuchet

import (
	"strconv"
	"strings"
)

var textDigits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var textDigitReplacer = strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")

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

// O(n^2) I think :(
// Start from the back, the first digit we see - we return it and stop
// This is trickier to do that front to back because we need to look ahead to the end of the string again
// to see if we have any words...
func getLastDigit(str string, includeWrittenDigits bool) int {
	for i := len(str) - 1; i >= 0; i -= 1 {
		char := str[i]
		digit, err := strconv.ParseInt(string(char), 10, 0)

		if err != nil && includeWrittenDigits {
			// This errored because the character wasn't a number, if that's the case lets
			// create a new str from now until the end and replace the words in it with digits and try again
			// This means we end up replacing word digits from their first occurence right to left, rather than left to right
			// i.e. twone -> tw1 instead of 2ne
			// We could potentially avoid this nested looping (worst case O(n^2)) by
			// Adding a seperate replacer for reversed strings (eno -> 1, owt -> two)
			// and simply running that against the reversed string and calling getFirst on it lmao

			// Now that we are doing a ParseInt on a large string, we need to make sure
			// that we remove non-digit characters at this point
			subStr := strings.Map(func(char rune) rune {
				if char >= '0' && char <= '9' {
					return char
				}

				// A negative return value drops the character!
				return -1
			}, textDigitReplacer.Replace(str[i:]))

			digit, err = strconv.ParseInt(subStr, 10, 0)

		}

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

func filter[T any](arr []T, predicate func(T) bool) []T {
	filteredArr := []T{}

	for _, value := range arr {
		if predicate(value) {
			filteredArr = append(filteredArr, value)
		}
	}

	return filteredArr
}
