package main

import (
	"fmt"
	"strings"

	"github.com/borgstrom/aoc2021/input"
)

func containsAll(s1, s2 string) bool {
	x1 := make(map[rune]bool)
	x2 := make(map[rune]bool)
	for _, x := range s1 {
		x1[x] = true
	}
	for _, x := range s2 {
		x2[x] = true
	}
	for k, v := range x1 {
		if x2[k] != v {
			return false
		}
	}
	return true
}

func containsDiff(s1, s2 string) string {
	x1 := make(map[rune]bool)
	x2 := make(map[rune]bool)
	for _, x := range s1 {
		x1[x] = true
	}
	for _, x := range s2 {
		x2[x] = true
	}

	out := ""
	for k, v := range x1 {
		if x2[k] != v {
			out += string(k)
		}
	}
	return out
}

func digitPattern(s string) string {
	parts := strings.Split(s, " | ")
	digits := strings.Fields(parts[1])

	out := make([]string, 4)
	found := make(map[int]string)

	// Pass 1, easy digits
	for i, d := range digits {
		switch len(d) {
		case 2:
			out[i] = "1"
			found[1] = d
		case 3:
			out[i] = "7"
			found[7] = d
		case 4:
			out[i] = "4"
			found[4] = d
		case 7:
			out[i] = "8"
			found[8] = d
		}
	}

	_, found1 := found[1]
	_, found4 := found[4]

	if found1 {
		for i, d := range digits {
			if len(d) == 5 && containsAll(d, found[1]) {
				out[i] = "3"
				found[3] = d
				break
			}
		}

		if found4 {
			for i, d := range digits {
				if len(d) == 5 && containsAll(d, containsDiff(found[1], found[4])) {
					out[i] = "5"
					found[3] = d
					break
				}
			}
		}
	}

	if found4 {
		for i, d := range digits {
			if len(d) == 6 && containsAll(d, found[4]) {
				out[i] = "9"
				found[9] = d
			}
		}
	}

	_, found3 := found[3]
	_, found5 := found[5]

	if found3 && found5 {
		for i, d := range digits {
			if len(d) == 5 {
				out[i] = "2"
				found[2] = d
				break
			}
		}
	}

	return strings.Join(out, "")
}

type Digit map[rune]bool

type Digits []Digit

func parseDigits(s string) Digits {
	output := make(Digits, 4)
	fields := strings.Fields(s)
	if len(fields) != 4 {
		panic("output does not have 4 fields")
	}
	for i, d := range fields {
		output[i] = parseDigit(d)
	}
	return output
}

func parseDigit(s string) Digit {
	out := make(Digit)
	for _, x := range s {
		out[x] = true
	}
	return out
}

func (d Digits) easyDigits() Digits {
	out := make(Digits, 0)

	for _, x := range d {
		switch len(x) {
		case 2:
			// 1
			out = append(out, x)
		case 4:
			// 4
			out = append(out, x)
		case 3:
			// 7
			out = append(out, x)
		case 7:
			// 8
			out = append(out, x)
		}
	}

	return out
}

func countEasyDigits(raw []string) int {
	count := 0
	for _, s := range raw {
		parts := strings.Split(s, " | ")
		d := parseDigits(parts[1])
		u := d.easyDigits()
		count += len(u)
	}
	return count
}

func main() {
	raw := input.MustLoad()
	count := countEasyDigits(raw)
	fmt.Println(count)
}
