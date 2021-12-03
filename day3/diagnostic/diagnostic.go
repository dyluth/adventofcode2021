package diagnostic

import (
	"fmt"
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
		if oneCount > (len(inputs) / 2) {
			gammaList = append(gammaList, "1")
			epsilonList = append(epsilonList, "0")
		} else {
			gammaList = append(gammaList, "0")
			epsilonList = append(epsilonList, "1")
		}
	}
	gammaString := strings.Join(gammaList, "")
	epsilonString := strings.Join(epsilonList, "")
	return binaryStringToInt(gammaString), binaryStringToInt(epsilonString)
}

func binaryStringToInt(binary string) int {
	val64, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(val64)
}

func GetOxygenRating(inputs []string) int {
	prefix := []string{}
	for i := 0; i < len(inputs[0]); i++ {
		v := getCommonValue(inputs, i, "1")
		prefix = append(prefix, v)
		inputs = subSliceByPrefix(inputs, strings.Join(prefix, ""))
		if len(inputs) == 1 {
			return binaryStringToInt(inputs[0])
		}
	}
	panic(fmt.Sprintf("erm.. still have: %+v\n", inputs))
}
func GetCarbonDioxideRating(inputs []string) int {
	prefix := []string{}
	for i := 0; i < len(inputs[0]); i++ {
		v := getCommonValue(inputs, i, "1")
		if v == "0" {
			v = "1"
		} else {
			v = "0"
		}
		prefix = append(prefix, v)
		inputs = subSliceByPrefix(inputs, strings.Join(prefix, ""))
		if len(inputs) == 1 {
			return binaryStringToInt(inputs[0])
		}
	}
	panic(fmt.Sprintf("erm.. still have: %+v\n", inputs))
}

// take a string and return all slices that start with that prefix
func subSliceByPrefix(inputs []string, prefix string) []string {
	vals := []string{}
	for i := range inputs {
		if strings.HasPrefix(inputs[i], prefix) {
			vals = append(vals, inputs[i])
		}
	}
	return vals
}

// returns the most common rune at that index in the slice of strings
// if there is a tie, it will use the prefer string instead
func getCommonValue(inputs []string, index int, prefer string) string {
	results := make(map[string]int)
	for _, in := range inputs {
		c := in[index : index+1]
		count := results[c]
		results[c] = count + 1
	}

	biggest := ""
	biggestValue := -1
	for val, count := range results {
		if count > biggestValue {
			biggestValue = count
			biggest = val
		} else if count == biggestValue {
			biggest = prefer
		}
	}
	return biggest
}
