package d1_test

import (
	"testing"

	"github.com/ostcar/aoc2020/d1"
)

const input = `1721
979
366
299
675
1456`

func TestD1a(t *testing.T) {
	got, err := d1.D1a(input)
	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}
	if got != "514579" {
		t.Errorf("Got %s, expected 514579", got)
	}
}

func TestD1b(t *testing.T) {
	got, err := d1.D1b(input)
	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}
	if got != "241861950" {
		t.Errorf("Got %s, expected 241861950", got)
	}
}

func BenchmarkD1b(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1.D1b(input)
	}
}
