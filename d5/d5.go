package d5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// D5a solves the first part of day 5.
func D5a(input string) string {
	codes := parse(input)
	return strconv.Itoa(decode(codes[len(codes)-1]))
}

// D5b solves the second part of day 5.
func D5b(input string) string {
	codes := parse(input)
	exist := make(map[int]bool, len(codes))
	for _, code := range codes {
		exist[decode(code)] = true
	}
	for i := 8; i < 1016; i++ {
		if !exist[i] && exist[i-1] && exist[i+1] {
			return strconv.Itoa(i)
		}
	}
	return fmt.Sprintf("No result :(")
}

func parse(input string) []string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	sort.Slice(lines, func(i, j int) bool {
		for idx := 0; idx < 10; idx++ {
			if lines[i][idx] == lines[j][idx] {
				continue
			}
			l := lines[j][idx]
			return l == 'B' || l == 'R'
		}
		return false
	})
	return lines
}

func decode(code string) int {
	var o int
	for _, l := range code {
		o <<= 1
		if l == 'B' || l == 'R' {
			o |= 1
		}
	}
	return o
}
