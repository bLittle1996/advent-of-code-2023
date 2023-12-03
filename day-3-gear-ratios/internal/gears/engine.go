package gears

import (
	"strconv"
	"strings"
)

const invalidSymbol rune = '.'
const gearSymbol rune = '*'

type Engine struct {
	rawSchematic string
	Schematic    []string
	parts        []part
	gears        []gear
}

type part struct {
	X     int
	Y     int
	Width int
	Value int
}

type gear struct {
	X int
	Y int
}

func NewEngine(rawSchematic string) Engine {
	lines := strings.Split(strings.Trim(rawSchematic, "\n"), "\n")

	engine := Engine{
		rawSchematic: rawSchematic,
		Schematic:    lines,
	}

	engine.computeBitsAndBobs()

	return engine
}

// Populates the parts and gears slices as well as the PartNumbers slice.
func (engine *Engine) computeBitsAndBobs() {
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
			} else if char == gearSymbol {
				engine.gears = append(engine.gears, gear{
					X: j,
					Y: i,
				})
			}

			// Can process a number if the character isn't a digit or we are at the end
			if numStr != "" && (!isDigit(char) || j == len(line)-1) {
				num, _ := strconv.Atoi(numStr)

				part := part{
					X:     startsAtIndex,
					Y:     i,
					Width: len(numStr),
					Value: num,
				}

				if engine.isValidPart(part) {
					engine.parts = append(engine.parts, part)
				}

				// Reset numStr so we can build new numbers
				numStr = ""
			}
		}
	}
}

func (engine *Engine) isValidPart(part part) bool {
	rowWidth := len(engine.Schematic[part.Y])
	rowCount := len(engine.Schematic)

	canLookLeft := part.X > 0
	canLookRight := part.X+(part.Width-1) < rowWidth-1
	canLookUp := part.Y > 0
	canLookDown := part.Y < rowCount-1

	adjacentLeft, adjacentRight, adjacentUp, adjacentDown := false, false, false, false

	leftBound := part.X
	rightBound := part.X + part.Width

	if canLookLeft {
		leftBound -= 1
	}

	if canLookRight {
		rightBound += 1
	}

	// can look to the left on the rows
	if canLookLeft {
		s := engine.Schematic[part.Y][leftBound:part.X]

		adjacentLeft = containsSymbol(s)
	}

	// can look to the right on the rights
	if canLookRight {
		s := engine.Schematic[part.Y][part.X+part.Width : rightBound]

		adjacentRight = containsSymbol(s)
	}

	// Can look at previous row
	if canLookUp {
		s := engine.Schematic[part.Y-1][leftBound:rightBound]

		adjacentUp = containsSymbol(s)
	}

	// can look at next row
	if canLookDown {
		s := engine.Schematic[part.Y+1][leftBound:rightBound]

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

	for _, part := range engine.parts {
		sum += part.Value
	}

	return sum
}
