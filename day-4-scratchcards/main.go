package main

import (
	"fmt"

	"github.com/bLittle1996/advent-of-code-2023/scratchcards/data"
	"github.com/bLittle1996/advent-of-code-2023/scratchcards/scratchcard"
)

func main() {
	scratchcards := scratchcard.NewPile(data.Data)
	pointSum := 0

	for _, sc := range scratchcards {
		pointSum += sc.Points()
	}

	fmt.Println(fmt.Sprintf("From all %d cards, you have a total of %d points!", len(scratchcards), pointSum))
}
