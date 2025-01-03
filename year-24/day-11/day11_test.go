package day11

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
		name:    "Day 11 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    55312,
	},
	{
		name:    "Day 11 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    202019,
	},
	{
		name:    "Day 11 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    65601038650482,
	},
	{
		name:    "Day 11 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    239321955280205,
	},
}

func TestDay11(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
