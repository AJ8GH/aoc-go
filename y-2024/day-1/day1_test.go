package day1

import (
	"testing"

	"github.com/aj8gh/aocgo/util"
)

var testCases = []struct {
	name  string
	input []string
}{
	{
		name:  "Day 1 Level 1 - Example",
		input: util.ReadExample(),
	},
	{
		name:  "Day 1 Level 1",
		input: util.ReadInput(),
	},
	{
		name:  "Day 1 Level 2 - Example",
		input: util.ReadExample(),
	},
	{
		name:  "Day 1 Level 2",
		input: util.ReadInput(),
	},
}

func TestDay1(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Level1(tt.input)
			want := 0
			util.AssertInt(t, got, want)
		})
	}
}
