package day4

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
		name:    "Day 4 Level 1 Example",
		subject: Level1,
		input:   util.ReadExample(),
		want:    18,
	},
	{
		name:    "Day 4 Level 1",
		subject: Level1,
		input:   util.ReadInput(),
		want:    2685,
	},
	{
		name:    "Day 4 Level 2 Example",
		subject: Level2,
		input:   util.ReadExample(),
		want:    9,
	},
	{
		name:    "Day 4 Level 2",
		subject: Level2,
		input:   util.ReadInput(),
		want:    2048,
	},
}

func TestDay4(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.subject(tt.input)
			util.Assert(t, got, tt.want)
		})
	}
}

/*

  | a  b  c  d  e  f  g  h  i  j
--+------------------------------
0 | .  .  .  .  X  X  M  A  S  .
1 | .  S  A  M  X  M  S  .  .  .
2 | .  .  .  S  .  .  A  .  .  .
3 | .  .  A  .  A  .  M  S  .  X
4 | X  M  A  S  A  M  X  .  M  M
5 | X  .  .  .  .  .  X  A  .  A
6 | S  .  S  .  S  .  S  .  S  S
7 | .  A  .  A  .  A  .  A  .  A
8 | .  .  M  .  M  .  M  .  M  M
9 | .  X  .  X  .  X  M  A  S  X


  | a  b  c  d  e  f  g  h  i  j
--+------------------------------
0 | M  M  M  S  X  X  M  A  S  M
1 | M  S  A  M  X  M  S  M  S  A
2 | A  M  X  S  X  M  A  A  M  M
3 | M  S  A  M  A  S  M  S  M  X
4 | X  M  A  S  A  M  X  A  M  M
5 | X  X  A  M  M  X  X  A  M  A
6 | S  M  S  M  S  A  S  X  S  S
7 | S  A  X  A  M  A  S  A  A  A
8 | M  A  M  M  M  X  M  M  M  M
9 | M  X  M  X  A  X  M  A  S  X




*/
