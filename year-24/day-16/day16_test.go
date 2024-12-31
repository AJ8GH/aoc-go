package day16

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
		name:    "Day 16 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    7036,
	},
	{
		name:    "Day 16 Level 1 Example 2",
		subject: Level1,
		input:   util.ReadExampleN(2),
		want:    11048,
	},
	{
		name:    "Day 16 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    0,
	},
	{
		name:    "Day 16 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    0,
	},
	{
		name:    "Day 16 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    0,
	},
}

func TestDay16(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
