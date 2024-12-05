package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/leakedmemory/aoc/helpers"
)

func part1(lines []string) int {
	result := 0
	for _, line := range lines {
		sLevels := strings.Split(line, " ")
		levels := helpers.SliceAtoi(sLevels)
		if isSafe(levels) {
			result++
		}
	}

	return result
}

func isSafe(levels []int) bool {
	ascending := levels[0] < levels[1]
	lastLevel := levels[0]
	for _, currLevel := range levels[1:] {
		if ascending && currLevel < lastLevel {
			return false
		}

		if !ascending && lastLevel < currLevel {
			return false
		}

		distance := int(math.Abs(float64(currLevel - lastLevel)))
		if distance == 0 || distance > 3 {
			return false
		}

		lastLevel = currLevel
	}

	return true
}

func part2(lines []string) int {
	result := 0
	for _, line := range lines {
		sLevels := strings.Split(line, " ")
		levels := helpers.SliceAtoi(sLevels)
		if isSafe(levels) || isSafeWithOneRemoval(levels) {
			result++
		}
	}

	return result
}

func isSafeWithOneRemoval(levels []int) bool {
	for i := range levels {
		tmp := append([]int{}, levels[:i]...)
		tmp = append(tmp, levels[i+1:]...)
		if isSafe(tmp) {
			return true
		}
	}

	return false
}

func main() {
	lines := helpers.ReadLines("2024/day02/input.txt")

	fmt.Printf("part1: %v\n", part1(lines))
	fmt.Printf("part2: %v\n", part2(lines))
}
