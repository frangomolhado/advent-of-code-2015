package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leakedmemory/aoc/helpers"
)

func parseInput(input string) (map[int][]int, [][]int) {
	sections := strings.Split(input, "\n\n")
	rules := strings.Split(sections[0], "\n")
	updates := strings.Split(sections[1], "\n")

	// parse rules into a directed graph
	graph := make(map[int][]int)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		graph[x] = append(graph[x], y)
	}

	// parse updates into int slices
	updateList := [][]int{}
	for _, update := range updates {
		pages := strings.Split(update, ",")
		parsed := []int{}
		for _, p := range pages {
			page, _ := strconv.Atoi(p)
			parsed = append(parsed, page)
		}
		updateList = append(updateList, parsed)
	}

	return graph, updateList
}

func part1(graph map[int][]int, updates [][]int) int {
	result := 0
	for _, update := range updates {
		if isCorrectOrder(graph, update) {
			result += update[len(update)/2]
		}
	}

	return result
}

func isCorrectOrder(graph map[int][]int, update []int) bool {
	pos := make(map[int]int)
	for i, page := range update {
		pos[page] = i
	}

	for x, ys := range graph {
		if xIdx, ok := pos[x]; ok {
			for _, y := range ys {
				if yIdx, ok := pos[y]; ok && xIdx > yIdx {
					return false
				}
			}
		}
	}

	return true
}

func part2(graph map[int][]int, updates [][]int) int {
	result := 0
	for _, update := range updates {
		if !isCorrectOrder(graph, update) {
			sorted := topologicalSort(graph, update)
			result += sorted[len(sorted)/2]
		}
	}

	return result
}

func topologicalSort(graph map[int][]int, update []int) []int {
	pos := make(map[int]int)
	for i, page := range update {
		pos[page] = i
	}

	inDegree := make(map[int]int)
	for x, ys := range graph {
		if _, ok := pos[x]; ok {
			for _, y := range ys {
				if _, ok := pos[y]; ok {
					inDegree[y]++
				}
			}
		}
	}

	// initialize queue with nodes that have no dependencies
	queue := []int{}
	for _, page := range update {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	sorted := []int{}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		minPos := len(update)
		minIdx := 0
		for i, node := range queue {
			if pos[node] < minPos {
				minPos = pos[node]
				minIdx = i
			}
		}

		node := queue[minIdx]
		queue = append(queue[:minIdx], queue[minIdx+1:]...)

		if !visited[node] {
			sorted = append(sorted, node)
			visited[node] = true

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					inDegree[neighbor]--
					if inDegree[neighbor] == 0 {
						queue = append(queue, neighbor)
					}
				}
			}
		}
	}

	return sorted
}

func main() {
	input := helpers.ReadFile("2024/day05/input.txt")

	graph, updates := parseInput(input)
	fmt.Printf("part1: %v\n", part1(graph, updates))
	fmt.Printf("part2: %v\n", part2(graph, updates))
}
