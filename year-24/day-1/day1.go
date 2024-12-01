package day1

import (
	"math"
	"regexp"
	"slices"
	"strconv"
)

var re = regexp.MustCompile(`\d+`)

func Level1(input []string) int64 {
	firstNums, lastNums := getNumbers(input)

	slices.Sort(firstNums)
	slices.Sort(lastNums)

	var sum int64 = 0
	for i, v := range firstNums {
		diff := v - lastNums[i]
		sum += int64(math.Abs(float64(diff)))
	}

	return sum
}

func Level2(input []string) int64 {
	return 0
}

func getNumbers(input []string) (first, last []int) {
	for _, v := range input {
		if len(v) == 0 {
			continue
		}

		matches := re.FindAllString(v, -1)
		n, _ := strconv.Atoi(matches[0])
		m, _ := strconv.Atoi(matches[1])
		first = append(first, n)
		last = append(last, m)
	}

	return first, last
}
