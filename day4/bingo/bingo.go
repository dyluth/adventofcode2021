package bingo

import (
	"regexp"
	"strconv"
	"strings"
)

type entry struct {
	S     string
	Found bool
}

type board struct {
	e [][]*entry
}

// returns the winning board and number
func DoBingo(input []string, first bool) (*board, string) {
	numbers, boards := Generate(input)

	for _, n := range numbers {
		for i := range boards {
			boards[i].Markoff(n)
			if boards[i].HasWon() {
				if first {
					return boards[i], n
				}
			}
			// yuck - this is nasty, but quick to hack in for part 2
			if !first {
				allWinners := true
				for b := range boards {
					if !boards[b].HasWon() {
						allWinners = false
					}
				}
				if allWinners {
					return boards[i], n
				}
			}
		}
	}
	panic("on the dance floor")
}

func Generate(input []string) (numbers []string, boards []*board) {
	boards = []*board{}
	//
	numbers = strings.Split(input[0], ",")

	for i := 1; i < len(input); i++ {
		lines := strings.Split(input[i], "\n")
		b := [][]*entry{}
		for _, line := range lines {
			line = strings.TrimSpace(line)
			line = regexp.MustCompile(`\s+`).ReplaceAllString(line, "-")

			elements := strings.Split(line, "-")
			entries := []*entry{}
			for i := range elements {
				entries = append(entries, &entry{S: elements[i]})
			}
			b = append(b, entries)
		}

		boards = append(boards, &board{b})
	}
	return numbers, boards
}

func (b board) Markoff(value string) {
	// find all entries with that value and mark it off
	for i := range b.e {
		for j := range b.e[i] {
			if b.e[i][j].S == value {
				b.e[i][j].Found = true
			}
		}
	}
}

func (b board) GetUnmarked() []string {
	results := []string{}
	for i := range b.e {
		for j := range b.e[i] {
			if !b.e[i][j].Found {
				results = append(results, b.e[i][j].S)
			}
		}
	}
	return results
}

func (b board) HasWon() bool {
	// search all vertical  for completely marked off lines
	for i := range b.e {
		found := true
		for j := range b.e[i] {
			if !b.e[i][j].Found {
				found = false
			}
		}
		if found {
			return true
		}
	}
	//search all horizontal for completely marked off lines
	for i := range b.e[0] {
		found := true
		for j := range b.e {
			if !b.e[j][i].Found {
				found = false
			}
		}
		if found {
			return true
		}
	}
	return false
}

func SumStrings(list []string) int {
	sum := 0
	for _, l := range list {
		s, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		sum += s
	}
	return sum
}
