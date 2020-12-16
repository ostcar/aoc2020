package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ostcar/aoc2020/d1"
	"github.com/ostcar/aoc2020/d10"
	"github.com/ostcar/aoc2020/d11"
	"github.com/ostcar/aoc2020/d12"
	"github.com/ostcar/aoc2020/d13"
	"github.com/ostcar/aoc2020/d14"
	"github.com/ostcar/aoc2020/d15"
	"github.com/ostcar/aoc2020/d2"
	"github.com/ostcar/aoc2020/d3"
	"github.com/ostcar/aoc2020/d4"
	"github.com/ostcar/aoc2020/d5"
	"github.com/ostcar/aoc2020/d6"
	"github.com/ostcar/aoc2020/d7"
	"github.com/ostcar/aoc2020/d8"
	"github.com/ostcar/aoc2020/d9"
)

func assignment(name string) (input string, fn func(string) string, err error) {
	assignments := map[string]func(string) string{
		"d1a":  d1.D1a,
		"d1b":  d1.D1b,
		"d2a":  d2.D2a,
		"d2b":  d2.D2b,
		"d3a":  d3.D3a,
		"d3b":  d3.D3b,
		"d4a":  d4.D4a,
		"d4b":  d4.D4b,
		"d5a":  d5.D5a,
		"d5b":  d5.D5b,
		"d6a":  d6.D6a,
		"d6b":  d6.D6b,
		"d7a":  d7.D7a,
		"d7b":  d7.D7b,
		"d8a":  d8.D8a,
		"d8b":  d8.D8b,
		"d9a":  d9.D9a,
		"d9b":  d9.D9b,
		"d10a": d10.D10a,
		"d10b": d10.D10b,
		"d11a": d11.D11a,
		"d11b": d11.D11b,
		"d12a": d12.D12a,
		"d12b": d12.D12b,
		"d13a": d13.D13a,
		"d13b": d13.D13b,
		"d14a": d14.D14a,
		"d14b": d14.D14b,
		"d15a": d15.D15a,
		"d15b": d15.D15b,
	}

	fn, ok := assignments[name]
	if !ok {
		return "", nil, fmt.Errorf("unknown assignment %s", name)
	}

	day := name[:len(name)-1]
	file := fmt.Sprintf("%s/input.txt", day)

	f, err := os.Open(file)
	if err != nil {
		return "", nil, fmt.Errorf("open file %s: %w", file, err)
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return "", nil, fmt.Errorf("reading file %s: %w", file, err)
	}
	return string(bs), fn, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s d1a\n", os.Args[0])
		os.Exit(1)
	}

	input, fn, err := assignment(os.Args[1])
	if err != nil {
		log.Fatalf("Can not get assignment: %v", err)
	}

	result := fn(input)
	if err != nil {
		log.Fatalf("Can not solve %s: %v", os.Args[1], err)
	}
	fmt.Println(result)
}
