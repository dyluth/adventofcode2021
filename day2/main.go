package main

import (
	"fmt"

	"github.com/dyluth/adventofcode2021/day2/submarine"
	"github.com/dyluth/adventofcode2021/selectionbox"
)

func main() {
	lines := selectionbox.ReadInput()
	forwards, down := submarine.PlotInstructions(lines)
	fmt.Printf("part1 result: %v\n", forwards*down)
	forwards, down = submarine.PlotInstructionsAim(lines)
	fmt.Printf("part2 result: %v\n", forwards*down)
}
