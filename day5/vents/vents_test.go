package vents

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewCoord(t *testing.T) {
	coord := NewCoord("1,2 ")
	assert.Equal(t, 1, coord.x)
	assert.Equal(t, 2, coord.y)
	coord = NewCoord("   3,4 ")
	assert.Equal(t, 3, coord.x)
	assert.Equal(t, 4, coord.y)
}
func Test_plotVector(t *testing.T) {
	//plotVector()
	start := NewCoord("1,2")
	endDown := NewCoord("1,4")
	endAcross := NewCoord("3,2")
	v, err := plotVector(start, start)
	require.NoError(t, err)
	require.Len(t, v, 1)
	assert.Equal(t, v[0].x, 1)
	assert.Equal(t, v[0].y, 2)

	v, err = plotVector(start, endDown)
	require.NoError(t, err)
	require.Len(t, v, 3)
	assert.Equal(t, v[0].x, 1)
	assert.Equal(t, v[1].x, 1)
	assert.Equal(t, v[2].x, 1)
	assert.Equal(t, v[0].y, 2)
	assert.Equal(t, v[1].y, 3)
	assert.Equal(t, v[2].y, 4)

	v, err = plotVector(start, endAcross)
	require.NoError(t, err)
	require.Len(t, v, 3)
	assert.Equal(t, v[0].x, 1)
	assert.Equal(t, v[1].x, 2)
	assert.Equal(t, v[2].x, 3)
	assert.Equal(t, v[0].y, 2)
	assert.Equal(t, v[1].y, 2)
	assert.Equal(t, v[2].y, 2)

}

var input = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

var inputSmall = []string{"0,1 -> 0,2",
	"1,1 -> 2,1"}

func TestBuildBoard(t *testing.T) {

	board := BuildBoard(inputSmall)
	require.Len(t, board, 4)
}
