package day3

import (
	"regexp"
	"strconv"
)

const (
	DO     = "do()"
	DO_NOT = "don't()"
)

var (
	re  = regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)
	num = regexp.MustCompile(`\d+`)
)

func Level1(input []string) int {
	return solve(input, false)
}

func Level2(input []string) int {
	return solve(input, true)
}

func solve(input []string, filter bool) int {
	out := 0
	for _, v := range parse(input, filter) {
		out += v[0] * v[1]
	}
	return out
}

func parse(input []string, filter bool) [][]int {
	out := [][]int{}
	doing := true

	for _, v := range input {
		for _, w := range re.FindAllString(v, -1) {
			switch {
			case w == DO:
				doing = true
			case filter && w == DO_NOT:
				doing = false
			case doing && w != DO && w != DO_NOT && len(w) > 0:
				nums := []int{}
				for _, x := range num.FindAllString(w, 2) {
					n, _ := strconv.Atoi(x)
					nums = append(nums, n)
				}
				out = append(out, nums)
			}
		}
	}
	return out
}
