package util

import (
	"fmt"
)

type Grid[T comparable] struct {
	Rows [][]T
}

func (g *Grid[T]) Get(x, y int) T {
	return g.Rows[y][x]
}

func (g *Grid[T]) GetLoc(l Location) T {
	return g.Rows[l.Y][l.X]
}

func (g *Grid[T]) Put(x, y int, val T) {
	g.Rows[y][x] = val
}

func (g *Grid[T]) PutLoc(l Location, val T) {
	g.Rows[l.Y][l.X] = val
}

func (g *Grid[T]) Height() int {
	return len(g.Rows)
}

func (g *Grid[T]) Width() int {
	if g.Height() == 0 {
		return 0
	}
	return len(g.Rows[0])
}

func (g Grid[T]) String() string {
	s := ""

	for _, row := range g.Rows {
		line := ""
		for _, v := range row {
			line += fmt.Sprintf("%v ", v)
		}
		s += fmt.Sprintf("%v%v", line, "\n")
	}

	return s
}
