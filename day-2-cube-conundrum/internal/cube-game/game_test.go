package cubegame_test

import (
	"strings"
	"testing"

	"github.com/bLittle1996/advent-of-code-2023/day-2-cube-conundrum/data"
	cubegame "github.com/bLittle1996/advent-of-code-2023/day-2-cube-conundrum/internal/cube-game"
)

var testBag = cubegame.CubeBag{
	Red:   12,
	Green: 13,
	Blue:  14,
}

const testInputAnswer = 8

func TestThatItSolvesTheExample(t *testing.T) {
	gameLines := strings.Split(data.TestData, "\n")
	games := []cubegame.CubeGame{}

	for _, line := range gameLines {
		if line == "" {
			continue
		}
		game, _ := cubegame.FromString(line)

		games = append(games, *game)
	}

	legalIdSum := 0
	for _, game := range games {
		if game.IsLegal(testBag) {
			legalIdSum += game.Id
		}
	}

	if legalIdSum != testInputAnswer {
		t.Errorf("Expected %d, received %d", testInputAnswer, legalIdSum)
		return
	}
}
