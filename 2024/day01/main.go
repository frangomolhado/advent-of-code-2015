package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/leakedmemory/aoc/helpers"
)

func part1(lines []string) int {
	left := []int{}
	right := []int{}
	for _, line := range lines {
		pair := strings.Split(line, "   ")
		lval, _ := strconv.Atoi(pair[0])
		rval, _ := strconv.Atoi(pair[1])
		left = append(left, lval)
		right = append(right, rval)
	}

	sfunc := func(a, b int) int {
		return a - b
	}
	slices.SortFunc(left, sfunc)
	slices.SortFunc(right, sfunc)

	result := 0
	for i := range left {
		result += int(math.Abs(float64(right[i] - left[i])))
	}

	return result
}

func part2(lines []string) int {
	left := map[string]struct{}{}
	right := map[string]int{}
	for _, line := range lines {
		pair := strings.Split(line, "   ")
		left[pair[0]] = struct{}{}
		right[pair[1]]++
	}

	result := 0
	for k := range left {
		ik, _ := strconv.Atoi(k)
		result += ik * right[k]
	}

	return result
}

func main() {
	lines := helpers.ReadLines("2024/day01/input.txt")

	fmt.Printf("part1: %v\n", part1(lines))
	fmt.Printf("part2: %v\n", part2(lines))
}
