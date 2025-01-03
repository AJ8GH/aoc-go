package day8

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
		name:    "Day 8 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    14,
	},
	{
		name:    "Day 8 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    285,
	},
	{
		name:    "Day 8 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    34,
	},
	{
		name:    "Day 8 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    944,
	},
}

func TestDay8(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
