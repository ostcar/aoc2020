package d1

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, error) {
	var nums []int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid input `%s`: %w", line, err)
		}
		nums = append(nums, i)
	}
	return nums, nil
}

// D1a solves day1 a
func D1a(input string) (string, error) {
	nums, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("can not load input: %w", err)
	}

	for i, n1 := range nums {
		for j, n2 := range nums {
			if i == j {
				continue
			}
			if n1+n2 != 2020 {
				continue
			}
			return strconv.Itoa(n1 * n2), nil
		}
	}
	return "", fmt.Errorf("no correct values :(")
}

// D1b solves day1 b
func D1b(input string) (string, error) {
	nums, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("can not load input: %w", err)
	}

	for i, n1 := range nums {
		for j, n2 := range nums {
			for k, n3 := range nums {
				if i == j || i == k || j == k {
					continue
				}
				if n1+n2+n3 != 2020 {
					continue
				}
				return strconv.Itoa(n1 * n2 * n3), nil
			}

		}
	}
	return "", fmt.Errorf("no correct values :(")
}
