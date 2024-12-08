package day2

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
		name:    "Day 2 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    2,
	},
	{
		name:    "Day 2 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    236,
	},
	{
		name:    "Day 2 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    4,
	},
	{
		name:    "Day 2 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    308,
	},
}

func TestDay2(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
