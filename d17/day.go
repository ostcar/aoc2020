package d17

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// D17a solves first part of day 17.
func D17a(input string) string {
	w := newWorld(input, 3)
	for i := 0; i < 6; i++ {
		w = w.step()
	}
	return strconv.Itoa(w.count())
}

// D17b solves first part of day 17.
func D17b(input string) string {
	w := newWorld(input, 4)
	for i := 0; i < 6; i++ {
		w = w.step()
	}
	return strconv.Itoa(w.count())
}

type world struct {
	cubes map[point]bool
	max   point
	min   point

	dim int
}

func newWorld(input string, dim int) *world {
	w := &world{
		cubes: make(map[point]bool),
		max:   point{},
		dim:   dim,
	}
	for y, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		w.max[1] = max(w.max.y(), y)
		for x, c := range line {
			w.max[0] = max(w.max.x(), x)
			if c == '#' {
				w.cubes[point{x, y}] = true
			}
		}
	}
	return w
}

func (w *world) step() *world {
	nw := &world{
		cubes: make(map[point]bool),
		max:   w.max.move(1, w.dim),
		min:   w.min.move(-1, w.dim),
		dim:   w.dim,
	}

	for pw := nw.min.w(); pw <= nw.max.w(); pw++ {
		for z := nw.min.z(); z <= nw.max.z(); z++ {
			for y := nw.min.y(); y <= nw.max.y(); y++ {
				for x := nw.min.x(); x <= nw.max.x(); x++ {
					p := point{x, y, z, pw}
					n := w.neighbor(p)

					// Was active.
					if w.cubes[p] {
						if n == 2 || n == 3 {
							nw.cubes[p] = true
						}
						continue
					}

					// Was inactive.
					if n == 3 {
						nw.cubes[p] = true
					}
				}
			}
		}
	}
	return nw
}

func (w *world) count() int {
	var count int
	for pw := w.min.w(); pw <= w.max.w(); pw++ {
		for z := w.min.z(); z <= w.max.z(); z++ {
			for y := w.min.y(); y <= w.max.y(); y++ {
				for x := w.min.x(); x <= w.max.x(); x++ {
					if w.cubes[point{x, y, z, pw}] {
						count++
					}
				}
			}
		}
	}
	return count
}

func (w *world) neighbor(p point) int {
	var count int

	for pw := p.w() - 1; pw <= p.w()+1; pw++ {
		for z := p.z() - 1; z <= p.z()+1; z++ {
			for y := p.y() - 1; y <= p.y()+1; y++ {
				for x := p.x() - 1; x <= p.x()+1; x++ {
					if z == p.z() && y == p.y() && x == p.x() && pw == p.w() {
						continue
					}
					if w.cubes[point{x, y, z, pw}] {
						count++
					}
				}
			}
		}
	}
	return count
}

func (w *world) String() string {
	var buf bytes.Buffer
	for z := w.min.z(); z <= w.max.z(); z++ {
		buf.WriteString(fmt.Sprintf("z=%d\n", z))
		for y := w.min.y(); y <= w.max.y(); y++ {
			for x := w.min.x(); x <= w.max.x(); x++ {
				var b byte = '.'
				if w.cubes[point{x, y, z}] {
					b = '#'
				}
				buf.WriteByte(b)
			}
			buf.WriteByte('\n')
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

type point [4]int

func (p point) move(n int, dim int) point {
	var np point
	for i := 0; i < dim; i++ {
		np[i] = p[i] + n
	}
	return np
}

func (p point) x() int {
	return p[0]
}

func (p point) y() int {
	return p[1]
}

func (p point) z() int {
	return p[2]
}

func (p point) w() int {
	return p[3]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
