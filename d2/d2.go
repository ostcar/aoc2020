package d2

import (
	"fmt"
	"strconv"
	"strings"
)

type entry struct {
	min, max int
	letter   string
	password string
}

func parseInput(input string) ([]entry, error) {
	var entries []entry
	for _, line := range strings.Split(input, "\n") {
		var e entry
		if _, err := fmt.Sscanf(line, "%d-%d %s %s", &e.min, &e.max, &e.letter, &e.password); err != nil {
			return nil, fmt.Errorf("scanning line `%s`: %w", line, err)
		}

		e.letter = e.letter[:1]

		entries = append(entries, e)
	}

	return entries, nil
}

// D2a solves day2 a.
func D2a(input string) string {
	entries, err := parseInput(input)
	if err != nil {
		return fmt.Sprintf("Can not parse input: %v", err)
	}

	var valid int
	for _, e := range entries {
		count := strings.Count(e.password, e.letter)
		if count > e.max || count < e.min {
			continue
		}
		valid++
	}
	return strconv.Itoa(valid)
}

// D2b solves day2 b.
func D2b(input string) string {
	entries, err := parseInput(input)
	if err != nil {
		return fmt.Sprintf("Can not parse input: %v", err)
	}

	var valid int
	for _, e := range entries {
		l1 := e.password[e.min-1 : e.min]
		l2 := e.password[e.max-1 : e.max]
		if l1 == l2 {
			continue
		}
		if l1 != e.letter && l2 != e.letter {
			continue
		}
		valid++
	}
	return strconv.Itoa(valid)
}
