package day2

import (
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`\d+`)

func Level1(input []string) (safeCount int) {
	return countSafe(input, 0)
}

func Level2(input []string) int {
	return countSafe(input, 1)
}

func countSafe(input []string, tolerance int) (safeCount int) {
	reports := parse(input)
	for _, report := range reports {

		if len(report) < 2 {
			safeCount++
			continue
		}

		if report[0] == report[1] {
			continue
		}

		increasing := false
		if report[0] < report[1] {
			increasing = true
		}

		safe := true
		for i := 1; i < len(report); i++ {
			a, b := report[i-1], report[i]
			if !check(a, b, increasing) {

				if tolerance > 0 {
					if i == len(report)-1 {
						continue
					}

					if check(a, report[i+1], increasing) {
						i++
						continue
					}

					// [1, 5]
				}
				safe = false
				break
			}
		}

		if safe {
			safeCount++
		}
	}

	return safeCount
}

func checkReport(report []int) {

}

func check(a, b int, increasing bool) bool {
	return a != b &&
		a < b == increasing &&
		(increasing && b-a <= 3 || !increasing && a-b <= 3)
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
