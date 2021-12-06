package main

import (
	"fmt"

	"github.com/dyluth/adventofcode2021/day6/lanternfish"
	"github.com/dyluth/adventofcode2021/selectionbox"
)

func main() {
	lines := selectionbox.ReadInput()
	sea := lanternfish.NewSea(lines[0])
	sea.NextDays(80)

	fmt.Printf("part1 total fish: %v\n", sea.CountFish())
	sea.NextDays(256 - 80)
	fmt.Printf("part2 total fish: %v\n", sea.CountFish())
}
