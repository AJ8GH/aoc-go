package day13

import (
	"regexp"
	"strconv"
)

const (
	bCost    = 1
	aCost    = 3
	maxCount = 100
)

var re = regexp.MustCompile(`\d+`)

type component struct {
	x, y int
}

type game struct {
	a, b, prize component
}

func Level1(input []string) (result int) {
	games := parse(input)
	for _, v := range games {
		result += score(v)
	}
	return result
}

func Level2(input []string) (result int) {
	return result
}

/*
check 100 b
if < <

	check 100 a

while b > 0

	minus 1 b
	check 100 a

return
*/
func score(game game) (score int) {
	bCount := 0
	x := 0
	y := 0

	for bCount <= maxCount {
		x += game.b.x
		y += game.b.y
		bCount++
		if x == game.prize.x && y == game.prize.y {
			return bCount * bCost
		}
		if x > game.prize.x || y > game.prize.y {
			break
		}
	}

	aCount := 0
	if x <= game.prize.x && y <= game.prize.y {
		for aCount <= maxCount {
			x += game.a.x
			y += game.a.y
			aCount++
			if x == game.prize.x && y == game.prize.y {
				return bCount*bCost + aCount*aCost
			}
			if x > game.prize.x || y > game.prize.y {
				break
			}
		}
	}

	x -= aCount * game.a.x
	y -= aCount * game.a.y
	aCount = 0
	for aCount <= maxCount {
		if bCount > 0 {
			bCount--
			x -= game.b.x
			y -= game.b.y
		}
		tempACount := aCount
		tempX := x
		tempY := y
		for tempACount <= maxCount {
			tempX += game.a.x
			tempY += game.a.y
			tempACount++
			if tempX == game.prize.x && tempY == game.prize.y {
				return bCount*bCost + tempACount*aCost
			}
			if tempX > game.prize.x || y > game.prize.y {
				break
			}
		}
		if bCount == 0 {
			break
		}
	}

	return 0
}

func parse(input []string) (games []game) {
	for i := 0; i < len(input); i += 4 {
		a := toComponent(input[i])
		b := toComponent(input[i+1])
		prize := toComponent(input[i+2])
		games = append(games, game{a: a, b: b, prize: prize})
	}
	return games
}

func toComponent(line string) (c component) {
	digits := re.FindAllString(line, 2)
	x, _ := strconv.Atoi(digits[0])
	y, _ := strconv.Atoi(digits[1])
	return component{x: x, y: y}
}
