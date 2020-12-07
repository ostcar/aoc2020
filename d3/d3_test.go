package d3_test

import (
	"strings"
	"testing"

	"github.com/ostcar/aoc2020/d3"
)

const input = `
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestD3a(t *testing.T) {
	got := d3.D3a(cleared(input))
	if got != "7" {
		t.Errorf("Got %s, expected 7", got)
	}
}

func TestD3b(t *testing.T) {
	got := d3.D3b(cleared(input))
	if got != "336" {
		t.Errorf("Got %s, expected 336", got)
	}
}

func cleared(input string) string {
	var out string
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		out += strings.TrimSpace(line) + "\n"
	}
	return out
}
