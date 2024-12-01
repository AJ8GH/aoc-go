package day1

import (
	"math"
	"regexp"
	"slices"
	"strconv"
)

var re = regexp.MustCompile(`\d+`)

func Level1(input []string) int {
	first, last := getNumbers(input)

	slices.Sort(first)
	slices.Sort(last)

	sum := 0
	for i, v := range first {
		diff := v - last[i]
		sum += int(math.Abs(float64(diff)))
	}

	return sum
}

func Level2(input []string) int {
	first, last := getNumbers(input)
	count := map[int]int{}

	for _, v := range last {
		count[v]++
	}

	sum := 0
	for _, v := range first {
		sum += v * count[v]
	}
	return sum
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
