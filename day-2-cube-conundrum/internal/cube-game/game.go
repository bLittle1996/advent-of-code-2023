package cubegame

import (
	"fmt"
	"strconv"
	"strings"
)

type CubeGame struct {
	Id      int
	Reveals []Reveal
}

type Reveal struct {
	Red   int
	Green int
	Blue  int
}

type CubeBag struct {
	Red   int
	Green int
	Blue  int
}

func (game *CubeGame) IsLegal(bag CubeBag) bool {
	if game == nil {
		return false
	}

	for _, reveal := range game.Reveals {
		if reveal.Red > bag.Red || reveal.Green > bag.Green || reveal.Blue > bag.Blue {
			return false
		}
	}

	return true
}

func (game *CubeGame) FewestNumberOfCubes() CubeBag {
	bag := CubeBag{}

	if game == nil {
		return bag
	}

	reachOfRed := 0
	greatestGreen := 0
	biggestBlue := 0

	for _, reveal := range game.Reveals {
		if reveal.Red > reachOfRed {
			reachOfRed = reveal.Red
		}

		if reveal.Green > greatestGreen {
			greatestGreen = reveal.Green
		}

		if reveal.Blue > biggestBlue {
			biggestBlue = reveal.Blue
		}
	}

	bag.Red = reachOfRed
	bag.Green = greatestGreen
	bag.Blue = biggestBlue

	return bag
}

func (game *CubeGame) ComputePower() int {
	if game == nil {
		return 0
	}

	minimumBag := game.FewestNumberOfCubes()

	gamePower := minimumBag.Red * minimumBag.Green * minimumBag.Blue

	return gamePower
}

// Creates a game from a string in the following format, individual colours are optional:
// Game %d: %d red, %d green, %d blue; %d red, %d green, %d blue ...
//
// Returns *Game, nil or nil, error.
func FromString(str string) (*CubeGame, error) {
	parts := strings.Split(str, ":")

	if len(parts) != 2 {
		return nil, fmt.Errorf("Unable to parse game from string, too many colons")
	}

	gameIdStr := parts[0]
	revelationsStr := parts[1]

	gameId, err := parseGameId(gameIdStr)

	if err != nil {
		return nil, err
	}

	game := &CubeGame{
		Id:      gameId,
		Reveals: parseGameReveals(revelationsStr),
	}

	return game, nil
}

func parseGameId(gameIdStr string) (int, error) {
	genericError := fmt.Errorf("Unable to parse game id, incorrect format. Expected Game #d, received %s", gameIdStr)
	parts := strings.Split(gameIdStr, " ")

	if len(parts) != 2 {
		return 0, genericError
	}

	gameText, gameId := parts[0], parts[1]

	parsedId, err := strconv.ParseInt(gameId, 10, 0)

	if gameText != "Game" || err != nil {
		return 0, genericError
	}

	return int(parsedId), nil
}

func parseGameReveals(revelationsStr string) []Reveal {
	var reveals []Reveal
	revelationsStr = strings.Trim(revelationsStr, " \n")

	parts := strings.Split(revelationsStr, ";")

	for _, part := range parts {
		reveal, err := parseGameReveal(part)

		if err != nil {
			continue
		}

		reveals = append(reveals, *reveal)
	}

	return reveals
}

func parseGameReveal(revealStr string) (*Reveal, error) {
	reveal := &Reveal{}
	revealStr = strings.Trim(revealStr, " ")
	reveals := strings.Split(revealStr, ",")

	for _, revealColour := range reveals {
		revealColour = strings.Trim(revealColour, " ")

		parts := strings.Split(revealColour, " ")

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid reveal format: expected # colour, received %s", revealColour)
		}

		num, colour := parts[0], parts[1]

		numAsInt, err := strconv.ParseInt(num, 10, 0)

		if err != nil {
			return nil, err
		}

		switch colour {
		case "red":
			reveal.Red = int(numAsInt)
		case "green":
			reveal.Green = int(numAsInt)
		case "blue":
			reveal.Blue = int(numAsInt)
		default:
			return nil, fmt.Errorf("Invalid colour provided, only red, green, and blue are allowed")

		}
	}

	return reveal, nil
}

func contains[T comparable](haystack []T, needle T) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}

	return false
}
