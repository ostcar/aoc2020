package d11

import (
	"bytes"
	"strconv"
	"strings"
)

// D11a solves first part of day 11.
func D11a(input string) string {
	g := newGrid(input)
	g.neighborLimit = 4
	g.calcNeighborA()
	for changed := true; changed; g, changed = g.step() {
	}

	return strconv.Itoa(g.count())
}

// D11b solves second part of day 11.
func D11b(input string) string {
	g := newGrid(input)
	g.neighborLimit = 5
	g.calcNeighborB()

	for changed := true; changed; g, changed = g.step() {
	}

	return strconv.Itoa(g.count())
}

type grid struct {
	seats  map[pos]int
	width  int
	height int

	neighborLimit int
	neighbors     map[pos][]pos
}

func newGrid(input string) *grid {
	lines := strings.Split(input[:len(input)-1], "\n")
	g := &grid{
		seats:     make(map[pos]int, len(lines[0])*len(lines)),
		width:     len(lines[0]),
		height:    len(lines),
		neighbors: make(map[pos][]pos, len(lines[0])*len(lines)),
	}

	for y, line := range lines {
		for x, c := range line {
			if c == 'L' {
				g.seats[pos{x, y}] = posEmtpy
			}
		}
	}

	return g
}

func (g *grid) calcNeighborA() {
	for p := range g.seats {

		for x := -1; x <= +1; x++ {
			for y := -1; y <= +1; y++ {
				if x == 0 && y == 0 {
					continue
				}

				np := pos{p.x + x, p.y + y}
				if g.seats[np] != posFloor {
					g.neighbors[p] = append(g.neighbors[p], np)
				}
			}
		}
	}
}

func (g *grid) calcNeighborB() {
	for p := range g.seats {
		for x := -1; x <= +1; x++ {
			for y := -1; y <= +1; y++ {
				if x == 0 && y == 0 {
					continue
				}

				for i := 1; ; i++ {
					np := pos{x*i + p.x, y*i + p.y}
					if np.x < 0 || np.x >= g.width || np.y < 0 || np.y >= g.height {
						break
					}
					if g.seats[np] != posFloor {
						g.neighbors[p] = append(g.neighbors[p], np)
						break
					}
				}
			}
		}
	}
}

func (g *grid) neighborCount(p pos) int {
	var count int
	for _, n := range g.neighbors[p] {
		if g.seats[n] == posOccupied {
			count++
		}
	}
	return count
}

func (g *grid) step() (*grid, bool) {
	gn := &grid{
		width:         g.width,
		height:        g.height,
		seats:         make(map[pos]int, len(g.seats)),
		neighborLimit: g.neighborLimit,
		neighbors:     g.neighbors,
	}

	var changed bool
	for k, v := range g.seats {
		gn.seats[k] = v
		if v == posEmtpy && g.neighborCount(k) == 0 {
			gn.seats[k] = posOccupied
			changed = true
		}

		if v == posOccupied && g.neighborCount(k) >= g.neighborLimit {
			gn.seats[k] = posEmtpy
			changed = true
		}
	}
	return gn, changed
}

func (g *grid) count() int {
	var c int
	for _, v := range g.seats {
		if v == posOccupied {
			c++
		}
	}
	return c
}

func (g *grid) String() string {
	buf := bytes.NewBuffer(make([]byte, 0, (g.width+1)*g.height))
	for y := 0; y < g.height; y++ {
		lb := make([]byte, g.width)
		for x := 0; x < g.width; x++ {
			switch g.seats[pos{x, y}] {
			case posFloor:
				lb[x] = '.'
			case posEmtpy:
				lb[x] = 'L'
			case posOccupied:
				lb[x] = '#'
			default:
				lb[x] = '?'
			}
		}
		buf.Write(lb)
		buf.WriteByte('\n')
	}
	return buf.String()
}

type pos struct {
	x, y int
}

const (
	posFloor = iota
	posEmtpy
	posOccupied
)
