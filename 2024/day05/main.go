package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leakedmemory/aoc/helpers"
)

type rulesMap map[string]map[string]bool

func parseInput(input string) (rulesMap, []string) {
	tmp := strings.Split(input, "\n\n")
	rules := strings.Split(tmp[0], "\n")
	updates := strings.Split(tmp[1], "\n")

	rMap := makeRulesMap(rules)

	return rMap, updates
}

func part1(rMap rulesMap, updates []string) int {
	result := 0
	for _, update := range updates {
		u := strings.Split(update, ",")
		if isCorrectOrder(rMap, u) {
			mid, _ := strconv.Atoi(u[len(u)/2])
			result += mid
		}
	}

	return result
}

func makeRulesMap(rules []string) rulesMap {
	m := rulesMap{}
	for _, r := range rules {
		tmp := strings.Split(r, "|")
		left := tmp[0]
		right := tmp[1]

		if m[left] == nil {
			m[left] = make(map[string]bool)
		}
		if m[right] == nil {
			m[right] = make(map[string]bool)
		}

		// false == parent value should come before child value
		m[left][right] = false
		// true == parent value should come after child value
		m[right][left] = true
	}

	return m
}

func isCorrectOrder(rMap rulesMap, update []string) bool {
	for i, page := range update {
		rules, ok := rMap[page]
		if !ok {
			continue
		}

		for j, p := range update {
			if p == page {
				continue
			}

			after, ok := rules[p]
			if !ok {
				continue
			}

			if after && i < j {
				return false
			}

			if !after && j < i {
				return false
			}
		}
	}

	return true
}

func part2(rMap rulesMap, updates []string) int {
	result := 0
	for _, update := range updates {
		u := strings.Split(update, ",")
		if !isCorrectOrder(rMap, u) {
			fixed := fixUpdateOrder(rMap, u)
			mid, _ := strconv.Atoi(fixed[len(fixed)/2])
			result += mid
		}
	}

	return result
}

// probably should not take 20s to execute, but works for now.
func fixUpdateOrder(rMap rulesMap, update []string) []string {
	fixed := update
Outer:
	for {
		for i, page := range fixed {
			rules, ok := rMap[page]
			if !ok {
				continue
			}

			for j, p := range fixed {
				if p == page {
					continue
				}

				after, ok := rules[p]
				if !ok {
					continue
				}

				if after && i < j {
					fixed = moveToIndex(fixed, i, j+1)
					if isCorrectOrder(rMap, fixed) {
						return fixed
					}

					continue Outer
				}

				if !after && j < i {
					fixed = moveToIndex(fixed, j, i+1)
					if isCorrectOrder(rMap, fixed) {
						return fixed
					}

					continue Outer
				}
			}
		}
	}
}

func moveToIndex(update []string, from, to int) []string {
	result := make([]string, 0, len(update))
	copy(result, update)

	if from < to {
		result = append(result, update[:from]...)
		result = append(result, update[from+1:to]...)
		result = append(result, update[from])
		if to <= len(update)-1 {
			result = append(result, update[to:]...)
		}
		return result
	} else {
		result := update[:to]
		result = append(result, update[from])
		result = append(result, update[to+1:from]...)
		if from <= len(update)-1 {
			result = append(result, update[from:]...)
		}
		return result
	}
}

func main() {
	input := helpers.ReadFile("2024/day05/input.txt")

	rMap, updates := parseInput(input)
	fmt.Printf("part1: %v\n", part1(rMap, updates))
	fmt.Printf("part2: %v\n", part2(rMap, updates))
}
