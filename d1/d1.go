package d1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseSlice(input string) ([]int, error) {
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
	lines := strings.Split(input, "\n")
	m := make(map[int]bool, len(lines))
	nums := make([]int, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			return "", fmt.Errorf("invalid input `%s`: %w", line, err)
		}

		if i < 1010 {
			nums = append(nums, i)
		} else {
			m[i] = true
		}
	}

	for _, num := range nums {
		if m[2020-num] {
			return strconv.Itoa(num * (2020 - num)), nil
		}
	}

	return "", fmt.Errorf("no correct values :(")
}

// D1b solves day1 b
func D1b(input string) (string, error) {
	nums, err := parseSlice(input)
	if err != nil {
		return "", fmt.Errorf("can not load input: %w", err)
	}

	sort.Ints(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[i]+nums[j] > 2020 {
				continue
			}

			//
			for k := 0; k < j; k++ {
				if nums[i]+nums[j]+nums[k] > 2020 {
					break
				}
				if nums[i]+nums[j]+nums[k] != 2020 {
					continue
				}
				return strconv.Itoa(nums[i] * nums[j] * nums[k]), nil
			}

		}
	}
	return "", fmt.Errorf("no correct values :(")
}
