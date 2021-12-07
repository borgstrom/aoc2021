package main

import (
	"math"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/borgstrom/aoc2021/input"
)

type crabs struct {
	positions []int
	min int
	max int
}

func loadCrabs(s string) crabs {
	in := strings.Split(s, ",")
	out := crabs{
		positions: make([]int, len(in)),
	}

	for i, xS := range in {
		x := input.MustAtoi(xS)
		out.positions[i] = x
		if x < out.min {
			out.min = x
		}
		if x > out.max {
			out.max = x
		}
	}

	return out
}

type alignment struct {
	position int
	cost int
	start time.Time
	duration time.Duration
}

func (c crabs) align() alignment {
	out := alignment{
		start: time.Now(),
		cost: math.MaxInt,
	}

	for _, i := range c.positions {
		cost := c.alignTo(i)
		if cost < out.cost {
			out.position = i
			out.cost = cost
		}
	}

	out.duration = time.Since(out.start)
	return out
}

func (c crabs) alignTo(pos int) (cost int) {
	cost = 0
	for _, i := range c.positions {
		if i > pos {
			cost += i - pos
		} else if i < pos {
			cost += pos - i
		}
	}
	return
}

func main() {
	raw := input.MustLoad()
	c := loadCrabs(raw[0])
	a := c.align()
	spew.Dump(a)
}
