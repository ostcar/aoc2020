package d22_test

import (
	"testing"

	"github.com/ostcar/aoc2020/d22"
)

const input = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`

func TestD22a(t *testing.T) {
	expect := "306"
	if got := d22.D22a(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}

func TestD12b(t *testing.T) {
	expect := "291"
	if got := d22.D22b(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}

func TestD12bLoop(t *testing.T) {
	input := `Player 1:
43
19

Player 2:
2
29
14
`
	expect := "105"
	if got := d22.D22b(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}
