package day5

import (
	"testing"

	"github.com/aj8gh/aocgo/util"
)

var testCases = []struct {
	name    string
	subject func([]string) int
	input   []string
	want    int
}{
	{
		name:    "Day 5 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    143,
	},
	{
		name:    "Day 5 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    4790,
	},
	{
		name:    "Day 5 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    123,
	},
	{
		name:    "Day 5 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    6319,
	},
}

func TestDay5(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}

/*

47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13


75,47,61,53,29

[75, 47, 61] 53

{
	75 -> [47, 61, 53, 29]
	47 -> [53, 61, 29]
	53 -> [29, 13]
}

97,61,53,29,13

75,29,13

75,97,47,61,53



61,13,29

97,13,75,29,47



*/
