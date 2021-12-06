package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/borgstrom/aoc2021/input"
)

type line struct {
	x1, x2, y1, y2 int
}

func (l *line) horizontal() bool {
	return l.x1 == l.x2
}

func (l *line) vertical() bool {
	return l.y1 == l.y2
}

func (l *line) minX() int {
	if l.x1 < l.x2 {
		return l.x1
	}
	return l.x2
}

func (l *line) maxX() int {
	if l.x1 > l.x2 {
		return l.x1
	}
	return l.x2
}

func (l *line) minY() int {
	if l.y1 < l.y2 {
		return l.y1
	}
	return l.y2
}

func (l *line) maxY() int {
	if l.y1 > l.y2 {
		return l.y1
	}
	return l.y2
}

func parseLine(raw string) *line {
	parts := strings.Split(raw, "->")
	xy1 := strings.Split(parts[0], ",")
	xy2 := strings.Split(parts[1], ",")

	return &line{
		x1: input.MustAtoi(xy1[0]),
		y1: input.MustAtoi(xy1[1]),
		x2: input.MustAtoi(xy2[0]),
		y2: input.MustAtoi(xy2[1]),
	}
}

func parseLines(raw []string) []*line {
	out := make([]*line, len(raw))
	for i, l := range raw {
		out[i] = parseLine(l)
	}
	return out
}

func countOverlap(lines []*line, diagonal bool) int {
	counter := make(map[[2]int]int)

	for _, l := range lines {
		if l.horizontal() {
			for y := l.minY(); y <= l.maxY(); y++ {
				counter[[2]int{l.x1, y}]++
			}
		} else if l.vertical() {
			for x := l.minX(); x <= l.maxX(); x++ {
				counter[[2]int{x, l.y1}]++
			}
		} else if diagonal {
			var startX, startY, endX, endY int
			if l.x2 > l.x1 {
				startX = l.x1
				startY = l.y1
				endX = l.x2
				endY = l.y2
			} else {
				startX = l.x2
				startY = l.y2
				endX = l.x1
				endY = l.y1
			}
			for i := 0; i <= endX-startX; i++ {
				var y int
				if startY < endY {
					y = startY + i
				} else {
					y = startY - i
				}
				counter[[2]int{startX + i, y}]++
			}
		}
	}

	count := 0
	for _, c := range counter {
		if c > 1 {
			count++
		}
	}
	return count
}

func main() {
	raw := input.MustLoad()
	lines := parseLines(raw)

	count := countOverlap(lines, len(os.Args) == 2 && os.Args[1] == "diagonal")
	fmt.Println(count)
}
