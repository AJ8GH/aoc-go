package day14

import (
	"regexp"
	"strconv"
)

const (
	seconds = 100
)

var re = regexp.MustCompile(`\-?\d+`)

type loc struct {
	x, y int
}

type robot struct {
	p, v loc
}

func Level1(input []string) (result int) {
	robots := parse(input)
	quads := []int{0, 0, 0, 0}
	height, width := max(robots)

	for _, r := range robots {
		loc := simulate(r, height, width)
		quad := quadrant(loc, height, width)
		if quad >= 0 {
			quads[quad]++
		}
	}

	result = 1
	for _, v := range quads {
		result *= v
	}

	return result
}

func Level2(input []string) (result int) {
	return result
}

func max(robots []robot) (height, width int) {
	for _, r := range robots {
		if r.p.x > width {
			width = r.p.x
		}
		if r.p.y > height {
			height = r.p.y
		}
	}
	return height + 1, width + 1
}

func simulate(r robot, height, width int) loc {
	for i := 0; i < seconds; i++ {
		r.p.x = getWrappedVal(r.p.x, r.v.x, width)
		r.p.y = getWrappedVal(r.p.y, r.v.y, height)
	}
	return r.p
}

func getWrappedVal(p, v, max int) int {
	n := p + v
	switch {
	case n < 0:
		return max + n
	default:
		return n % max
	}
}

func quadrant(l loc, height, width int) (quad int) {
	switch {
	case l.x < width/2 && l.y < height/2:
		return 0
	case l.x > width/2 && l.y < height/2:
		return 1
	case l.x < width/2 && l.y > height/2:
		return 2
	case l.x > width/2 && l.y > height/2:
		return 3
	default:
		return -1
	}
}

func parse(input []string) (robots []robot) {
	for _, line := range input[:len(input)-1] {
		digits := re.FindAllString(line, 4)
		robots = append(robots, robot{
			p: toLoc(digits[:2]),
			v: toLoc(digits[2:]),
		})
	}
	return robots
}

func toLoc(digits []string) loc {
	x, _ := strconv.Atoi(digits[0])
	y, _ := strconv.Atoi(digits[1])
	return loc{x: x, y: y}
}
