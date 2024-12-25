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
		want:    0,
	},
	{
		name:    "Day 8 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    0,
	},
	{
		name:    "Day 8 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    0,
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

/*

8,8
9,9

dist = -1, -1

8,8 + -1,-1
7,7
-----------
9,9 - -1,-1
10,10




1,8
2,5

dist = -1,3

1,8  + -1,3
0,11 + -1,3
-----------
2,5 - -1,3
3,2


*/
