package d14_test

import (
	"testing"

	"github.com/ostcar/aoc2020/d14"
)

func TestD14a(t *testing.T) {
	input := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
	`
	expect := "165"
	if got := d14.D14a(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}

func TestD12b(t *testing.T) {
	input := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
`
	expect := "208"
	if got := d14.D14b(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}
