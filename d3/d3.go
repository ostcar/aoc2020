package d3

import (
	"strconv"
	"strings"
)

// D3a solves day3 a
func D3a(input string) (string, error) {
	p := newPlan(input)
	trees := p.walk(3, 1)
	return strconv.Itoa(trees), nil
}

// D3b solves day3 b
func D3b(input string) (string, error) {
	p := newPlan(input)
	trees := 1
	for _, way := range [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	} {
		trees = trees * p.walk(way[0], way[1])
	}
	return strconv.Itoa(trees), nil
}

type plan struct {
	trees          map[[2]int]bool
	height, weight int
}

func newPlan(input string) *plan {
	p := new(plan)
	p.trees = make(map[[2]int]bool)

	for y, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		for x, column := range line {
			if column == '#' {
				p.set(x, y)
			}
		}
		p.height = y + 1
		p.weight = len(line)
	}
	return p
}

func (p *plan) set(x, y int) {
	p.trees[[2]int{x, y}] = true
}

func (p *plan) get(x, y int) bool {
	x = x % p.weight
	return p.trees[[2]int{x, y}]
}

func (p *plan) walk(right, down int) int {
	var trees int
	var x int
	var y int
	for y < p.height {
		if p.get(x, y) {
			trees++
		}
		x = x + right
		y = y + down
	}
	return trees
}
