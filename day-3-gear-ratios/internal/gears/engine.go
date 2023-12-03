package gears

import (
	"math"
	"strconv"
	"strings"
)

const invalidSymbol rune = '.'
const gearSymbol rune = '*'

type Engine struct {
	Schematic []string
	Parts     []part
	Gears     []gear
}

type part struct {
	X     int
	Y     int
	Width int
	Value int
}

type gear struct {
	X     int
	Y     int
	Ratio int
}

func NewEngine(rawSchematic string) Engine {
	lines := strings.Split(strings.Trim(rawSchematic, "\n"), "\n")

	engine := Engine{
		Schematic: lines,
	}

	// Populate the parts and gears slices
	parts, gears := getPartsAndGears(lines)
	engine.Parts = parts
	// At this point, we have only valid parts, but our gears might not be valid gears
	engine.Gears = getValidGears(parts, gears)

	return engine
}

func (engine Engine) GetPartNumberSum() int {
	sum := 0

	for _, part := range engine.Parts {
		sum += part.Value
	}

	return sum
}

func (engine Engine) GetGearRatioSum() int {
	sum := 0

	for _, gear := range engine.Gears {
		sum += gear.Ratio
	}

	return sum
}

// Populates the parts and gears slices
func getPartsAndGears(engineSchematics []string) ([]part, []gear) {
	parts, gears := []part{}, []gear{}

	for i, line := range engineSchematics {
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
				gears = append(gears, gear{
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

				if isValidPart(engineSchematics, part) {
					parts = append(parts, part)
				}

				// Reset numStr so we can build new numbers
				numStr = ""
			}
		}
	}

	return parts, gears
}

// Takes the list of gears and compares it against parts, returning a new list of gears containing only ones that are adjacent to two parts.
// Gear ratios are set in this function as well
func getValidGears(parts []part, potentialGears []gear) []gear {
	var validGears []gear

	for _, gear := range potentialGears {
		// We know the coordinate of a gear
		// We also know the positition of every part in the engine
		// We just need to compare indicies to see if it is adjacent or not
		// If the difference in coordinates is within +/- 1, it is adjacent
		// 1..   [0,0] -> [1,1] <- [2, 2]
		// .*.
		// ..1

		adjacentParts := []part{}
		for _, part := range parts {
			// Since a parts starting digit isn't guaranteed to be right next to the gear,
			// We'll need to check each coordinate that the number occupies based
			// off of it's width. offset.e. the number 100 takes up 3 possible spots that can be adjacent
			// to the gear
			for offset := 0; offset < part.Width; offset += 1 {
				if areCoordinatesAdjacent(gear.X, gear.Y, part.X+offset, part.Y) {
					adjacentParts = append(adjacentParts, part)
					break
				}
			}
		}

		// A valid gear can only have 2 adjacent parts!
		if len(adjacentParts) == 2 {
			gear.Ratio = adjacentParts[0].Value * adjacentParts[1].Value

			validGears = append(validGears, gear)
		}
	}

	return validGears
}

func isValidPart(engineSchematics []string, part part) bool {
	rowWidth := len(engineSchematics[part.Y])
	rowCount := len(engineSchematics)

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
		s := engineSchematics[part.Y][leftBound:part.X]

		adjacentLeft = containsSymbol(s)
	}

	// can look to the right on the rights
	if canLookRight {
		s := engineSchematics[part.Y][part.X+part.Width : rightBound]

		adjacentRight = containsSymbol(s)
	}

	// Can look at previous row
	if canLookUp {
		s := engineSchematics[part.Y-1][leftBound:rightBound]

		adjacentUp = containsSymbol(s)
	}

	// can look at next row
	if canLookDown {
		s := engineSchematics[part.Y+1][leftBound:rightBound]

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

func areCoordinatesAdjacent(x1 int, y1 int, x2 int, y2 int) bool {
	return math.Abs(float64(x1-x2)) <= 1 && math.Abs(float64(y1-y2)) <= 1
}
