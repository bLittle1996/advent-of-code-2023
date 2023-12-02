package main

import (
	"fmt"
	"strings"

	"github.com/bLittle1996/advent-of-code-2023/day-2-cube-conundrum/data"
	cubegame "github.com/bLittle1996/advent-of-code-2023/day-2-cube-conundrum/internal/cube-game"
)

func main() {
	var games []cubegame.CubeGame
	gameLines := strings.Split(data.Data, "\n")

	for _, line := range gameLines {
		if line == "" {
			continue
		}

		game, err := cubegame.FromString(line)

		if err != nil {
			fmt.Println(err)
			continue
		}

		games = append(games, *game)
	}

	fancyGameBag := cubegame.CubeBag{Red: 12, Green: 13, Blue: 14}

	cumulativePower := 0
	legalGameIdSum := 0
	legalGameIds := []string{}
	for _, game := range games {
		if game.IsLegal(fancyGameBag) {
			legalGameIdSum += game.Id
			legalGameIds = append(legalGameIds, fmt.Sprint(game.Id))
		}

		cumulativePower += game.ComputePower()
	}

	fmt.Println(fmt.Sprintf("The following games were legal for this bag: %v", fancyGameBag))
	fmt.Println(fmt.Sprintf("%s", strings.Join(legalGameIds, ", ")))

	fmt.Println(fmt.Sprintf("\nThe sum of the IDs of the legal games is %d", legalGameIdSum))

	fmt.Println(fmt.Sprintf("\nThe power of all games minimum cube sets is %d", cumulativePower))

}
