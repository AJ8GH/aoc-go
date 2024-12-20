package day6

import (
	"fmt"
)

var visited = map[position]bool{}

const (
	n       = "n"
	e       = "e"
	s       = "s"
	w       = "w"
	starter = "^"
)

type position struct {
	i, j int
}

func Level1(input []string) int {
	input = input[:len(input)-1]
	current := position{0, 0}
	for i, row := range input {
		for j, el := range row {
			if string(el) == starter {
				current = position{i, j}
			}
		}
	}

	direction := n
	for {
		visit(current)
		next := nextPosition(current, direction)
		nextVal := get(input, next)
		switch nextVal {
		case "#":
			direction = nextDirection(direction)
			current = nextPosition(current, direction)
		case "":
			return len(visited)
		default:
			current = nextPosition(current, direction)
		}
	}
}

func Level2(input []string) int {
	return 0
}

func nextPosition(p position, direction string) position {
	switch direction {
	case n:
		return position{p.i - 1, p.j}
	case e:
		return position{p.i, p.j + 1}
	case s:
		return position{p.i + 1, p.j}
	case w:
		return position{p.i, p.j - 1}
	default:
		panic(fmt.Sprintf("Bad direction %q", direction))
	}
}

func nextDirection(direction string) string {
	switch direction {
	case n:
		return e
	case e:
		return s
	case s:
		return w
	case w:
		return n
	default:
		panic(fmt.Sprintf("Bad direction %q", direction))
	}
}

func get(input []string, p position) string {
	if p.i < 0 || p.i >= len(input) || p.j < 0 || p.j >= len(input[p.i]) {
		return ""
	}

	return string(input[p.i][p.j])
}

func visit(p position) {
	visited[p] = true
}
