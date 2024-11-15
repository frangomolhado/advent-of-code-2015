package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leakedmemory/aoc/helpers"
)

func part1(input string) int {
	lights := [1000][1000]int{}
	for _, line := range strings.Split(input, "\n") {
		splitLine := strings.Split(line, " ")
		if splitLine[0] == "toggle" {
			fromX, fromY := getCoordinates(splitLine[1])
			toX, toY := getCoordinates(splitLine[3])
			for row := fromY; row <= toY; row++ {
				for column := fromX; column <= toX; column++ {
					lights[row][column] ^= 1
				}
			}
		} else if splitLine[1] == "on" {
			fromX, fromY := getCoordinates(splitLine[2])
			toX, toY := getCoordinates(splitLine[4])
			for row := fromY; row <= toY; row++ {
				for column := fromX; column <= toX; column++ {
					lights[row][column] = 1
				}
			}
		} else {
			fromX, fromY := getCoordinates(splitLine[2])
			toX, toY := getCoordinates(splitLine[4])
			for row := fromY; row <= toY; row++ {
				for column := fromX; column <= toX; column++ {
					lights[row][column] = 0
				}
			}
		}
	}

	result := 0
	for i := range 1000 {
		for j := range 1000 {
			result += lights[i][j]
		}
	}
	return result
}

func part2(input string) int {
	lights := [1000][1000]int{}
	for _, line := range strings.Split(input, "\n") {
		splitLine := strings.Split(line, " ")
		if splitLine[0] == "toggle" {
			fromX, fromY := getCoordinates(splitLine[1])
			toX, toY := getCoordinates(splitLine[3])
			for row := fromY; row <= toY; row++ {
				for column := fromX; column <= toX; column++ {
					lights[row][column] += 2
				}
			}
		} else if splitLine[1] == "on" {
			fromX, fromY := getCoordinates(splitLine[2])
			toX, toY := getCoordinates(splitLine[4])
			for row := fromY; row <= toY; row++ {
				for column := fromX; column <= toX; column++ {
					lights[row][column]++
				}
			}
		} else {
			fromX, fromY := getCoordinates(splitLine[2])
			toX, toY := getCoordinates(splitLine[4])
			for row := fromY; row <= toY; row++ {
				for column := fromX; column <= toX; column++ {
					lights[row][column] = max(0, lights[row][column]-1)
				}
			}
		}
	}

	result := 0
	for i := range 1000 {
		for j := range 1000 {
			result += lights[i][j]
		}
	}
	return result
}

func getCoordinates(s string) (x, y int) {
	coordinates := strings.Split(s, ",")
	x, _ = strconv.Atoi(coordinates[0])
	y, _ = strconv.Atoi(coordinates[1])
	return
}

func main() {
	input := helpers.ReadFile("2015/day06/input.txt")

	fmt.Printf("part1: %d\n", part1(input))
	fmt.Printf("part2: %d\n", part2(input))
}
