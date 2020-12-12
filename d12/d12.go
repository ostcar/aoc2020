package d12

import (
	"fmt"
	"strconv"
	"strings"
)

// D12a solves first part of day 12.
func D12a(input string) string {
	s := new(ship)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		value, _ := strconv.Atoi(line[1:])
		s.doA(line[0], value)
	}
	return strconv.Itoa(abs(s.x) + abs(s.y))
}

// D12b solves first part of day 12.
func D12b(input string) string {
	s := new(ship)
	s.vx = 10
	s.vy = -1
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		value, _ := strconv.Atoi(line[1:])
		s.doB(line[0], value)
	}
	return strconv.Itoa(abs(s.x) + abs(s.y))
}

type ship struct {
	direction int
	x, y      int
	vx, vy    int
}

func (s *ship) doA(cmd byte, value int) {
	switch cmd {
	case 'N':
		s.y -= value
	case 'S':
		s.y += value
	case 'E':
		s.x += value
	case 'W':
		s.x -= value
	case 'L':
		s.direction -= value
	case 'R':
		s.direction += value
	case 'F':
		for s.direction < 0 {
			s.direction += 360
		}
		s.direction %= 360
		switch s.direction {
		case 0:
			s.doA('E', value)
		case 90:
			s.doA('S', value)
		case 180:
			s.doA('W', value)
		case 270:
			s.doA('N', value)
		default:
			panic(fmt.Sprintf("Invalid direction %d", s.direction))
		}
	}
}

func (s *ship) doB(cmd byte, value int) {
	switch cmd {
	case 'N':
		s.vy -= value
	case 'S':
		s.vy += value
	case 'E':
		s.vx += value
	case 'W':
		s.vx -= value
	case 'L':
		s.doB('R', 360-value)

	case 'R':
		switch value {
		case 90:
			s.vy, s.vx = s.vx, -s.vy
		case 180:
			s.vy, s.vx = -s.vy, -s.vx
		case 270:
			s.vy, s.vx = -s.vx, s.vy
		}

	case 'F':
		s.x += s.vx * value
		s.y += s.vy * value
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
