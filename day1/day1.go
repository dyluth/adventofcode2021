package day1

func countIncreases(values []int) int {
	increases := 0
	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			increases++
		}
	}

	return increases
}

func constructSlidingWindow(values []int) []int {
	result := []int{}
	for i := 2; i < len(values); i++ {
		total := values[i] + values[i-1] + values[i-2]
		result = append(result, total)
	}

	return result
}

func countSlidingIncreases(values []int) int {
	values = constructSlidingWindow(values)
	increases := 0
	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			increases++
		}
	}

	return increases
}
