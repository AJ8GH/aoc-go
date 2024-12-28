package day12

type location struct {
	val         string
	i, j, edges int
}

func Level1(input []string) int {
	locs := parse(input)
	return calculatePrices(locs)
}

func Level2(input []string) int {
	return 0
}

func calculatePrices(locs [][]*location) (price int) {
	for _, v := range locs {
		price += calculatePrice(v)
	}
	return price
}

func calculatePrice(block []*location) (price int) {
	perimeter := 0
	for _, v := range block {
		perimeter += v.edges
	}
	return len(block) * perimeter
}

func parse(input []string) (locs [][]*location) {
	found := map[location]bool{}
	for i, row := range input {
		for j, v := range row {
			val := string(v)
			loc := location{i: i, j: j, val: val}
			if found[location{i: i, j: j}] {
				continue
			}
			found[location{i: i, j: j}] = true
			locs = append(locs, findAll(input, []*location{&loc}, &loc, found))
		}
	}
	return locs
}

func findAll(
	input []string,
	locs []*location,
	startLoc *location,
	found map[location]bool,
) []*location {

	for _, v := range [][]int{
		{startLoc.i + 1, startLoc.j},
		{startLoc.i - 1, startLoc.j},
		{startLoc.i, startLoc.j + 1},
		{startLoc.i, startLoc.j - 1},
	} {
		loc := get(input, v[0], v[1])
		if loc.val != startLoc.val {
			startLoc.edges++
			continue
		}

		if found[location{i: v[0], j: v[1]}] {
			continue
		}

		found[location{i: v[0], j: v[1]}] = true
		locs = append(locs, &loc)
		locs = findAll(input, locs, &loc, found)
	}

	return locs
}

func get(input []string, i, j int) location {
	if i < 0 || i >= len(input) || j < 0 || j >= len(input[i]) {
		return location{}
	}

	return location{i: i, j: j, val: string(input[i][j])}
}
