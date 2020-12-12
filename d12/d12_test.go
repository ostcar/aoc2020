package d12_test

import (
	"testing"

	"github.com/ostcar/aoc2020/d12"
)

const input = `F10
N3
F7
R90
F11
`

func TestD12a(t *testing.T) {
	expect := "25"
	if got := d12.D12a(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}

func TestD12b(t *testing.T) {
	expect := "286"
	if got := d12.D12b(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}
