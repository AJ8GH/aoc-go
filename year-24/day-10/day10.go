package day10

import "strconv"

type location struct {
	i, j, val int
}

func Level1(input []string) (count int) {
	rows, trailheads := parse(input)
	for _, v := range trailheads {
		stack := []location{v}
		for stack[0].val != 9 {
			stack = getNexts(stack, rows)
		}
		ends := map[location]bool{}
		for _, v := range stack {
			ends[v] = true
		}
		count += len(ends)
	}

	return
}

func Level2(input []string) int {
	return 0
}

func getNexts(locs []location, rows [][]int) (out []location) {
	for _, v := range locs {
		out = append(out, nexts(v, rows)...)
	}
	return
}

func nexts(l location, rows [][]int) (nexts []location) {
	for _, v := range [][]int{
		{l.i + 1, l.j},
		{l.i - 1, l.j},
		{l.i, l.j + 1},
		{l.i, l.j - 1},
	} {
		if v[0] >= 0 && v[0] < len(rows) && v[1] >= 0 && v[1] < len(rows[v[0]]) {
			next := get(v[0], v[1], rows)
			if next.val == l.val+1 {
				nexts = append(nexts, next)
			}
		}
	}
	return
}

func get(i, j int, rows [][]int) location {
	return location{i: i, j: j, val: rows[i][j]}
}

func parse(input []string) (rows [][]int, trailheads []location) {
	for i, v := range input[:len(input)-1] {
		row := []int{}
		for j, n := range v {
			digit, _ := strconv.Atoi(string(n))
			row = append(row, digit)
			if digit == 0 {
				trailheads = append(trailheads, location{i: i, j: j, val: 0})
			}
		}
		rows = append(rows, row)
	}
	return
}
