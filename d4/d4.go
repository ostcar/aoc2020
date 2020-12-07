package d4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// D4a solves first part of day 4.
func D4a(input string) string {
	var c int
	for _, p := range parse(input) {
		if validSimple(p) {
			c++
		}
	}
	return strconv.Itoa(c)
}

// D4b solves second part of day 4.
func D4b(input string) string {
	var c int
	for _, p := range parse(input) {
		if validKomplex(p) {
			c++
		}
	}
	return strconv.Itoa(c)
}

func validSimple(p map[string]string) bool {
	for _, key := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if _, ok := p[key]; !ok {
			return false
		}
	}
	return true
}

func validKomplex(p map[string]string) bool {
	for key, f := range map[string]func(string) bool{
		"byr": byr,
		"iyr": iyr,
		"eyr": eyr,
		"hgt": hgt,
		"hcl": hcl,
		"ecl": ecl,
		"pid": pid,
	} {
		v, ok := p[key]
		if !ok {
			return false
		}
		if !f(v) {
			return false
		}
	}
	return true
}

func checkYear(in string, min, max int) bool {
	var year int
	if _, err := fmt.Sscanf(in, "%d", &year); err != nil {
		return false
	}

	if year < min || year > max {
		return false
	}
	return true
}

func byr(in string) bool {
	return checkYear(in, 1920, 2002)
}

func iyr(in string) bool {
	return checkYear(in, 2010, 2020)
}

func eyr(in string) bool {
	return checkYear(in, 2020, 2030)
}

func hgt(in string) bool {
	var number int
	var unit string
	if _, err := fmt.Sscanf(in, "%d%s", &number, &unit); err != nil {
		return false
	}

	switch unit {
	case "cm":
		return number >= 150 && number <= 193
	case "in":
		return number >= 59 && number <= 76
	default:
		return false
	}
}

func hcl(in string) bool {
	ok, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, in)
	return ok
}

func ecl(in string) bool {
	for _, accept := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if in == accept {
			return true
		}
	}
	return false
}

func pid(in string) bool {
	ok, _ := regexp.MatchString(`^[0-d]{9}$`, in)
	return ok
}

func parse(input string) []map[string]string {
	var p []map[string]string

	for _, element := range strings.Split(input, "\n\n") {
		passport := make(map[string]string)
		for _, field := range strings.Fields(element) {
			parts := strings.Split(field, ":")
			passport[parts[0]] = parts[1]
		}
		p = append(p, passport)
	}
	return p
}
