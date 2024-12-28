package day11

import (
	"strconv"
	"strings"
)

const (
	level1Blinks = 25
	level2Blinks = 75
	multiplier   = 2024
	zero         = "0"
	one          = "1"
	delimiter    = " "
)

func Level1(input []string) int {
	return solve(parse(input), level1Blinks)
}

func Level2(input []string) (count int) {
	return solve(parse(input), level2Blinks)
}

func solve(stones map[string]int, blinks int) (count int) {
	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}

	for _, v := range stones {
		count += v
	}

	return count
}

func blink(stones map[string]int) (afterBlink map[string]int) {
	afterBlink = make(map[string]int, len(afterBlink))
	for k, v := range stones {
		switch {
		case k == zero:
			afterBlink[one] += v
		case len(k)%2 == 0:
			a, b := k[:len(k)/2], k[len(k)/2:]
			for len(b) > 1 && string(b[0]) == zero {
				b = b[1:]
			}
			afterBlink[a] += v
			afterBlink[b] += v
		default:
			n, _ := strconv.Atoi(k)
			afterBlink[strconv.Itoa(n*multiplier)] += v
		}
	}
	return afterBlink
}

func parse(input []string) (stones map[string]int) {
	stones = map[string]int{}
	for _, v := range strings.Split(input[0], delimiter) {
		stones[v]++
	}
	return stones
}
