package main

import (
	"fmt"
	"log"
	"os"
)

func part1(input string) int {
	// discount line feed
	result := 1
	for _, c := range input {
		if c == '(' {
			result++
		} else {
			result--
		}
	}

	return result
}

func part2(input string) int {
	floor := 0
	for i, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
			if floor == -1 {
				return i + 1
			}
		}
	}

	panic("unreachable")
}

func main() {
	content, err := os.ReadFile("2015/day01/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	input := string(content)
	fmt.Printf("part 1: %d\n", part1(input))
	fmt.Printf("part 2: %d\n", part2(input))
}
