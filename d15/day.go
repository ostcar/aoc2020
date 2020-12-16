package d15

import (
	"strconv"
	"strings"
)

// D15a solves first part of day 15.
func D15a(input string) string {
	input = strings.TrimSpace(input)
	inputNr := strings.Split(input, ",")

	turns := make(map[int]int)
	for i, v := range inputNr[:len(inputNr)-1] {
		vtoi, _ := strconv.Atoi(v)
		turns[vtoi] = i + 1
	}

	last, _ := strconv.Atoi(inputNr[len(inputNr)-1])

	for i := len(inputNr); i < 2020; i++ {
		before, ok := turns[last]
		turns[last] = i
		if !ok {
			last = 0
			continue
		}

		last = i - before
	}

	return strconv.Itoa(last)
}

// D15b solves first part of day 15.
func D15b(input string) string {
	input = strings.TrimSpace(input)
	inputNr := strings.Split(input, ",")

	turns := make(map[int]int)
	for i, v := range inputNr[:len(inputNr)-1] {
		vtoi, _ := strconv.Atoi(v)
		turns[vtoi] = i + 1
	}

	last, _ := strconv.Atoi(inputNr[len(inputNr)-1])

	for i := len(inputNr); i < 30_000_000; i++ {
		before, ok := turns[last]
		turns[last] = i
		if !ok {
			last = 0
			continue
		}

		last = i - before
	}

	return strconv.Itoa(last)
}
