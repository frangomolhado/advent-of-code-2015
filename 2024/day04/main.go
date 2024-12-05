package main

import (
	"fmt"

	"github.com/leakedmemory/aoc/helpers"
)

func part1(lines []string) int {
	result := 0
	for row, line := range lines {
		for column, c := range line {
			//nolint:nestif
			if c == 'X' {
				if isXMASRow(line, column) {
					result++
				}

				if isXMASColumn(lines[row+1:min(row+4, len(lines))], column) {
					result++
				}

				if isXMASRightDiagonal(lines[row+1:min(row+4, len(lines))], column+1) {
					result++
				}

				if isXMASLeftDiagonal(lines[row+1:min(row+4, len(lines))], column-1) {
					result++
				}
			} else if c == 'S' {
				if isSAMXRow(line, column) {
					result++
				}

				if isSAMXColumn(lines[row+1:min(row+4, len(lines))], column) {
					result++
				}

				if isSAMXRightDiagonal(lines[row+1:min(row+4, len(lines))], column+1) {
					result++
				}

				if isSAMXLeftDiagonal(lines[row+1:min(row+4, len(lines))], column-1) {
					result++
				}
			}
		}
	}

	return result
}

func isXMASRow(line string, column int) bool {
	return len(line)-column > 3 && line[column:column+4] == "XMAS"
}

func isXMASColumn(lines []string, column int) bool {
	return len(lines) == 3 &&
		lines[0][column] == 'M' &&
		lines[1][column] == 'A' &&
		lines[2][column] == 'S'
}

func isXMASRightDiagonal(lines []string, column int) bool {
	return len(lines) == 3 &&
		len(lines[0])-column >= 3 &&
		lines[0][column] == 'M' &&
		lines[1][column+1] == 'A' &&
		lines[2][column+2] == 'S'
}

func isXMASLeftDiagonal(lines []string, column int) bool {
	return len(lines) == 3 && column >= 2 &&
		lines[0][column] == 'M' &&
		lines[1][column-1] == 'A' &&
		lines[2][column-2] == 'S'
}

func isSAMXRow(line string, column int) bool {
	return len(line)-column > 3 && line[column:column+4] == "SAMX"
}

func isSAMXColumn(lines []string, column int) bool {
	return len(lines) == 3 &&
		lines[0][column] == 'A' &&
		lines[1][column] == 'M' &&
		lines[2][column] == 'X'
}

func isSAMXRightDiagonal(lines []string, column int) bool {
	return len(lines) == 3 &&
		len(lines[0])-column >= 3 &&
		lines[0][column] == 'A' &&
		lines[1][column+1] == 'M' &&
		lines[2][column+2] == 'X'
}

func isSAMXLeftDiagonal(lines []string, column int) bool {
	return len(lines) == 3 && column >= 2 &&
		lines[0][column] == 'A' &&
		lines[1][column-1] == 'M' &&
		lines[2][column-2] == 'X'
}

func part2(lines []string) int {
	result := 0
	for row, line := range lines[1 : len(lines)-1] {
		for column, c := range line[1 : len(line)-1] {
			if c == 'A' && isXMASCrux(lines[row:row+3], column+1) {
				result++
			}
		}
	}

	return result
}

func isXMASCrux(lines []string, column int) bool {
	if lines[0][column-1] != 'M' && lines[0][column-1] != 'S' {
		return false
	}

	if lines[2][column-1] != 'M' && lines[2][column-1] != 'S' {
		return false
	}

	// right diagonal
	if lines[0][column-1] == 'M' && lines[2][column+1] != 'S' ||
		lines[0][column-1] == 'S' && lines[2][column+1] != 'M' {
		return false
	}

	// left diagonal
	if lines[2][column-1] == 'M' && lines[0][column+1] != 'S' ||
		lines[2][column-1] == 'S' && lines[0][column+1] != 'M' {
		return false
	}

	return true
}

func main() {
	lines := helpers.ReadLines("2024/day04/input.txt")

	fmt.Printf("part1: %v\n", part1(lines))
	fmt.Printf("part2: %v\n", part2(lines))
}

//    Rows    5
// .....XMAS. 1
// .SAMX..... 1
// ..........
// ..........
// XMASAMX... 2
// ..........
// ..........
// ..........
// ..........
// .....XMAS. 1

//  Columns   3
// ..........
// ......S... 1
// ......A...
// ......M..X 1
// ......X..M
// .........A
// .........S 1
// .........A
// .........M
// .........X

// Diagonals  10 5 left + 5 right
// ....X.....  1 0 left + 1 right
// .....M....
// ...S..A...  2 1 left + 1 right
// ..A.A..S.X  1 1 left + 0 right
// .M...M..M.
// X.....XA..
// S.S.S.S.S.  6 3 left + 3 right
// .A.A.A.A..
// ..M.M.M.M.
// .X.X.X...X
