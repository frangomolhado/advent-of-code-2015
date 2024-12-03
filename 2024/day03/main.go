package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/leakedmemory/aoc/helpers"
)

const mulPattern = `mul\((\d{1,3}),(\d{1,3})\)`

var mulRegex = regexp.MustCompile(mulPattern)

func part1(input string) int {
	result := 0
	matches := mulRegex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		result += x * y
	}

	return result
}

func part2(input string) int {
	result := 0
	dos := strings.Split(input, "do()")
	for _, do := range dos {
		seq := strings.Split(do, "don't()")[0]
		matches := mulRegex.FindAllStringSubmatch(seq, -1)
		for _, match := range matches {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			result += x * y
		}
	}

	return result
}

func main() {
	input := helpers.ReadFile("2024/day03/input.txt")

	fmt.Printf("part1: %v\n", part1(input))
	fmt.Printf("part2: %v\n", part2(input))
}
