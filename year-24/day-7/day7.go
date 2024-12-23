package day7

import (
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`\d+`)

type operator string

var permsCache = map[int][][]operator{
	1: {{plus}, {mult}},
}

const (
	plus operator = "+"
	mult operator = "*"
)

type equation struct {
	answer int64
	inputs []int64
}

func Level1(input []string) int64 {
	equations := parse(input)

	valid := []equation{}
	for _, v := range equations {
		if isValid(v) {
			valid = append(valid, v)
		}
	}
	return sumAnswers(valid)
}

func Level2(input []string) int64 {
	return 0
}

func isValid(e equation) bool {
	perms := perms(len(e.inputs) - 1)

	for _, perm := range perms {
		acc := e.inputs[0]
		for i := 1; i < len(e.inputs); i++ {
			op := perm[i-1]
			acc = apply(acc, e.inputs[i], op)
			if acc == e.answer {
				return true
			}
		}
	}

	return false
}

func apply(a, b int64, op operator) int64 {
	switch op {
	case plus:
		return a + b
	case mult:
		return a * b
	default:
		panic("Unknown operator")
	}
}

func sumAnswers(equations []equation) (sum int64) {
	for _, v := range equations {
		sum += v.answer
	}
	return
}

func parse(input []string) (equations []equation) {
	for _, v := range input {
		nums := []int64{}
		vals := re.FindAllString(v, -1)
		for _, s := range vals {
			n, _ := strconv.Atoi(s)
			nums = append(nums, int64(n))
		}
		if len(nums) != 0 {
			equations = append(equations, equation{nums[0], nums[1:]})
		}
	}
	return
}

func perms(n int) (perms [][]operator) {
	if permsCache[n] != nil {
		return permsCache[n]
	}

	perms = permsCache[1]
	start := 1

	for i := n; i > 1; i-- {
		if permsCache[i] != nil {
			perms = permsCache[i]
			start = i
			break
		}
	}

	for i := start; i <= n; i++ {
		if permsCache[i] == nil {
			permsCache[i] = perms
		}

		if i == n {
			break
		}

		operators := [][]operator{}
		for _, v := range perms {
			copied := make([]operator, len(v))
			copy(copied, v)
			a := append(v, plus)
			b := append(copied, mult)
			operators = append(operators, a)
			operators = append(operators, b)
		}
		perms = operators
	}

	return
}

/*

	2

	[+]
	[*]

	---

	[++]
	[+*]
	[*+]
	[**]

	---

	[++*]
	[+++]
	[+**]
	[+*+]
	[*+*]
	[*++]
	[***]
	[**+]

	___

	[++*+]
	[++**]
	[++++]
	[+++*]
	[+**+]
	[+***]
	[+*++]
	[+*+*]
	[*+*+]
	[*+**]
	[*+++]
	[*++*]
	[***+]
	[****]
	[**++]
	[**+*]




*/
