package d7

import (
	"strconv"
	"strings"
)

func parseLine(line string) (string, map[string]int) {
	parts := strings.SplitN(line, " bags contain ", 2)
	subBags := strings.SplitAfter(parts[1], ", ")

	contains := make(map[string]int)
	for _, subBag := range subBags {
		if subBag == "no other bags." {
			continue
		}

		words := strings.Split(subBag, " ")
		name := words[1] + " " + words[2]
		amount, _ := strconv.Atoi(words[0])
		contains[name] = amount
	}
	return strings.TrimSpace(parts[0]), contains
}

func parse(input string) map[string]map[string]int {
	lines := strings.Split(input, "\n")
	bags := make(map[string]map[string]int, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		name, contains := parseLine(line)
		bags[name] = contains
	}
	return bags
}

// D7a solves first part of day 7.
func D7a(input string) string {
	bags := parse(input)
	contained := make(map[string][]string, len(bags))
	for name, contain := range bags {
		for inner := range contain {
			contained[inner] = append(contained[inner], name)
		}
	}
	return strconv.Itoa(len(eventualyContain("shiny gold", contained)))
}

func eventualyContain(name string, contained map[string][]string) map[string]bool {
	out := make(map[string]bool)
	for _, inner := range contained[name] {
		out[inner] = true
		for key := range eventualyContain(inner, contained) {
			out[key] = true
		}
	}
	return out
}

// D7b solves second part of day 7.
func D7b(input string) string {
	bags := parse(input)
	return strconv.Itoa(recContain("shiny gold", bags))
}

func recContain(name string, bags map[string]map[string]int) int {
	var count int
	for contain, amount := range bags[name] {
		count += amount
		count += amount * recContain(contain, bags)
	}
	return count
}
