package submarine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var data = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func Test_plotInstructions(t *testing.T) {
	forwards, down := PlotInstructions(data)
	require.Equal(t, 15, forwards)
	require.Equal(t, 10, down)
}

func TestPlotInstructionsAim(t *testing.T) {
	forwards, down := PlotInstructionsAim(data)
	require.Equal(t, 15, forwards)
	require.Equal(t, 60, down)
}
