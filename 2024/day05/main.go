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

func day05(graph map[int][]int, updates [][]int) (int, int) {
	part1 := 0
	part2 := 0
	for _, update := range updates {
		if isCorrectOrder(graph, update) {
			part1 += update[len(update)/2]
		} else {
			sorted := topologicalSort(graph, update)
			part2 += sorted[len(sorted)/2]
		}
	}

	return part1, part2
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

func topologicalSort(graph map[int][]int, update []int) []int {
	pagesPos := make(map[int]int)
	for i, page := range update {
		pagesPos[page] = i
	}

	inDegree := make(map[int]int)
	for x, ys := range graph {
		if _, ok := pagesPos[x]; ok {
			for _, y := range ys {
				if _, ok := pagesPos[y]; ok {
					inDegree[y]++
				}
			}
		}
	}

	queue := []int{}
	for _, page := range update {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	sorted := []int{}
	for len(queue) > 0 {
		minPos := len(update)
		minIdx := 0
		for i, node := range queue {
			if pagesPos[node] < minPos {
				minPos = pagesPos[node]
				minIdx = i
			}
		}

		node := queue[minIdx]
		queue = append(queue[:minIdx], queue[minIdx+1:]...)
		sorted = append(sorted, node)
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func main() {
	input := helpers.ReadFile("2024/day05/input.txt")
	graph, updates := parseInput(input)
	p1, p2 := day05(graph, updates)
	fmt.Printf("part1: %v\npart2: %v\n", p1, p2)
}
