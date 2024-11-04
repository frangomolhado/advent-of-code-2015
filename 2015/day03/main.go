package main

import (
	"fmt"

	"github.com/leakedmemory/aoc/helpers"
)

type point struct {
	x, y int
}

func part1(input string) int {
	set := make(map[point]struct{})
	currPos := point{0, 0}
	set[currPos] = struct{}{}
	for _, c := range input {
		switch c {
		case '>':
			currPos.x++
		case '<':
			currPos.x--
		case '^':
			currPos.y--
		case 'v':
			currPos.y++
		}

		set[currPos] = struct{}{}
	}

	return len(set)
}

func part2(input string) int {
	set := make(map[point]struct{})
	santa := point{0, 0}
	robo := point{0, 0}
	set[santa] = struct{}{}
	for i, c := range input {
		var mover *point
		if i%2 == 0 {
			mover = &santa
		} else {
			mover = &robo
		}

		switch c {
		case '>':
			mover.x++
		case '<':
			mover.x--
		case '^':
			mover.y--
		case 'v':
			mover.y++
		}

		set[*mover] = struct{}{}
	}

	return len(set)
}

func main() {
	input := helpers.ReadFile("2015/day03/input.txt")
	fmt.Printf("part 1: %d\n", part1(input))
	fmt.Printf("part 2: %d\n", part2(input))
}
