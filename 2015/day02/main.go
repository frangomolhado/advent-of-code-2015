package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(input string) int {
	result := 0
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		l, w, h := rectFromString(line)
		side1 := l * w
		side2 := w * h
		side3 := h * l

		result += 2 * (side1 + side2 + side3)
		result += min(min(side1, side2), side3)
	}

	return result
}

func part2(input string) int {
	result := 0
	for _, line := range strings.Split(strings.TrimRight(input, "\n"), "\n") {
		l, w, h := rectFromString(line)
		volume := l * w * h
		result += volume

		min1 := l
		min2 := w
		min3 := h
		if min1 > min2 {
			min1, min2 = min2, min1
		}
		if min1 > min3 {
			min1, min3 = min3, min1
		}
		if min2 > min3 {
			min2 = min3
		}
		result += 2 * (min1 + min2)
	}

	return result
}

func rectFromString(s string) (length, width, height int) {
	dimensions := strings.Split(s, "x")
	length, _ = strconv.Atoi(dimensions[0])
	width, _ = strconv.Atoi(dimensions[1])
	height, _ = strconv.Atoi(dimensions[2])
	return
}

func main() {
	content, err := os.ReadFile("2015/day02/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	input := string(content)
	fmt.Printf("part 1: %d\n", part1(input))
	fmt.Printf("part 2: %d\n", part2(input))
}
