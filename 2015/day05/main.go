package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/leakedmemory/aoc/helpers"
)

var (
	vowels           = []rune{'a', 'e', 'i', 'o', 'u'}
	invalidSequences = []string{"ab", "cd", "pq", "xy"}
)

func part1(input string) int {
	result := 0
	for _, word := range strings.Split(input, "\n") {
		vowelsSeen := []rune{}
		lastC := ' '
		twiceInRow := false
		hasInvalidSequence := false

		for _, c := range word {
			if slices.Contains(vowels, c) {
				vowelsSeen = append(vowelsSeen, c)
			}

			if !twiceInRow && lastC == c {
				twiceInRow = true
				// No need to update lastC since they are equal and no need to
				// check for invalid sequence since it is impossible.
				continue
			}

			if slices.Contains(invalidSequences, string(lastC)+string(c)) {
				hasInvalidSequence = true
				break
			}

			lastC = c
		}

		if len(vowelsSeen) >= 3 && twiceInRow && !hasInvalidSequence {
			result++
		}
	}

	return result
}

func part2(input string) int {
	result := 0
	for _, word := range strings.Split(input, "\n") {
		pairs := make(map[string][]int)
		repeatedWithSingleLetter := false
		overlaps := true
		for i, c := range word {
			if !repeatedWithSingleLetter && i > 1 && c == rune(word[i-2]) {
				repeatedWithSingleLetter = true
			}

			if overlaps && i > 0 {
				pastC := string(word[i-1])
				pair := pastC + string(c)
				if indices, ok := pairs[pair]; ok {
					for _, idx := range indices {
						if i-1 != idx {
							overlaps = false
							break
						}
					}
				}

				pairs[pair] = append(pairs[pair], i)
			}

			if repeatedWithSingleLetter && !overlaps {
				result++
				break
			}
		}
	}

	return result
}

func main() {
	input := helpers.ReadFile("2015/day05/input.txt")

	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}
