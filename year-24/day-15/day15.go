package day15

import (
	"fmt"

	"github.com/aj8gh/aocgo/util"
)

type dir string
type object string

const (
	u dir = "^"
	r dir = ">"
	d dir = "v"
	l dir = "<"

	robot    object = "@"
	wall     object = "#"
	box      object = "O"
	leftBox  object = "["
	rightBox object = "]"
	empty    object = "."

	multiplier = 100
)

func Level1(input []string) (result int) {
	grid, dirs, loc := parse(input, 1)
	result = solve(grid, dirs, loc)
	return result
}

func Level2(input []string) (result int) {
	grid, dirs, loc := parse(input, 2)
	fmt.Sprintf("%v %v %v", grid, dirs, loc)
	return result
}

func solve(grid util.Grid[object], dirs []dir, loc util.Location) (result int) {
	for _, dir := range dirs {
		if move(grid, dir, loc) {
			loc = next(loc, dir)
		}
	}
	return gpsSum(grid)
}

func move(grid util.Grid[object], dir dir, loc util.Location) bool {
	nextLoc := next(loc, dir)
	nextObj := grid.GetLoc(nextLoc)
	switch nextObj {
	case box:
		moved := move(grid, dir, nextLoc)
		if !moved {
			return false
		}
		switchLocs(grid, loc, nextLoc)
		return true
	case wall:
		return false
	case empty:
		switchLocs(grid, loc, nextLoc)
		return true
	case leftBox:
		right := nextSide(loc, dir, r)
		return moveDouble(grid, dir, loc, right)
	case rightBox:
		left := nextSide(loc, dir, l)
		return moveDouble(grid, dir, loc, left)
	default:
		panic("Bad object")
	}
}

func moveDouble(grid util.Grid[object], dir dir, a util.Location, b util.Location) bool {
	nextA := next(a, dir)
	nextB := next(b, dir)
	nextObjA := grid.GetLoc(nextA)
	nextObjB := grid.GetLoc(nextB)

	/*

	[][]
	[]

	*/
	var resultA bool
	switch nextObjA {
	case leftBox:
		return moveDouble(grid, dir, nextA, nextB)
	case rightBox:
		if nextObjA == rightBox {

		}
		resultA = false
	case wall:
		resultA = false
	case empty:
		resultA = false
	default:
		panic("Bad object")
	}

	switch nextObjB {
	case leftBox:
		return resultA && false
	case rightBox:
		return resultA && false
	case wall:
		return resultA && false
	case empty:
		return resultA && false
	default:
		panic("Bad object")
	}
}

func switchLocsDouble() {

}

/*

[][]
[]

[]..
[]

..[]
[]

[][][]
[][]
[]

*/

func switchLocs(grid util.Grid[object], loc, nextLoc util.Location) {
	currentObj := grid.GetLoc(loc)
	grid.PutLoc(nextLoc, currentObj)
	grid.PutLoc(loc, empty)
}

func next(loc util.Location, dir dir) util.Location {
	switch dir {
	case u:
		return util.Location{X: loc.X, Y: loc.Y - 1}
	case r:
		return util.Location{X: loc.X + 1, Y: loc.Y}
	case d:
		return util.Location{X: loc.X, Y: loc.Y + 1}
	case l:
		return util.Location{X: loc.X - 1, Y: loc.Y}
	default:
		panic("Bad dir")
	}
}

func nextSide(loc util.Location, vDir, hDir dir) util.Location {
	var newLoc util.Location

	switch hDir {
	case l:
		newLoc = util.Location{X: loc.X - 1, Y: loc.Y}
	case r:
		newLoc = util.Location{X: loc.X + 1, Y: loc.Y}
	}

	return next(newLoc, vDir)
}

func gpsSum(grid util.Grid[object]) (sum int) {
	for i, row := range grid.Rows {
		for j, o := range row {
			if o == box {
				sum += multiplier*i + j
			}
		}
	}
	return sum
}

func parse(input []string, level int) (
	grid util.Grid[object],
	directions []dir,
	loc util.Location,
) {
	startOfDirs := 0
	rows := [][]object{}
	for i, line := range input {
		if line == "" {
			startOfDirs = i + 1
			break
		}

		row := []object{}
		for j, v := range line {
			o := object(v)
			if o == robot {
				loc.X = j
				loc.Y = i
			}
			if level == 2 {
				switch o {
				case robot:
					o = object(fmt.Sprintf("%v %v", robot, empty))
				case box:
					o = object(fmt.Sprintf("%v %v", leftBox, rightBox))
				case wall, empty:
					o = object(fmt.Sprintf("%v %v", o, o))
				}
			}
			row = append(row, o)
		}
		rows = append(rows, row)
	}
	grid.Rows = rows

	for i := startOfDirs; i < len(input); i++ {
		for _, v := range input[i] {
			directions = append(directions, dir(v))
		}
	}

	return grid, directions, loc
}
