package day15

import (
	"github.com/aj8gh/aocgo/util"
)

type dir string
type object string

const (
	u dir = "^"
	r dir = ">"
	d dir = "v"
	l dir = "<"

	robot object = "@"
	wall  object = "#"
	box   object = "O"
	empty object = "."

	multiplier = 100
)

func Level1(input []string) (result int) {
	grid, dirs, loc := parse(input)
	result = solve(grid, dirs, loc)
	return result
}

func Level2(input []string) (result int) {
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
	default:
		panic("Bad object")
	}
}

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

func parse(input []string) (
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
