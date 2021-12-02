package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/borgstrom/aoc2021/input"
)

/*
--- Day 2: Dive! ---

Now, you need to figure out how to pilot this thing.

It seems like the submarine can take a series of commands like forward 1, down 2, or up 3:

    forward X increases the horizontal position by X units.
    down X increases the depth by X units.
    up X decreases the depth by X units.

Note that since you're on a submarine, down and up affect your depth, and so they have the opposite result of what you might expect.

The submarine seems to already have a planned course (your puzzle input). You should probably figure out where it's going. For example:

forward 5
down 5
forward 8
up 3
down 8
forward 2

Your horizontal position and depth both start at 0. The steps above would then modify them as follows:

    forward 5 adds 5 to your horizontal position, a total of 5.
    down 5 adds 5 to your depth, resulting in a value of 5.
    forward 8 adds 8 to your horizontal position, a total of 13.
    up 3 decreases your depth by 3, resulting in a value of 2.
    down 8 adds 8 to your depth, resulting in a value of 10.
    forward 2 adds 2 to your horizontal position, a total of 15.

After following these instructions, you would have a horizontal position of 15 and a depth of 10. (Multiplying these together produces 150.)

Calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?

 */

type command struct {
	direction string
	amount int
}

func loadCommands() []command {
	out := make([]command, 0)
	for _, s := range input.MustLoad() {
		parts := strings.Split(s, " ")
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(fmt.Errorf("failed to convert command amount: %w", err))
		}
		out = append(out, command{
			direction: parts[0],
			amount: amount,
		})
	}
	return out
}

var useAim = flag.Bool("aim", false, "Use aim to calculate depth")

func main() {
	flag.Parse()
	if *useAim {
		part2()
	} else {
		part1()
	}
}

func part1() {
	commands := loadCommands()
	horizontal := 0
	depth := 0
	for _, command := range commands {
		switch command.direction {
		case "forward":
			horizontal += command.amount
		case "down":
			depth += command.amount
		case "up":
			depth -= command.amount
		default:
			panic(fmt.Errorf("unknown command: %s", command.direction))
		}
	}

	fmt.Println("horizontal:", horizontal)
	fmt.Println("depth:", depth)
	fmt.Println("horizontal * depth:", horizontal*depth)
}

func part2() {
	commands := loadCommands()
	horizontal := 0
	depth := 0
	aim := 0

	for _, command := range commands {
		switch command.direction {
		case "forward":
			horizontal += command.amount
			depth += command.amount * aim
		case "down":
			aim += command.amount
		case "up":
			aim -= command.amount
		default:
			panic(fmt.Errorf("unknown command: %s", command.direction))
		}
	}

	fmt.Println("horizontal:", horizontal)
	fmt.Println("depth:", depth)
	fmt.Println("horizontal * depth:", horizontal*depth)
}
