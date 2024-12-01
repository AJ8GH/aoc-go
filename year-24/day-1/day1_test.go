package day1

import (
	"testing"

	"github.com/aj8gh/aocgo/util"
)

var testCases = []struct {
	name    string
	subject func([]string) int64
	input   []string
	want    int64
}{
	{
		name:    "Day 1 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    11,
	},
	{
		name:    "Day 1 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    2344935,
	},
	{
		name:    "Day 1 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    31,
	},
	{
		name:    "Day 1 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    0,
	},
}

func TestDay1(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}
