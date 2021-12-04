package bingo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var input = []string{
	"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
	`22 13 17 11  0
	8  2 23  4 24
   21  9 14 16  7
	6 10  3 18  5
	1 12 20 15 19`,
	` 3 15  0  2 22
	9 18 13 17  5
   19  8  7 25 23
   20 11 10 24  4
   14 21 16 12  6`,
	`14 21 17 24  4
   10 16 15  9 19
   18  8 23 26 20
   22 11 13  6  5
	2  0 12  3  7`,
}

func TestGenerate(t *testing.T) {
	numbers, boards := Generate(input)
	require.Equal(t, numbers[0], "7")
	require.Equal(t, numbers[4], "11")
	require.Equal(t, 3, len(boards))
	require.Equal(t, "22", boards[0].e[0][0].S)
	require.Equal(t, "13", boards[0].e[0][1].S)
	require.Equal(t, "7", boards[0].e[2][4].S)
	require.Equal(t, "18", boards[1].e[1][1].S)
	require.Equal(t, "3", boards[2].e[4][3].S)
}

func Test_board_HasWon(t *testing.T) {
	b := board{
		e: [][]*entry{
			{
				{Found: true},
				{Found: false},
			},
			{
				{Found: false},
				{Found: true},
			},
		},
	}
	require.Equal(t, false, b.HasWon())

	b = board{
		e: [][]*entry{
			{
				{Found: true},
				{Found: true},
			},
			{
				{Found: false},
				{Found: false},
			},
		},
	}
	require.Equal(t, true, b.HasWon())

	b = board{
		e: [][]*entry{
			{

				{Found: false},
				{Found: true},
			},
			{
				{Found: false},
				{Found: true},
			},
		},
	}
	require.Equal(t, true, b.HasWon())

}

func TestDoBingo(t *testing.T) {

	board, score := DoBingo(input, true)
	fmt.Printf("board: %v, score: %v\n", board.GetUnmarked(), score)

	list := board.GetUnmarked()
	sum := SumStrings(list)
	fmt.Printf("board: %v, score: %v sum: %v\n", board.GetUnmarked(), score, sum)
	assert.Equal(t, "24", score)
	assert.Equal(t, 188, sum)

	board, score = DoBingo(input, false)
	fmt.Printf("board2: %v, score: %v\n", board.GetUnmarked(), score)

	list = board.GetUnmarked()
	sum = SumStrings(list)
	fmt.Printf("board2: %v, score: %v sum: %v\n", board.GetUnmarked(), score, sum)
	assert.Equal(t, "13", score)
	assert.Equal(t, 148, sum)

}

func Test_board_GetUnmarked(t *testing.T) {
	b := board{
		e: [][]*entry{
			{
				{Found: true, S: "f1"},
				{Found: false, S: "u1"},
			},
			{
				{Found: false, S: "u2"},
				{Found: true, S: "f2"},
			},
		},
	}
	r := b.GetUnmarked()
	fmt.Printf("results: %+v'n", r)
	require.Len(t, r, 2)

	require.Equal(t, r[0], "u1")
	require.Equal(t, r[1], "u2")
}
