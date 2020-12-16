package d15_test

import (
	"testing"

	"github.com/ostcar/aoc2020/d15"
)

func TestD15a(t *testing.T) {
	for _, tt := range []struct {
		input  string
		expect string
	}{
		{"0,3,6", "436"},
		{"1,3,2", "1"},
		{"2,1,3", "10"},
		{"1,2,3", "27"},
		{"2,3,1", "78"},
		{"3,2,1", "438"},
		{"3,1,2", "1836"},
	} {
		t.Run(tt.input, func(t *testing.T) {
			if got := d15.D15a(tt.input); got != tt.expect {
				t.Errorf("Got %s, expected %s", got, tt.expect)
			}
		})
	}
}

func TestD12b(t *testing.T) {
	for _, tt := range []struct {
		input  string
		expect string
	}{
		{"0,3,6", "175594"},
		// {"1,3,2", "2578"},
		// {"2,1,3", "3544142"},
		// {"1,2,3", "261214"},
		// {"2,3,1", "6895259"},
		// {"3,2,1", "18"},
		// {"3,1,2", "362"},
	} {
		t.Run(tt.input, func(t *testing.T) {
			if got := d15.D15b(tt.input); got != tt.expect {
				t.Errorf("Got %s, expected %s", got, tt.expect)
			}
		})
	}
}
