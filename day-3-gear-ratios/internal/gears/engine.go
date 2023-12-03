package gears

import (
	"strconv"
	"strings"
)

const invalidSymbol rune = '.'

type Engine struct {
	rawSchematic string
	Schematic    []string
	PartNumbers  []int
}

func NewEngine(rawSchematic string) Engine {
	lines := strings.Split(strings.Trim(rawSchematic, "\n"), "\n")

	engine := Engine{
		rawSchematic: rawSchematic,
		Schematic:    lines,
	}

	engine.computePartNumbers()

	return engine
}

func (engine *Engine) computePartNumbers() {
	if engine == nil {
		return
	}

	for i, line := range engine.Schematic {
		numStr := ""
		startsAtIndex := 0
		for j, char := range line {
			// If we haven't added a digit yet, we mark the starting point
			// of any potential digits as the current index
			if numStr == "" {
				startsAtIndex = j
			}

			if isDigit(char) {
				numStr += string(char)
			}

			// Can process a number if the character isn't a digit or we are at the end
			if numStr != "" && (!isDigit(char) || j == len(line)-1) {
				// Added a number and got to a non-number? Time to check if it's a part number!!!
				num, err := strconv.Atoi(numStr)

				if err == nil && engine.isValidPartNumber(startsAtIndex, i, len(numStr)) {
					engine.PartNumbers = append(engine.PartNumbers, num)
				}

				// Reset numStr so we can build new numbers
				numStr = ""
			}
		}
	}
}

func (engine *Engine) isValidPartNumber(x int, y int, width int) bool {
	rowWidth := len(engine.Schematic[y])
	rowCount := len(engine.Schematic)

	canLookLeft := x > 0
	canLookRight := x+(width-1) < rowWidth-1
	canLookUp := y > 0
	canLookDown := y < rowCount-1

	adjacentLeft, adjacentRight, adjacentUp, adjacentDown := false, false, false, false

	leftBound := x
	rightBound := x + width

	if canLookLeft {
		leftBound -= 1
	}

	if canLookRight {
		rightBound += 1
	}

	// can look to the left on the rows
	if canLookLeft {
		s := engine.Schematic[y][leftBound:x]

		adjacentLeft = containsSymbol(s)
	}

	// can look to the right on the rights
	if canLookRight {
		s := engine.Schematic[y][x+width : rightBound]

		adjacentRight = containsSymbol(s)
	}

	// Can look at previous row
	if canLookUp {
		s := engine.Schematic[y-1][leftBound:rightBound]

		adjacentUp = containsSymbol(s)
	}

	// can look at next row
	if canLookDown {
		s := engine.Schematic[y+1][leftBound:rightBound]

		adjacentDown = containsSymbol(s)
	}

	return adjacentLeft || adjacentRight || adjacentUp || adjacentDown
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func containsSymbol(lineSegment string) bool {
	for _, char := range lineSegment {
		if !isDigit(char) && char != invalidSymbol {
			return true
		}
	}

	return false
}

func PartNumberSum(engine Engine) int {
	sum := 0

	for _, partNum := range engine.PartNumbers {
		sum += partNum
	}

	return sum
}
