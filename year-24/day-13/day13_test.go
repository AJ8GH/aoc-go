package day13

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
		name:    "Day 13 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    480,
	},
	{
		name:    "Day 13 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    37901,
	},
	{
		name:    "Day 13 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    0,
	},
	{
		name:    "Day 13 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    0,
	},
}

func TestDay13(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
