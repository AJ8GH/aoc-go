package day13

import (
	"math"
	"regexp"
	"strconv"
)

const (
	bCost       = 1
	aCost       = 3
	scaleLevel1 = 0
	scaleLevel2 = 10_000_000_000_000.0
)

var re = regexp.MustCompile(`\d+`)

type component struct {
	x, y float64
}

type game struct {
	a, b, prize component
}

func Level1(input []string) (result int) {
	return solve(input, scaleLevel1)
}

func Level2(input []string) (result int) {
	return solve(input, scaleLevel2)
}

func solve(input []string, scale float64) (result int) {
	games := parse(input, scale)
	for _, v := range games {
		result += score(v)
	}
	return result
}

func score(game game) (score int) {
	aX, aY := game.a.x, game.a.y
	bX, bY := game.b.x, game.b.y
	pX, pY := game.prize.x, game.prize.y

	b := (aY*pX - aX*pY) / (aY*bX - aX*bY)
	a := (pX - (bX * b)) / aX

	if math.Ceil(a) != a || math.Ceil(b) != b {
		return 0
	}

	return int(b)*bCost + int(a)*aCost
}

func parse(input []string, scale float64) (games []game) {
	for i := 0; i < len(input); i += 4 {
		a := toComponent(input[i])
		b := toComponent(input[i+1])
		prize := toComponent(input[i+2])
		prize.x += scale
		prize.y += scale
		games = append(games, game{a: a, b: b, prize: prize})
	}
	return games
}

func toComponent(line string) (c component) {
	digits := re.FindAllString(line, 2)
	x, _ := strconv.ParseFloat(digits[0], 64)
	y, _ := strconv.ParseFloat(digits[1], 64)
	return component{x: x, y: y}
}
