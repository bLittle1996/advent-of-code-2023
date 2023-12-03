package main

import (
	"fmt"

	"github.com/bLittle1996/advent-of-code-2023/gear-ratios/data"
	"github.com/bLittle1996/advent-of-code-2023/gear-ratios/internal/gears"
)

func main() {

	fmt.Println("Here's your engine schematic:")
	fmt.Print(data.Data)

	engine := gears.NewEngine(data.Data)

	fmt.Println(fmt.Sprintf("The sum of part numbers is %d", engine.GetPartNumberSum()))
	fmt.Println(fmt.Sprintf("The combined gear ratio is %d", engine.GetGearRatioSum()))
}
