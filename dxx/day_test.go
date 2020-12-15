package dxx_test

import (
	"testing"

	"github.com/ostcar/aoc2020/dxx"
)

const input = `
`

func TestDxxa(t *testing.T) {
	expect := ""
	if got := dxx.Dxxa(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}

func TestD12b(t *testing.T) {
	expect := ""
	if got := dxx.Dxxb(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}
