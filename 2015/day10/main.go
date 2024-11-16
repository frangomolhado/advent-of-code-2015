package main

import (
	"fmt"
	"strings"
)

func part1(input string) int {
	return len(lookAndSay(input, 40))
}

func part2(input string) int {
	return len(lookAndSay(input, 50))
}

func lookAndSay(start string, iterations int) string {
	curr := start
	for range iterations {
		var sb strings.Builder
		lastC := rune(curr[0])
		count := 1
		for _, c := range curr[1:] {
			if c == lastC {
				count++
				continue
			}

			sb.WriteString(fmt.Sprintf("%d%c", count, lastC))
			lastC = c
			count = 1
		}

		sb.WriteString(fmt.Sprintf("%d%c", count, lastC))
		curr = sb.String()
	}

	return curr
}

func main() {
	input := "1113222113"

	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}
