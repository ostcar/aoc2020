package d8

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// D8a solves first part of day 8.
func D8a(input string) string {
	s, err := newState(input)
	if err != nil {
		return fmt.Sprintf("Can not create state: %v", err)
	}

	s.Run()

	return strconv.Itoa(s.acumulator)
}

// D8b solves second part of day 8.
func D8b(input string) string {
	s, err := newState(input)
	if err != nil {
		return fmt.Sprintf("Can not create state: %v", err)
	}

	for i := 0; i < len(s.code); i++ {
		t := s.code[i]

		switch v := s.code[i].(type) {
		case nop:
			s.code[i] = jmp(v)
		case jmp:
			s.code[i] = nop(v)
		default:
			continue
		}

		if err := s.Run(); err == nil {
			break
		}

		s.code[i] = t

	}

	return strconv.Itoa(s.acumulator)
}

type state struct {
	index      int
	acumulator int
	called     map[int]bool
	code       []operation
}

func newState(input string) (*state, error) {
	lines := strings.Split(input, "\n")
	s := &state{
		called: make(map[int]bool),
		code:   make([]operation, 0, len(lines)),
	}

	for i, line := range lines {
		if line == "" {
			continue
		}

		var cmd string
		var arg int
		if _, err := fmt.Sscanf(line, "%s %d", &cmd, &arg); err != nil {
			return nil, fmt.Errorf("scanning line %d: %v", i, err)
		}

		var op operation
		switch cmd {
		case "nop":
			op = nop(arg)
		case "acc":
			op = acc(arg)
		case "jmp":
			op = jmp(arg)
		default:
			return nil, fmt.Errorf("unknown command `%s`", cmd)
		}

		s.code = append(s.code, op)
	}
	return s, nil
}

func (s *state) Run() error {
	for k := range s.called {
		delete(s.called, k)
	}
	s.acumulator = 0
	s.index = 0

	for s.index < len(s.code) {
		// Do no call same line twice.
		if s.called[s.index] {
			return errInititeLoop
		}

		s.called[s.index] = true
		s.code[s.index].Run(s)
		s.index++
	}
	return nil
}

var errInititeLoop = errors.New("infinite loop")

type operation interface {
	Run(*state)
}

type nop int

func (nop) Run(s *state) {}

type acc int

func (i acc) Run(s *state) {
	s.acumulator += int(i)
}

type jmp int

func (j jmp) Run(s *state) {
	s.index += int(j) - 1
}
