package d5

import "testing"

func TestDecode(t *testing.T) {
	for _, tt := range []struct {
		code   string
		expect int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	} {
		t.Run(tt.code, func(t *testing.T) {
			got := decode(tt.code)
			if got != tt.expect {
				t.Errorf("Got %d, expected %d", got, tt.expect)
			}
		})
	}
}
