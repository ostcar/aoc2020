package d6

import (
	"strconv"
	"strings"
)

// D6a solves first part of day6.
func D6a(input string) string {
	input = strings.ReplaceAll(input, "\n\n", "$")
	input = strings.ReplaceAll(input, "\n", "")
	data := strings.Split(input, "$")

	var count int
	set := make(map[rune]bool, 26)
	for _, group := range data {
		for k := range set {
			delete(set, k)
		}
		for _, char := range group {
			set[char] = true
		}
		count += len(set)
	}

	return strconv.Itoa(count)
}

//D6b solves second part of day6.
func D6b(input string) string {
	data := strings.Split(input[:len(input)-1], "\n\n")

	var count int
	for _, group := range data {
		persons := strings.Count(group, "\n") + 1

		for i := 'a'; i <= 'z'; i++ {
			if strings.Count(group, string(i)) == persons {
				count++
			}
		}
	}

	return strconv.Itoa(count)

}
