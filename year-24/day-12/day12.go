package day12

type direction string

type location struct {
	val         string
	i, j, edges int
	dir         direction
}

type edge struct {
	first, last location
	dir         direction
}

const (
	n direction = "n"
	e direction = "e"
	s direction = "s"
	w direction = "w"
)

func Level1(input []string) int {
	locs, _ := parse(input)
	return calculatePrices(locs)
}

func Level2(input []string) (price int) {
	locs, edges := parse(input)
	for i, v := range locs {
		price += len(v) * len(edges[i])
	}
	return price
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

func parse(input []string) (locs [][]*location, edges [][]edge) {
	found := map[location]bool{}
	edgesFound := map[location]bool{}
	for i, row := range input {
		for j, v := range row {
			val := string(v)
			loc := location{i: i, j: j, val: val}
			if found[location{i: i, j: j}] {
				continue
			}
			found[location{i: i, j: j}] = true
			locsToAppend, edgesToAppend := findAll(input, []*location{&loc}, &loc, found, edgesFound, []edge{})
			locs = append(locs, locsToAppend)
			edges = append(edges, edgesToAppend)
		}
	}
	return locs, edges
}

func findAll(
	input []string,
	locs []*location,
	startLoc *location,
	found map[location]bool,
	edgesFound map[location]bool,
	edges []edge,
) ([]*location, []edge) {
	for _, v := range []location{
		{i: startLoc.i + 1, j: startLoc.j, dir: s},
		{i: startLoc.i - 1, j: startLoc.j, dir: n},
		{i: startLoc.i, j: startLoc.j + 1, dir: e},
		{i: startLoc.i, j: startLoc.j - 1, dir: w},
	} {
		loc := get(input, v.i, v.j)
		if loc.val != startLoc.val {
			startLoc.edges++
			if !edgesFound[location{i: startLoc.i, j: startLoc.j, dir: v.dir}] {
				edge := findEdge(input, startLoc, v.dir, edgesFound)
				edges = append(edges, edge)
			}
			continue
		}

		if found[location{i: v.i, j: v.j}] {
			continue
		}

		found[location{i: v.i, j: v.j}] = true
		locs = append(locs, loc)
		locs, edges = findAll(input, locs, loc, found, edgesFound, edges)
	}

	return locs, edges
}

func findEdge(input []string, startLoc *location, dir direction, edgesFound map[location]bool) edge {
	var startDir, endDir direction
	switch dir {
	case n, s:
		startDir, endDir = w, e
	case e, w:
		startDir, endDir = s, n
	}

	edgesFound[location{i: startLoc.i, j: startLoc.j, dir: dir}] = true
	first := getCornerLoc(input, startLoc, dir, startDir, edgesFound)
	last := getCornerLoc(input, startLoc, dir, endDir, edgesFound)
	return edge{first: first, last: last, dir: dir}
}

func getCornerLoc(
	input []string,
	startLoc *location,
	dir, traversalDir direction,
	edgesFound map[location]bool,
) location {
	loc := startLoc
	next := nextLoc(input, loc, traversalDir)
	for next.val == loc.val && nextLoc(input, next, dir).val != loc.val {
		loc = next
		edgesFound[location{i: loc.i, j: loc.j, dir: dir}] = true
		next = nextLoc(input, loc, traversalDir)
	}
	return location{i: loc.i, j: loc.j, val: loc.val, edges: loc.edges, dir: loc.dir}
}

func get(input []string, i, j int) *location {
	if i < 0 || i >= len(input) || j < 0 || j >= len(input[i]) {
		return &location{}
	}

	return &location{i: i, j: j, val: string(input[i][j])}
}

func nextLoc(input []string, loc *location, dir direction) *location {
	switch dir {
	case n:
		return get(input, loc.i-1, loc.j)
	case e:
		return get(input, loc.i, loc.j+1)
	case s:
		return get(input, loc.i+1, loc.j)
	case w:
		return get(input, loc.i, loc.j-1)
	default:
		panic("bad direction")
	}
}
