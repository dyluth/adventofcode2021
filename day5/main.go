package main

import (
	"fmt"

	"github.com/dyluth/adventofcode2021/day5/vents"
	"github.com/dyluth/adventofcode2021/selectionbox"
)

func main() {
	lines := selectionbox.ReadInput()
	board := vents.BuildBoard(lines)
	overlaps := vents.GetOverlaps(board)
	fmt.Printf("overlaps: %v\n", overlaps)
}
