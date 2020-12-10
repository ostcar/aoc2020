package d10_test

import (
	"testing"

	"github.com/ostcar/aoc2020/d10"
)

const input1 = `16
10
15
5
1
11
7
19
6
12
4
`

const input2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`

func TestD10a(t *testing.T) {
	if got := d10.D10a(input1); got != "35" {
		t.Errorf("Got %s, expected 35", got)
	}

	if got := d10.D10a(input2); got != "220" {
		t.Errorf("Got %s, expected 220", got)
	}
}

func TestD10b(t *testing.T) {
	if got := d10.D10b(input1); got != "8" {
		t.Errorf("Got %s, expected 8", got)
	}

	if got := d10.D10b(input2); got != "19208" {
		t.Errorf("Got %s, expected 19208", got)
	}
}
