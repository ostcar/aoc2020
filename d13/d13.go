package d13

import (
	"fmt"
	"strconv"
	"strings"
)

// D13a solves first part of day 13.
func D13a(input string) string {
	lines := strings.Split(input, "\n")
	arrive, _ := strconv.Atoi(lines[0])

	minWait := arrive
	var busID int
	for _, b := range strings.Split(lines[1], ",") {
		if b == "x" {
			continue
		}
		bus, _ := strconv.Atoi(b)

		if newMinWait := min(minWait, bus-(arrive%bus)); newMinWait != 0 {
			minWait = newMinWait
			busID = bus
		}
	}

	return strconv.Itoa(minWait * busID)
}

// D13b solves first part of day 13.
func D13b(input string) string {
	lines := strings.Split(input, "\n")
	rawBusIDs := strings.Split(lines[1], ",")
	busIDs := make([]int, len(rawBusIDs))
	var highestID int
	for i, r := range rawBusIDs {
		if r == "x" {
			continue
		}
		busIDs[i], _ = strconv.Atoi(r)
		if busIDs[i] > highestID {
			highestID = i
		}
	}

	var count int
	for t := busIDs[highestID] - highestID; ; t += busIDs[highestID] {
		count++
		if count%1_000_000_000 == 0 {
			fmt.Println(t)
		}

		if validate(t, busIDs) {
			return strconv.Itoa(t)
		}
	}
}

func validate(t int, busIDs []int) bool {
	for i, bus := range busIDs {
		if bus == 0 {
			continue
		}
		if (t+i)%bus != 0 {
			return false
		}
	}
	return true
}

// min returns b, if it is lower then a, in other case, it returns 0.
func min(a, b int) int {
	if b < a {
		return b
	}
	return 0
}
