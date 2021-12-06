package main

import (
	"fmt"
	"strings"

	"github.com/borgstrom/aoc2021/input"
)

type school map[int]int

func (s school) tick() {
	s0 := s[0]
	for i := 0; i < 8; i++ {
		s[i] = s[i+1]
	}
	s[6] += s0
	s[8] = s0
}

func (s school) count() int {
	count := 0
	for i := 0; i <= 8; i++ {
		count += s[i]
	}
	return count
}

func loadInput(raw string) school {
	nums := strings.Split(raw, ",")

	s := school{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}

	for _, num := range nums {
		s[input.MustAtoi(num)]++
	}

	return s
}
func main() {
	i := input.MustLoad()
	s := loadInput(i[0])
	for x := 0; x < 256; x++ {
		s.tick()
	}
	fmt.Println(s.count())
}
