package day17

import (
	"testing"

	"github.com/aj8gh/aocgo/util"
)

var testCases = []struct {
	name    string
	subject func([]string) string
	input   []string
	want    string
}{
	{
		name:    "Day 17 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    "4,6,3,5,6,3,5,2,1,0",
	},
	{
		name:    "Day 17 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    "",
	},
	{
		name:    "Day 17 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    "",
	},
	{
		name:    "Day 17 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    "",
	},
}

func TestDay17(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
