package day5

import (
	"strconv"
	"strings"
)

func Level1(input []string) int {
	count := 0
	order, pages := parse(input)
	for _, v := range pages {

		pagesToCheck := []int{}

	pageCheck:
		for i, page := range v {
			for _, toCheck := range pagesToCheck {
				isCorrectOrder := order[toCheck][page]
				if !isCorrectOrder {
					break pageCheck
				}
			}
			if i == len(v)-1 {
				count += v[len(v)/2]
				break
			}
			pagesToCheck = append(pagesToCheck, page)
		}
	}
	return count
}

func Level2(input []string) int {
	return 0
}

func parse(input []string) (order map[int]map[int]bool, pages [][]int) {
	order = map[int]map[int]bool{}

	for i, v := range input {
		if len(v) == 0 {
			pages = parsePages(i+1, input)
			break
		}

		spl := strings.Split(v, "|")
		a, _ := strconv.Atoi(spl[0])
		b, _ := strconv.Atoi(spl[1])
		if order[a] == nil {
			order[a] = map[int]bool{}
		}
		order[a][b] = true
	}
	return
}

func parsePages(line int, input []string) (pages [][]int) {
	for i := line; i < len(input); i++ {
		out := []int{}
		for _, v := range strings.Split(input[i], ",") {
			n, _ := strconv.Atoi(v)
			out = append(out, n)
		}
		pages = append(pages, out)
	}
	return
}
