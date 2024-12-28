package day11

import (
	"strconv"
	"strings"
)

const (
	blinks     = 25
	multiplier = 2024
	zero       = "0"
	one        = "1"
	delimiter  = " "
)

func Level1(input []string) int {
	stones := parse(input)

	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}

	return len(stones)
}

func Level2(input []string) int {
	return 0
}

func blink(stones []string) (afterBlink []string) {
	for _, v := range stones {
		switch {
		case v == zero:
			afterBlink = append(afterBlink, one)
		case len(v)%2 == 0:
			a, b := v[:len(v)/2], v[len(v)/2:]
			for len(b) > 1 && string(b[0]) == zero {
				b = b[1:]
			}
			afterBlink = append(afterBlink, a, b)
		default:
			n, _ := strconv.Atoi(v)
			afterBlink = append(afterBlink, strconv.Itoa(n*multiplier))
		}
	}
	return afterBlink
}

func parse(input []string) (stones []string) {
	return strings.Split(input[0], delimiter)
}
