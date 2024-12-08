package day2

import (
	"math"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`\d+`)

func Level1(input []string) (safeCount int) {
	levels := parse(input)
	diffs := allToDiffs(levels)
	return len(keepValid(diffs))
}

func Level2(input []string) int {
	levels := parse(input)
	diffs := allToDiffs(levels)
	return len(keepValidTolerant(diffs, levels))
}

func allToDiffs(levels [][]int) [][]int {
	out := [][]int{}
	for _, v := range levels {
		out = append(out, toDiffs(v))
	}
	return out
}

func toDiffs(levels []int) []int {
	out := []int{}
	for i, v := range levels {
		if i+1 >= len(levels) {
			continue
		}

		out = append(out, v-levels[i+1])
	}
	return out
}

func parse(input []string) (reports [][]int) {
	for _, v := range input {
		report := []int{}
		nums := re.FindAllString(v, -1)
		for _, v := range nums {
			level, _ := strconv.Atoi(v)
			report = append(report, level)
		}
		if len(report) == 0 {
			continue
		}
		reports = append(reports, report)
	}
	return reports
}

func keepValid(allDiffs [][]int) [][]int {
	out := [][]int{}
	for _, v := range allDiffs {
		if isValid(v) {
			out = append(out, v)
		}
	}
	return out
}

func keepValidTolerant(allDiffs [][]int, levels [][]int) [][]int {
	out := [][]int{}
	for i, v := range allDiffs {
		if isValid(v) {
			out = append(out, v)
		} else {
			levelsToCheck := levels[i]
			for j := range levelsToCheck {
				levelsCopy := make([]int, len(levelsToCheck))
				copy(levelsCopy, levelsToCheck)
				levelsModified := append(levelsCopy[:j], levelsCopy[j+1:]...)
				diffsToCheck := toDiffs(levelsModified)
				if isValid(diffsToCheck) {
					out = append(out, diffsToCheck)
					break
				}
			}
		}
	}
	return out
}

func isValid(diffs []int) bool {
	return all(diffs, func(n int) bool { return math.Abs(float64(n)) <= 3 }) &&
		(all(diffs, func(n int) bool { return n > 0 }) ||
			all(diffs, func(n int) bool { return n < 0 }))
}

func all(diffs []int, predicate func(int) bool) bool {
	for _, v := range diffs {
		if !predicate(v) {
			return false
		}
	}
	return true
}
