package day24

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
		name:    "Day 24 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    4,
	},
	{
		name:    "Day 24 Level 1 Example 2",
		subject: Level1,
		input:   util.ReadExampleN(2),
		want:    2024,
	},
	{
		name:    "Day 24 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    0,
	},
	{
		name:    "Day 24 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    0,
	},
	{
		name:    "Day 24 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    0,
	},
}

func TestDay24(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
