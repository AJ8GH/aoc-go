package day12

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
		name:    "Day 12 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    140,
	},
	{
		name:    "Day 12 Level 1 Example 2",
		subject: Level1,
		input:   util.ReadExampleN(2),
		want:    772,
	},
	{
		name:    "Day 12 Level 1 Example 3",
		subject: Level1,
		input:   util.ReadExampleN(3),
		want:    1930,
	},
	{
		name:    "Day 12 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    1533644,
	},
	{
		name:    "Day 12 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    80,
	},
	{
		name:    "Day 12 Level 2 Example 2",
		subject: Level2,
		input:   util.ReadExampleN(2),
		want:    436,
	},
	{
		name:    "Day 12 Level 2 Example 4",
		subject: Level2,
		input:   util.ReadExampleN(4),
		want:    236,
	},
	{
		name:    "Day 12 Level 2 Example 5",
		subject: Level2,
		input:   util.ReadExampleN(5),
		want:    368,
	},
	{
		name:    "Day 12 Level 2 Example 3",
		subject: Level2,
		input:   util.ReadExampleN(3),
		want:    1206,
	},
	{
		name:    "Day 12 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    936718,
	},
}

func TestDay12(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
