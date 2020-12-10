package d10

import (
	"sort"
	"strconv"
	"strings"
)

// D10a solces first part of day 10.
func D10a(input string) string {
	lines := strings.Split(input[:len(input)-1], "\n")
	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
	}

	sort.Ints(numbers)

	var d1 int
	var d3 int
	for i := 0; i < len(numbers); i++ {
		last := 0
		if i > 0 {
			last = numbers[i-1]
		}
		switch numbers[i] - last {
		case 1:
			d1++
		case 3:
			d3++
		}
	}
	d3++
	return strconv.Itoa(d1 * d3)
}

// D10b solces second part of day 10.
func D10b(input string) string {
	lines := strings.Split(input[:len(input)-1], "\n")
	numbers := make([]int, len(lines), len(lines)+2)
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
	}
	numbers = append(numbers, 0)
	sort.Ints(numbers)

	cache := make(map[int]int, len(numbers))
	v := value(numbers, 0, cache)

	return strconv.Itoa(v)
}

func value(numbers []int, idx int, cache map[int]int) int {
	if v, ok := cache[idx]; ok {
		return v
	}

	if idx == len(numbers)-1 {
		return 1
	}

	var count int
	for i := idx + 1; i <= idx+3; i++ {
		if i >= len(numbers) || numbers[i] > numbers[idx]+3 {
			break
		}

		count += value(numbers, i, cache)
	}

	cache[idx] = count
	return count
}
