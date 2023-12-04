package main

import (
	"fmt"

	"github.com/bLittle1996/advent-of-code-2023/scratchcards/data"
	"github.com/bLittle1996/advent-of-code-2023/scratchcards/scratchcard"
)

func main() {
	scratchcards := scratchcard.NewPile(data.Data)

	fmt.Println(fmt.Sprintf("From all %d scratchcards, you have a total of %d points!", len(scratchcards), scratchcards.Points()))
	fmt.Println("But points aren't everything, in fact that's now how these scratchcards work at all!")
	fmt.Println(fmt.Sprintf("You actually just win more scratchcards, %d in your case, and those scratchcards don't even win you anything useful! Ha!", scratchcards.Process()))
}
