package diagnostic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

func TestGetRate(t *testing.T) {
	gamma, epsilon := GetRate(input)
	assert.Equal(t, 22, gamma)
	assert.Equal(t, 9, epsilon)
}
