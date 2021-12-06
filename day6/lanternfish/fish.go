package lanternfish

import (
	"strconv"
	"strings"
)

type Sea struct {
	FishOfAge []int
}

func NewSea(input string) Sea {
	s := Sea{
		FishOfAge: make([]int, 9),
	}
	input = strings.TrimSpace(input)
	items := strings.Split(input, ",")
	for _, entry := range items {
		val, err := strconv.Atoi(entry)
		if err != nil {
			panic(err)
		}
		s.FishOfAge[val] += 1
	}
	return s
}

func (s *Sea) NextDay() {
	newDay := make([]int, 9)
	for i := 1; i < len(s.FishOfAge); i++ {
		newDay[i-1] = s.FishOfAge[i]
	}
	newDay[8] += s.FishOfAge[0]
	newDay[6] += s.FishOfAge[0]
	s.FishOfAge = newDay
}

func (s *Sea) NextDays(dayCount int) {
	for i := 0; i < dayCount; i++ {
		s.NextDay()
	}
}

func (s *Sea) CountFish() int {
	sum := 0
	for _, val := range s.FishOfAge {
		sum += val
	}
	return sum
}
