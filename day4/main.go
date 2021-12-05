package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/borgstrom/aoc2021/input"
)

type board struct {
	data    [5][5]int
	matched [5][5]bool
	matches [][2]int
}

type boards []*board

func (b *board) match(n int) bool {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if b.data[x][y] == n {
				b.matched[x][y] = true
				b.matches = append(b.matches, [2]int{x, y})
				return true
			}
		}
	}
	return false
}

func (b *board) winner() bool {
	// horizontal
	for x := 0; x < 5; x++ {
		if b.matched[x] == [5]bool{true, true, true, true, true} {
			return true
		}
	}

	// vertical
	for y := 0; y < 5; y++ {
		if b.matched[0][y] && b.matched[1][y] && b.matched[2][y] && b.matched[3][y] && b.matched[4][y] {
			return true
		}
	}

	return false
}

func (b *board) score() int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !b.matched[x][y] {
				sum += b.data[x][y]
			}
		}
	}
	return sum
}

func main() {
	raw := input.MustLoad()
	if len(os.Args) == 2 && os.Args[1] == "last" {
		findLast(raw)
	} else {
		findFirst(raw)
	}
}

func findLast(raw []string) {
	numbers := parseNumbers(raw[0])
	boards := parseBoards(raw[2:])

	won := make(map[int]bool)

	for _, x := range numbers {
		for n, b := range boards {
			if _, ok := won[n]; !ok {
				b.match(x)
				if b.winner() {
					won[n] = true
					if len(won) == len(boards) {
						fmt.Println("Winner! Board number", n)
						fmt.Println("Score:", b.score())
						fmt.Println("Answer:", b.score()*x)
						return
					}
				}
			}
		}
	}
}

func findFirst(raw []string) {
	numbers := parseNumbers(raw[0])
	boards := parseBoards(raw[2:])

	for _, x := range numbers {
		for n, b := range boards {
			b.match(x)
			if b.winner() {
				fmt.Println("Winner! Board number", n)
				fmt.Println("Score:", b.score())
				fmt.Println("Answer:", b.score() * x)
				return
			}
		}
	}
}

func parseNumbers(raw string) []int {
	s := strings.Split(raw, ",")
	l := len(s)
	n := make([]int, l)
	for x := 0; x < l; x++ {
		i, err := strconv.Atoi(s[x])
		if err != nil {
			panic(err)
		}
		n[x] = i
	}
	return n
}

func parseBoards(raw []string) boards {
	boards := make(boards, 0)
	for i := 0; i < len(raw); i += 6 {
		if raw[i] == "" {
			continue
		}

		b := &board{}
		for n := 0; n < 5; n++ {
			for x, s := range strings.Fields(raw[i+n]) {
				i, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				b.data[n][x] = i
			}
		}
		boards = append(boards, b)
	}
	return boards
}
