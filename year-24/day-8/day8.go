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
		diff := location{i: v.a.i - v.b.i, j: v.a.j - v.b.j}
		antiA := location{i: v.a.i + diff.i, j: v.a.j + diff.j}
		antiB := location{i: v.b.i - diff.i, j: v.b.j - diff.j}

		if antiA.i >= 0 && antiA.i < len(input) && antiA.j >= 0 && antiA.j < len(input[antiA.i]) {
			antinodes[antiA] = true
		}
		if antiB.i >= 0 && antiB.i < len(input) && antiB.j >= 0 && antiB.j < len(input[antiB.i]) {
			antinodes[antiB] = true
		}
	}
	return len(antinodes)
}

func Level2(input []string) int {
	return 0
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
