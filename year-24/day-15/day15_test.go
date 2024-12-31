package day15

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
		name:    "Day 15 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    2028,
	},
	{
		name:    "Day 15 Level 1 Example 2",
		subject: Level1,
		input:   util.ReadExampleN(2),
		want:    10092,
	},
	{
		name:    "Day 15 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    1526018,
	},
	{
		name:    "Day 15 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    105,
	},
	{
		name:    "Day 15 Level 2 Example 2",
		subject: Level2,
		input:   util.ReadExampleN(2),
		want:    9021,
	},
	{
		name:    "Day 15 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    0,
	},
}

func TestDay15(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
