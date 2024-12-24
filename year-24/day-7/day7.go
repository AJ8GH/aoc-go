package day7

import (
	"fmt"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`\d+`)
var cache1 = map[int][][]operator{1: {{plus}, {mult}}}
var cache2 = map[int][][]operator{1: {{plus}, {mult}, {concat}}}

type operator string

func permsCache(level int) map[int][][]operator {
	switch level {
	case 1:
		return cache1
	case 2:
		return cache2
	default:
		panic("unknown level")
	}
}

const (
	plus   operator = "+"
	mult   operator = "*"
	concat operator = "||"
)

type equation struct {
	answer    int64
	inputs    []int64
	operators []operator
}

func Level1(input []string) int64 {
	return solve(input, 1)
}

func Level2(input []string) int64 {
	return solve(input, 2)
}

func solve(input []string, level int) int64 {
	equations := parse(input)

	valid := []equation{}
	for _, v := range equations {
		if isValid(&v, level) {
			valid = append(valid, v)
		}
	}
	return sumAnswers(valid)
}

func isValid(e *equation, level int) bool {
	perms := perms(len(e.inputs)-1, level)

	for _, perm := range perms {
		acc := e.inputs[0]
		for i := 1; i < len(e.inputs); i++ {
			op := perm[i-1]
			acc = apply(acc, e.inputs[i], op)
		}
		if acc == e.answer {
			e.operators = perm
			return true
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
	case concat:
		s := fmt.Sprint(a) + fmt.Sprint(b)
		n, _ := strconv.ParseInt(s, 0, 64)
		return n
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
			equations = append(equations, equation{nums[0], nums[1:], nil})
		}
	}
	return
}

func perms(n, level int) (perms [][]operator) {
	cache := permsCache(level)
	if cache[n] != nil {
		return cache[n]
	}

	perms = cache[1]
	start := 1

	for i := n; i > 1; i-- {
		if cache[i] != nil {
			perms = cache[i]
			start = i
			break
		}
	}

	for i := start; i <= n; i++ {
		if cache[i] == nil {
			cache[i] = perms
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
			if level == 2 {
				c := append(copied, concat)
				operators = append(operators, c)
			}
		}
		perms = operators
	}

	return
}
