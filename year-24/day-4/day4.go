package day4

import (
	"unicode/utf8"
)

const (
	x = "X"
	m = "M"
	a = "A"
	s = "S"

	north     = "n"
	northEast = "ne"
	east      = "e"
	southEast = "se"
	south     = "s"
	southWest = "sw"
	west      = "w"
	northWest = "nw"
)

var directions = []string{
	north,
	northEast,
	east,
	southEast,
	south,
	southWest,
	west,
	northWest,
}

func Level1(input []string) int {
	xLocs := [][]int{}
	for i, line := range input {
		for j, char := range line {
			if string(char) == x {
				xLocs = append(xLocs, []int{i, j})
			}
		}
	}
	return countXmas(xLocs, input)
}

func Level2(input []string) int {
	return 0
}

func countXmas(xLocs [][]int, input []string) int {
	xmasCount := 0
	for _, loc := range xLocs {
		for _, dir := range directions {
			xmasCount += countForLoc(loc, input, x, dir)
		}
	}
	return xmasCount
}

func countForLoc(loc []int, input []string, letter string, direction string) int {
	nextLoc := nextLoc(loc, direction)
	nextLetter := getLetter(nextLoc, input)
	if nextLetter == next(letter) {
		if nextLetter == s {
			return 1
		}
		return countForLoc(nextLoc, input, nextLetter, direction)
	}
	return 0
}

func nextLoc(loc []int, direction string) []int {
	switch direction {
	case north:
		return []int{loc[0] + 1, loc[1]}
	case northEast:
		return []int{loc[0] + 1, loc[1] + 1}
	case east:
		return []int{loc[0], loc[1] + 1}
	case southEast:
		return []int{loc[0] - 1, loc[1] + 1}
	case south:
		return []int{loc[0] - 1, loc[1]}
	case southWest:
		return []int{loc[0] - 1, loc[1] - 1}
	case west:
		return []int{loc[0], loc[1] - 1}
	case northWest:
		return []int{loc[0] + 1, loc[1] - 1}
	default:
		panic("Bad direction")
	}
}

func next(r string) string {
	switch r {
	case x:
		return m
	case m:
		return a
	case a:
		return s
	default:
		return ""
	}
}

func getLetter(loc []int, input []string) string {
	i, j := loc[0], loc[1]
	if i >= 0 && i < len(input) {
		line := input[i]
		if j >= 0 && j < utf8.RuneCountInString(line) {
			return string(line[j])
		}
	}

	return ""
}
