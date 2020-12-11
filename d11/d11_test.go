package d11_test

import (
	"testing"

	"github.com/ostcar/aoc2020/d11"
)

const input = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`

func TestD11a(t *testing.T) {
	if got := d11.D11a(input); got != "37" {
		t.Errorf("Got %s, expected 37", got)
	}
}

func TestD11b(t *testing.T) {
	if got := d11.D11b(input); got != "26" {
		t.Errorf("Got %s, expected 26", got)
	}
}
