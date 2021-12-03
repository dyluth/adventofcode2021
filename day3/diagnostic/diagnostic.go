package diagnostic

import (
	"strconv"
	"strings"
)

func GetRate(inputs []string) (gamma, epsilon int) {
	gammaList := []string{}
	epsilonList := []string{}

	for i := 0; i < len(inputs[0]); i++ {
		oneCount := 0
		for _, in := range inputs {
			// subslice in to get the i'th character
			c := in[i : i+1]
			// count 1s
			if c == "1" {
				oneCount++
			}
		}
		// maaan this code...
		if oneCount > (len(inputs[0]) / 2) {
			gammaList = append(gammaList, "1")
			epsilonList = append(epsilonList, "0")
		} else {
			gammaList = append(gammaList, "0")
			epsilonList = append(epsilonList, "1")
		}
	}
	gammaString := strings.Join(gammaList, "")
	gamma64, err := strconv.ParseInt(gammaString, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilonString := strings.Join(epsilonList, "")
	epsilon64, err := strconv.ParseInt(epsilonString, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(gamma64), int(epsilon64)
}
