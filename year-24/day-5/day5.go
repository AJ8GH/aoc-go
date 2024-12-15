package day5

import (
	"strconv"
	"strings"
)

func Level1(input []string) int {
	order, pages := parse(input)
	valid, _ := filterByValidOrder(order, pages)
	return sumMiddles(valid)
}

func Level2(input []string) int {
	order, pages := parse(input)
	_, invalid := filterByValidOrder(order, pages)
	sorted := sortAll(invalid, order)
	return sumMiddles(sorted)
}

func sortAll(pages [][]int, order map[int]map[int]bool) (sorted [][]int) {
	for _, v := range pages {
		sorted = append(sorted, sort(v, order))
	}
	return
}

func sort(pages []int, order map[int]map[int]bool) (sorted []int) {
	for i := 0; i < len(pages); i++ {
		page := pages[i]
		if len(sorted) == 0 {
			sorted = append(sorted, page)
		} else {
			for i, current := range sorted {
				isPageLessThanCurrent := order[page][current]
				if isPageLessThanCurrent {
					sortedCopy := make([]int, len(sorted))
					copy(sortedCopy, sorted)
					sorted = append(sorted[:i], page)
					sorted = append(sorted, sortedCopy[i:]...)
					break
				}
				if i == len(sorted)-1 {
					sorted = append(sorted, page)
					break
				}
			}
		}
	}
	return
}

func sumMiddles(pages [][]int) int {
	sum := 0
	for _, v := range pages {
		if len(v) > 0 {
			sum += v[len(v)/2]
		}
	}
	return sum
}

func filterByValidOrder(
	order map[int]map[int]bool,
	pages [][]int,
) (valid [][]int, invalid [][]int) {
	for _, v := range pages {
		pagesToCheck := []int{}
	pageCheck:
		for i, page := range v {
			for _, toCheck := range pagesToCheck {
				isCorrectOrder := order[toCheck][page]
				if !isCorrectOrder {
					invalid = append(invalid, v)
					break pageCheck
				}
			}
			if i == len(v)-1 {
				valid = append(valid, v)
			}
			pagesToCheck = append(pagesToCheck, page)
		}
	}
	return
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
