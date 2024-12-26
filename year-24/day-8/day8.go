package day8

var locations map[string][]location
var antinodes map[location]bool

type location struct {
	i, j int
	val  string
}

type pair struct {
	a, b location
}

func Level1(input []string) int {
	return solve(input, 1)
}

func Level2(input []string) int {
	return solve(input, 2)
}

func solve(input []string, level int) int {
	input = input[:len(input)-1]
	parse(input)
	pairs := pairs()
	pairMap := map[string][]pair{}
	for _, v := range pairs {
		if pairMap[v.a.val] == nil {
			pairMap[v.a.val] = []pair{v}
		} else {
			pairMap[v.a.val] = append(pairMap[v.a.val], v)
		}
	}

	antinodes = map[location]bool{}
	for _, v := range pairs {
		if level == 2 {
			antinodes[location{i: v.a.i, j: v.a.j}] = true
			antinodes[location{i: v.b.i, j: v.b.j}] = true
		}

		diff := location{i: v.a.i - v.b.i, j: v.a.j - v.b.j}
		antiA := location{i: v.a.i + diff.i, j: v.a.j + diff.j}
		antiB := location{i: v.b.i - diff.i, j: v.b.j - diff.j}

		for {
			if antiA.i < 0 || antiA.i >= len(input) || antiA.j < 0 || antiA.j >= len(input[antiA.i]) {
				break
			}
			antinodes[antiA] = true
			runes := []rune(input[antiA.i])
			runes[antiA.j] = '#'
			input[antiA.i] = string(runes)
			if level == 1 {
				break
			}
			antiA = location{i: antiA.i + diff.i, j: antiA.j + diff.j}
		}
		for {
			if antiB.i < 0 || antiB.i >= len(input) || antiB.j < 0 || antiB.j >= len(input[antiB.i]) {
				break
			}
			antinodes[antiB] = true
			runes := []rune(input[antiB.i])
			runes[antiB.j] = '#'
			input[antiB.i] = string(runes)
			if level == 1 {
				break
			}
			antiB = location{i: antiB.i - diff.i, j: antiB.j - diff.j}
		}
	}
	return len(antinodes)
}

func pairs() (pairs []pair) {
	for _, v := range locations {
		for i, a := range v {
			for j := i + 1; j < len(v); j++ {
				b := v[j]
				if a == b {
					continue
				}
				pairs = append(pairs, pair{a: a, b: b})
			}
		}
	}
	return
}

func parse(input []string) {
	locations = map[string][]location{}
	for i, row := range input {
		for j, v := range row {
			s := string(v)
			if s == "." {
				continue
			}

			if locations[s] == nil {
				locations[s] = []location{{i: i, j: j, val: s}}
			} else {
				slice := locations[s]
				slice = append(slice, location{i: i, j: j, val: s})
				locations[s] = slice
			}
		}
	}
}
