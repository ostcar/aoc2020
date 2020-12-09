package d9

import (
	"strconv"
	"strings"
)

func parse(input string) []int64 {
	lines := strings.Split(input[:len(input)-1], "\n")
	numbers := make([]int64, len(lines))
	for i, line := range lines {
		n, _ := strconv.ParseInt(line, 10, 64)
		numbers[i] = n
	}
	return numbers
}

// D9a solves the first part of day 9.
func D9a(input string) string {
	numbers := parse(input)
	return strconv.FormatInt(numbers[firstAnomalyIndex(numbers)], 10)
}

// D9b solves the secod part of day 9.
func D9b(input string) string {
	numbers := parse(input)
	anomaly := firstAnomalyIndex(numbers)

numberLoop:
	for i := 0; i < anomaly; i++ {
		sum := numbers[i]
		for j := i + 1; j < anomaly; j++ {
			sum += numbers[j]
			if sum == numbers[anomaly] {
				return strconv.FormatInt(minMax(numbers, i, j), 10)
			}
			if sum > numbers[anomaly] {
				continue numberLoop
			}
		}
	}

	return "No number found"
}

func firstAnomalyIndex(numbers []int64) int {
numberLoop:
	for i := 25; i < len(numbers); i++ {
		for a := i - 25; a < i; a++ {
			for b := a + 1; b < i; b++ {
				if numbers[a]+numbers[b] == numbers[i] {
					continue numberLoop
				}
			}
		}
		return i
	}
	return -1
}

func minMax(numbers []int64, a, b int) int64 {
	var min int64 = 9223372036854775807
	var max int64
	for i := a; i <= b; i++ {
		if numbers[i] < min {
			min = numbers[i]
		}

		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return min + max
}
