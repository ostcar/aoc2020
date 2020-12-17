package d17_test

import (
	"testing"

	"github.com/ostcar/aoc2020/d17"
)

const input = `.#.
..#
###
`

func TestD17a(t *testing.T) {
	expect := "112"
	if got := d17.D17a(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}

func TestD12b(t *testing.T) {
	expect := "848"
	if got := d17.D17b(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}
