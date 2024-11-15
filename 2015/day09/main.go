package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/leakedmemory/aoc/helpers"
)

type edge struct {
	to     string
	weight int
}

func createGraph(lines []string) map[string][]edge {
	graph := map[string][]edge{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		from := split[0]
		to := split[2]
		weight, _ := strconv.Atoi(split[len(split)-1])
		graph[from] = append(graph[from], edge{to, weight})
		graph[to] = append(graph[to], edge{from, weight})
	}

	return graph
}

func part1(lines []string) int {
	graph := createGraph(lines)
	result := math.MaxInt
	for startNode := range graph {
		result = min(result, tsp(graph, startNode))
	}

	return result
}

func tsp(graph map[string][]edge, start string) int {
	nodes := make([]string, 0, len(graph))
	nodeIndex := map[string]int{}
	for node := range graph {
		nodeIndex[node] = len(nodes)
		nodes = append(nodes, node)
	}

	n := len(nodes)

	// adjacency matrix representation
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, n)
		for j := range adj[i] {
			adj[i][j] = math.MaxInt
		}
	}

	for from, edges := range graph {
		fromIndex := nodeIndex[from]
		for _, e := range edges {
			toIndex := nodeIndex[e.to]
			adj[fromIndex][toIndex] = e.weight
		}
	}

	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}

	parent := make([][]int, 1<<n)
	for i := range parent {
		parent[i] = make([]int, n)
		for j := range parent[i] {
			parent[i][j] = -1
		}
	}

	startIndex := nodeIndex[start]
	// base case
	dp[1<<startIndex][startIndex] = 0
	// solve using bitmask
	for mask := range 1 << n {
		for u := range n {
			if dp[mask][u] == math.MaxInt {
				continue
			}

			for v := range n {
				if mask&(1<<v) == 0 && adj[u][v] != math.MaxInt {
					newMask := mask | (1 << v)
					newCost := dp[mask][u] + adj[u][v]
					if newCost < dp[newMask][v] {
						dp[newMask][v] = newCost
						parent[newMask][v] = u
					}
				}
			}
		}
	}

	// find shortest path
	endMask := (1 << n) - 1
	shortestCost := math.MaxInt
	for i := range n {
		if dp[endMask][i] < shortestCost {
			shortestCost = dp[endMask][i]
		}
	}

	return shortestCost
}

func part2(lines []string) int {
	graph := createGraph(lines)
	result := math.MinInt
	for startNode := range graph {
		result = max(result, tspLongest(graph, startNode))
	}

	return result
}

func tspLongest(graph map[string][]edge, start string) int {
	nodes := make([]string, 0, len(graph))
	nodeIndex := map[string]int{}
	for node := range graph {
		nodeIndex[node] = len(nodes)
		nodes = append(nodes, node)
	}

	n := len(nodes)

	// adjacency matrix representation
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, n)
		for j := range adj[i] {
			adj[i][j] = math.MinInt
		}
	}

	for from, edges := range graph {
		fromIndex := nodeIndex[from]
		for _, e := range edges {
			toIndex := nodeIndex[e.to]
			adj[fromIndex][toIndex] = e.weight
		}
	}

	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MinInt
		}
	}

	startIndex := nodeIndex[start]
	// base case
	dp[1<<startIndex][startIndex] = 0
	// solve using bitmask
	for mask := range 1 << n {
		for u := range n {
			if dp[mask][u] == math.MinInt {
				continue
			}

			for v := range n {
				if mask&(1<<v) == 0 && adj[u][v] != math.MinInt {
					newMask := mask | (1 << v)
					newCost := dp[mask][u] + adj[u][v]
					if newCost > dp[newMask][v] {
						dp[newMask][v] = newCost
					}
				}
			}
		}
	}

	// find longest path
	endMask := (1 << n) - 1
	longestCost := math.MinInt
	for i := range n {
		if dp[endMask][i] > longestCost {
			longestCost = dp[endMask][i]
		}
	}

	return longestCost
}

func main() {
	lines := helpers.ReadLines("2015/day09/input.txt")

	fmt.Printf("part1: %d\n", part1(lines))
	fmt.Printf("part2: %d\n", part2(lines))
}
