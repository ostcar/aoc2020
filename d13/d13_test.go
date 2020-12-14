package d13_test

import (
	"testing"

	"github.com/ostcar/aoc2020/d13"
)

func TestD13a(t *testing.T) {
	const input = `939
7,13,x,x,59,x,31,19
`
	expect := "295"
	if got := d13.D13a(input); got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}

func TestD12b(t *testing.T) {
	for _, tt := range []struct {
		input  string
		expect string
	}{
		{"7,13,x,x,59,x,31,19", "1068781"},
		{"17,x,13,19", "3417"},
		{"67,7,59,61", "754018"},
		{"67,x,7,59,61", "779210"},
		{"67,7,x,59,61", "1261476"},
		{"1789,37,47,1889", "1202161486"},
	} {
		t.Run(tt.input, func(t *testing.T) {
			if got := d13.D13b("\n" + tt.input); got != tt.expect {
				t.Errorf("Got %s, expected %s", got, tt.expect)
			}
		})
	}
}
