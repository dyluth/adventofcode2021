package lanternfish

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSea(t *testing.T) {
	s := NewSea("3,4,3,1,2")
	assert.Equal(t, 0, s.FishOfAge[0])
	assert.Equal(t, 1, s.FishOfAge[1])
	assert.Equal(t, 1, s.FishOfAge[2])
	assert.Equal(t, 2, s.FishOfAge[3])
	assert.Equal(t, 1, s.FishOfAge[4])
}

func TestNextDay(t *testing.T) {
	s := NewSea("3,4,3,1,2")
	s.NextDay()
	assert.Equal(t, 1, s.FishOfAge[0])
	assert.Equal(t, 1, s.FishOfAge[1])
	assert.Equal(t, 2, s.FishOfAge[2])
	assert.Equal(t, 1, s.FishOfAge[3])
	assert.Equal(t, 0, s.FishOfAge[4])
	s.NextDay()
	assert.Equal(t, 1, s.FishOfAge[0])
	assert.Equal(t, 2, s.FishOfAge[1])
	assert.Equal(t, 1, s.FishOfAge[2])
	assert.Equal(t, 0, s.FishOfAge[3])
	assert.Equal(t, 0, s.FishOfAge[4])
	assert.Equal(t, 1, s.FishOfAge[6])
	assert.Equal(t, 1, s.FishOfAge[8])
}

func TestSea_NextDays(t *testing.T) {
	s := NewSea("3,4,3,1,2")
	s.NextDays(18)
	require.Equal(t, 26, s.CountFish())

	s = NewSea("3,4,3,1,2")
	s.NextDays(80)
	require.Equal(t, 5934, s.CountFish())

}
