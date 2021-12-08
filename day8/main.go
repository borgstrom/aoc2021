package main

import (
	"fmt"
	"strings"

	"github.com/borgstrom/aoc2021/input"
)

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
