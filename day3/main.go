package main

import (
	"fmt"

	"github.com/dyluth/adventofcode2021/day3/diagnostic"
	"github.com/dyluth/adventofcode2021/selectionbox"
)

func main() {
	lines := selectionbox.ReadInput()
	gamma, epsilon := diagnostic.GetRate(lines)
	fmt.Printf("part1 result: %v\n", gamma*epsilon)

	o2 := diagnostic.GetOxygenRating(lines)
	co2 := diagnostic.GetCarbonDioxideRating(lines)

	fmt.Printf("part2 result: %v\n", o2*co2)
}
