package d14

import (
	"fmt"
	"strconv"
	"strings"
)

// D14a solves first part of day XX.
func D14a(input string) string {
	var mask0, mask1 int
	mem := make(map[int]int)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.SplitN(line, "=", 2)
		cmd := strings.TrimSpace(parts[0])

		if cmd == "mask" {
			mask0, mask1 = parseMaskA(strings.TrimSpace(parts[1]))
			continue
		}

		if strings.HasPrefix(cmd, "mem") {
			var addr, value int
			fmt.Sscanf(line, "mem[%d] = %d", &addr, &value)
			mem[addr] = applyMaskA(value, mask0, mask1)
			continue
		}
	}

	var sum int
	for _, v := range mem {
		sum += v
	}
	return strconv.Itoa(sum)
}

// D14b solves first part of day XX.
func D14b(input string) string {
	var mask1, maskX int
	mem := make(map[int]int)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.SplitN(line, "=", 2)
		cmd := strings.TrimSpace(parts[0])

		if cmd == "mask" {
			mask1, maskX = parseMaskB(strings.TrimSpace(parts[1]))
			continue
		}

		if strings.HasPrefix(cmd, "mem") {
			var addr, value int
			if _, err := fmt.Sscanf(line, "mem[%d] = %d", &addr, &value); err != nil {
				return fmt.Sprintf("Can not parse line `%s`: %v", line, err)
			}
			applyMaskB(mem, addr, value, mask1, maskX)
			continue
		}
	}

	var sum int
	for _, v := range mem {
		sum += v
	}
	return strconv.Itoa(sum)
}

func parseMaskA(mask string) (int, int) {
	var mask0, mask1 int
	for _, c := range mask {
		mask1 <<= 1
		mask0 <<= 1
		if c == '1' {
			mask1 |= 1
		}
		if c != '0' {
			mask0 |= 1
		}
	}
	return mask0, mask1
}

func applyMaskA(value, mask0, mask1 int) int {
	value |= mask1
	value &= mask0
	return value
}

func parseMaskB(mask string) (int, int) {
	var mask1, maskX int
	for _, c := range mask {
		maskX <<= 1
		mask1 <<= 1
		if c == 'X' {
			maskX |= 1
		}
		if c == '1' {
			mask1 |= 1
		}
	}
	return mask1, maskX
}

func applyMaskB(mem map[int]int, addr, value, mask1, maskX int) {
	addr |= mask1

	addrFloating := []int{addr}

	for i := 0; i < 36; i++ {
		if maskX&(1<<i) == 0 {
			continue
		}
		addrFloating = float(addrFloating, i)
	}

	for _, v := range addrFloating {
		mem[v] = value
	}
}

func float(addrList []int, bit int) []int {
	l := len(addrList)
	for i := 0; i < l; i++ {
		addrList[i] &= ^(1 << bit)

		v := addrList[i]
		v |= 1 << bit
		addrList = append(addrList, v)
	}
	return addrList
}
