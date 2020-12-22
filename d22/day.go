package d22

import (
	"fmt"
	"strconv"
	"strings"
)

//var debug = fmt.Printf
var debug = func(string, ...interface{}) {}

// D22a solves first part of day 22.
func D22a(input string) string {
	d1, d2 := parse(input)

	win := winNo
	for win == winNo {
		win = gameA(d1, d2)
	}

	w := d1
	if win == win2 {
		w = d2
	}

	return strconv.Itoa(w.score())
}

func gameA(d1, d2 *deck) int {
	if d1.top == nil {
		return win2
	}
	if d2.top == nil {
		return win1
	}

	w := d1
	l := d2
	if d1.top.value < d2.top.value {
		w = d2
		l = d1
	}

	c1, c2 := w.top, l.top
	w.top = w.top.next
	l.top = l.top.next

	c1.next = c2
	c2.next = nil
	w.bottom.next = c1
	w.bottom = c2

	w.length++
	l.length--

	return winNo
}

// D22b solves first part of day 22.
func D22b(input string) string {
	d1, d2 := parse(input)

	win := gameB(d1, d2, 1)

	w := d1
	if win == win2 {
		w = d2
	}

	return strconv.Itoa(w.score())
}

func parse(input string) (d1 *deck, d2 *deck) {
	d1 = new(deck)
	d2 = new(deck)
	d := d1
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		var player int
		if _, err := fmt.Sscanf(line, "Player %d:", &player); err == nil {
			if player == 2 {
				d = d2
			}
			continue
		}

		n, _ := strconv.Atoi(line)
		c := &card{value: n}
		d.length++
		if d.top == nil {
			d.top = c
			d.bottom = c
			continue
		}

		d.bottom.next = c
		d.bottom = c
	}
	return d1, d2
}

func gameB(d1, d2 *deck, g int) int {
	debug("=== Game %d ===\n", g)

	hashes := make(map[string]bool)
	var i int
	for {
		h := d1.hash() + d2.hash()
		if hashes[h] {
			return win1
		}
		hashes[h] = true

		if d1.top == nil {
			debug("The winner of game %d is player 2!\n", g)
			return win2
		}
		if d2.top == nil {
			debug("The winner of game %d is player 1!\n", g)
			return win1
		}

		i++
		debug("\n-- Round %d (Game %d) --\n", i, g)
		debug("Player 1's deck: %s\n", d1)
		debug("Player 2's deck: %s\n", d2)
		debug("Player 1 plays: %s\n", d1.top)
		debug("Player 2 plays: %s\n", d2.top)

		w := d1
		l := d2
		p := 1

		if d1.top.value <= d1.length-1 && d2.top.value <= d2.length-1 {
			debug("Playing a sub-game to determine the winner...\n\n")
			nd1 := d1.copy()
			nd2 := d2.copy()
			nd1.top = nd1.top.next
			nd2.top = nd2.top.next
			nd1.length--
			nd2.length--

			if won := gameB(nd1, nd2, g+1); won == win2 {
				w = d2
				l = d1
				p = 2
			}
			debug("\n...anyway, back to game %d.\n", g)
		} else {
			if d1.top.value < d2.top.value {
				w = d2
				l = d1
				p = 2
			}
		}

		debug("Player %d wins round %d of game %d!\n", p, i, g)

		c1, c2 := w.top, l.top
		w.top = w.top.next
		l.top = l.top.next

		c1.next = c2
		c2.next = nil
		w.bottom.next = c1
		w.bottom = c2

		w.length++
		l.length--
	}
}

type deck struct {
	top    *card
	bottom *card
	length int
}

func (d *deck) hash() string {
	var cards []string
	for c := d.top; c != nil; c = c.next {
		cards = append(cards, strconv.Itoa(c.value))
	}
	return "D" + strings.Join(cards, ",")
}

func (d *deck) copy() *deck {
	nd := new(deck)
	nd.length = d.length
	for n := d.top; n != nil; n = n.next {
		if nd.top == nil {
			nd.top = &card{value: n.value}
			nd.bottom = nd.top
			continue
		}

		nd.bottom.next = &card{value: n.value}
		nd.bottom = nd.bottom.next
	}
	return nd
}

func (d *deck) score() int {
	multi := d.length
	var result int
	for c := d.top; c != nil; c = c.next {
		result += c.value * multi
		multi--
	}
	return result
}

func (d *deck) String() string {
	var cards []string
	for c := d.top; c != nil; c = c.next {
		cards = append(cards, strconv.Itoa(c.value))
	}
	return strings.Join(cards, ", ")
}

type card struct {
	value int
	next  *card
}

func (c *card) String() string {
	if c == nil {
		return ""
	}
	return strconv.Itoa(c.value)
}

const (
	winNo = iota
	win1
	win2
)
