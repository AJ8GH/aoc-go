package day3

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
		name:    "Day 3 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    161,
	},
	{
		name:    "Day 3 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    170807108,
	},
	{
		name:    "Day 3 Level 2 Example",
		subject: Level2,
		input:   util.ReadExampleN(2),
		want:    48,
	},
	{
		name:    "Day 3 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    74838033,
	},
}

func TestDay3(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
