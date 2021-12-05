package vents

import (
	"fmt"
	"strconv"
	"strings"
)

type VentVector struct {
	Coords []coord
}

type coord struct {
	x int
	y int
}

func (c coord) key() string {
	return fmt.Sprintf("%v-%v", c.x, c.y)
}

func BuildBoard(input []string) map[string]int {
	board := make(map[string]int)
	for _, line := range input {
		vv, err := NewVentVector(line)
		if err != nil {
			//panic(err)
			continue // skip diagonals
		}
		for _, c := range vv.Coords {
			key := c.key()
			val, ok := board[key]
			if !ok {
				val = 0
			}
			board[key] = val + 1
		}
	}
	return board
}

func GetOverlaps(board map[string]int) int {
	count := 0
	for _, v := range board {
		if v > 1 {
			count++
		}
	}
	return count
}

// expect in to look like: `0,9 -> 5,9`
func NewVentVector(in string) (VentVector, error) {
	in = strings.TrimSpace(in)
	ends := strings.Split(in, "->")
	if len(ends) != 2 {
		panic(fmt.Sprintf("WTF? %v\n", in))
	}
	start := NewCoord(ends[0])
	end := NewCoord(ends[1])
	coords, err := plotVector(start, end)
	if err != nil {
		return VentVector{}, err
	}
	return VentVector{Coords: coords}, nil
}

// plot out all squares in the vector
func plotVector(start, end coord) ([]coord, error) {
	// first check horizontal or vertical
	// if (start.x == end.x) == (start.y == end.y) {
	// 	if (start.x != end.x) || (start.y != end.y) {
	// 		// error if both change
	// 		// these are bad - diagonals
	// 		return nil, fmt.Errorf("no diagonals")
	// 	}
	// }
	// now build the stepVecrtor
	stepVector := coord{}
	if start.x < end.x {
		stepVector.x = 1
	} else if start.x > end.x {
		stepVector.x = -1
	}

	if start.y < end.y {
		stepVector.y = 1
	} else if start.y > end.y {
		stepVector.y = -1
	}

	// now we have the change, repeatedly apply the step to the start till we hit the end.
	coords := []coord{start}

	latest := start
	for {
		if latest.x == end.x && latest.y == end.y {
			return coords, nil
		}
		next := coord{x: latest.x + stepVector.x,
			y: latest.y + stepVector.y}
		coords = append(coords, next)
		latest = next
	}
}

// expect in to be like `0,9`
func NewCoord(in string) coord {
	in = strings.TrimSpace(in)
	parts := strings.Split(in, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return coord{x: x, y: y}
}
