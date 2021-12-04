package main

import (
	"fmt"
	"strconv"

	"github.com/dyluth/adventofcode2021/day4/bingo"
	"github.com/dyluth/adventofcode2021/selectionbox"
)

func main() {
	lines := selectionbox.ReadGroupedInput()
	board, score := bingo.DoBingo(lines, true)
	unmarked := board.GetUnmarked()
	sum := bingo.SumStrings(unmarked)
	s, err := strconv.Atoi(score)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part1 s: %v, sum: %v total: %v", s, sum, s*sum)

	board, score = bingo.DoBingo(lines, false)
	unmarked = board.GetUnmarked()
	sum = bingo.SumStrings(unmarked)
	s, err = strconv.Atoi(score)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part2 s: %v, sum: %v total: %v", s, sum, s*sum)
}
