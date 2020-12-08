package d8_test

import (
	"testing"

	"github.com/ostcar/aoc2020/d8"
)

const input = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

func TestD8a(t *testing.T) {
	if got := d8.D8a(input); got != "5" {
		t.Errorf("Got %s, expected 5", got)

	}
}

func TestD8b(t *testing.T) {
	if got := d8.D8b(input); got != "8" {
		t.Errorf("Got %s, expected 8", got)

	}
}
