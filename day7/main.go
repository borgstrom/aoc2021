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

func (c crabs) align(constant bool) alignment {
	out := alignment{
		start: time.Now(),
		cost: math.MaxInt,
	}

	for i := c.min; i <= c.max; i++ {
		cost := c.alignTo(i, constant)
		if cost < out.cost {
			out.position = i
			out.cost = cost
		}
	}

	out.duration = time.Since(out.start)
	return out
}

func (c crabs) alignTo(pos int, constant bool) (cost int) {
	cost = 0
	for _, i := range c.positions {
		var x int
		if i > pos {
			x = i - pos
		} else if i < pos {
			x = pos - i
		}
		if constant {
			cost += x
		} else {
			for y := 1; y <= x; y++ {
				cost += y
			}
		}
	}
	return
}

func main() {
	raw := input.MustLoad()
	c := loadCrabs(raw[0])
	a := c.align(false)
	spew.Dump(a)
}
