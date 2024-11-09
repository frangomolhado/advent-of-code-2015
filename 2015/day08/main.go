package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/leakedmemory/aoc/helpers"
)

func part1(input string) (result, codeSum int) {
	stringSum := 0
	for _, line := range strings.Split(input, "\n") {
		for i := 0; i < len(line); i++ {
			if !unicode.IsLetter(rune(line[i])) && rune(line[i]) == '\\' {
				if rune(line[i+1]) == '"' || rune(line[i+1]) == '\\' {
					codeSum += 2
					i++
				} else {
					codeSum += 4
					i += 3
				}
				stringSum++
			} else if unicode.IsLetter(rune(line[i])) {
				stringSum++
				codeSum++
			} else {
				codeSum++
			}
		}
	}

	result = codeSum - stringSum
	return
}

func part2(input string) int {
	result := 0
	for _, line := range strings.Split(input, "\n") {
		result += 2 // opening and closing quotes
		for _, c := range line {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				result++
			} else {
				result += 2
			}
		}
	}

	return result
}

func main() {
	input := helpers.ReadFile("2015/day08/input.txt")

	p1, c1 := part1(input)
	p2 := part2(input) - c1
	fmt.Printf("part1: %d\n", p1)
	fmt.Printf("part2: %d\n", p2)
}
