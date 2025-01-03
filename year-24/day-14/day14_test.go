package day14

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
		name:    "Day 14 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    12,
	},
	{
		name:    "Day 14 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    215476074,
	},
	{
		name:    "Day 14 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    6285,
	},
}

func TestDay14(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
