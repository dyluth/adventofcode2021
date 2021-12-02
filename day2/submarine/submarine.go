package submarine

import (
	"fmt"
	"strconv"
	"strings"
)

func PlotInstructions(instructions []string) (forwards, down int) {
	for _, line := range instructions {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			amount, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			switch parts[0] {
			case "up":
				down = down - amount
			case "down":
				down = down + amount
			case "forward":
				forwards = forwards + amount
			default:
				panic(fmt.Sprintf("dont know what to do with instruction `%v`", line))
			}
		}
	}
	return
}

func PlotInstructionsAim(instructions []string) (forwards, down int) {
	aim := 0
	for _, line := range instructions {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			amount, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			switch parts[0] {
			case "up":
				aim = aim - amount
			case "down":
				aim = aim + amount
			case "forward":
				forwards = forwards + amount
				down = down + (aim * amount)
			default:
				panic(fmt.Sprintf("dont know what to do with instruction `%v`", line))
			}
		}
	}
	return
}
