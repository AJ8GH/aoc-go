package day8

var locations map[string][]location
var antiNodes map[location]bool

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

	antiNodes = map[location]bool{}
	for _, v := range pairs {
		if level == 2 {
			antiNodes[location{i: v.a.i, j: v.a.j}] = true
			antiNodes[location{i: v.b.i, j: v.b.j}] = true
		}

		diff := location{i: v.a.i - v.b.i, j: v.a.j - v.b.j}
		antiA := location{i: v.a.i + diff.i, j: v.a.j + diff.j}
		antiB := location{i: v.b.i - diff.i, j: v.b.j - diff.j}

		findAntiNodes(antiA, diff, input, level, plus)
		findAntiNodes(antiB, diff, input, level, minus)
	}
	return len(antiNodes)
}

func findAntiNodes(
	antiNode,
	diff location,
	input []string,
	level int,
	op func(int, int) int,
) {
	for {
		if antiNode.i < 0 || antiNode.i >= len(input) || antiNode.j < 0 || antiNode.j >= len(input[antiNode.i]) {
			break
		}
		antiNodes[antiNode] = true
		if level == 1 {
			break
		}
		antiNode = location{i: op(antiNode.i, diff.i), j: op(antiNode.j, diff.j)}
	}
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
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
