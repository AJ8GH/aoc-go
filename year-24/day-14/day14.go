package day14

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

const (
	level1Seconds    = 100
	level2Seconds    = 1_000_000_000
	lengthToConsider = 10
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

	for i := 0; i < level1Seconds; i++ {
		robots = simulate(robots, height, width)
	}

	for _, r := range robots {
		quad := quadrant(r.p, height, width)
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
	robots := parse(input)
	height, width := max(robots)

	for i := 0; i < level2Seconds; i++ {
		robots = simulate(robots, height, width)
		if continuousRowExists(robots, lengthToConsider) {
			show(robots, height, width, i)
			return i + 1
		}
	}

	return result
}

func max(robots []*robot) (height, width int) {
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

func simulate(robots []*robot, height, width int) []*robot {
	for _, r := range robots {
		r.p.x = getWrappedVal(r.p.x, r.v.x, width)
		r.p.y = getWrappedVal(r.p.y, r.v.y, height)
	}
	return robots
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

func parse(input []string) (robots []*robot) {
	for _, line := range input[:len(input)-1] {
		digits := re.FindAllString(line, 4)
		robots = append(robots, &robot{
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

func continuousRowExists(robots []*robot, lengthToConsider int) bool {
	points := map[int][]int{}
	for _, r := range robots {
		if points[r.p.y] == nil {
			points[r.p.y] = []int{}
		}
		points[r.p.y] = append(points[r.p.y], r.p.x)
	}

	for _, v := range points {
		if len(v) < lengthToConsider {
			continue
		}

		slices.Sort(v)
		currentLength := 0
		for i, point := range v {
			if i == 0 {
				continue
			}

			if point-v[i-1] == 1 {
				currentLength++
				if currentLength == lengthToConsider {
					return true
				}
			} else {
				currentLength = 0
			}
		}
	}
	return false
}

func show(robots []*robot, height, width, second int) {
	grid := make([][]string, height)
	for i := range grid {
		row := make([]string, width)
		for j := range row {
			row[j] = "."
		}
		grid[i] = row
	}
	for _, r := range robots {
		grid[r.p.y][r.p.x] = "O"
	}
	out := fmt.Sprintf("\nSecond %d:\n", second)
	for _, v := range grid {
		out += fmt.Sprintf("%v\n", v)
	}

	println(out)

	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(out); err != nil {
		panic(err)
	}
}
