package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/ostcar/aoc2020/d1"
	"github.com/ostcar/aoc2020/d2"
	"github.com/ostcar/aoc2020/d3"
	"github.com/ostcar/aoc2020/d4"
	"github.com/ostcar/aoc2020/d5"
)

func input(day string) (io.ReadCloser, error) {
	f, err := os.Open(day + ".txt")
	if err != nil {
		return nil, fmt.Errorf("can not open file: %w", err)
	}
	return f, nil
}

func assignment(name string) (input string, fn func(string) (string, error), err error) {
	assignments := map[string]func(string) (string, error){
		"d1a": d1.D1a,
		"d1b": d1.D1b,
		"d2a": d2.D2a,
		"d2b": d2.D2b,
		"d3a": d3.D3a,
		"d3b": d3.D3b,
		"d4a": d4.D4a,
		"d4b": d4.D4b,
		"d5a": d5.D5a,
		"d5b": d5.D5b,
	}

	fn, ok := assignments[name]
	if !ok {
		return "", nil, fmt.Errorf("unknown assignment %s", name)
	}

	day := name[:len(name)-1]
	file := fmt.Sprintf("%s/%s.txt", day, day)

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

	result, err := fn(input)
	if err != nil {
		log.Fatalf("Can not solve %s: %v", os.Args[1], err)
	}
	fmt.Println(result)
}
