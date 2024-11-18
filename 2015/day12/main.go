package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/leakedmemory/aoc/helpers"
)

func part1(input string) int {
	result := 0
	var sb strings.Builder
	for _, c := range input {
		if unicode.IsNumber(c) || c == '-' {
			sb.WriteRune(c)
		} else if sb.Len() > 0 {
			number, _ := strconv.Atoi(sb.String())
			result += number
			sb.Reset()
		}
	}

	return result
}

func part2(input string) int {
	var parsed any
	_ = json.Unmarshal([]byte(input), &parsed)
	return sumJSON(parsed)
}

func sumJSON(data any) int {
	switch v := data.(type) {
	case float64: // JSON numbers are parsed as float64
		return int(v)
	case []any:
		sum := 0
		for _, elem := range v {
			sum += sumJSON(elem)
		}
		return sum
	case map[string]any:
		for _, val := range v {
			if str, ok := val.(string); ok && str == "red" {
				return 0
			}
		}

		sum := 0
		for _, val := range v {
			sum += sumJSON(val)
		}
		return sum
	default:
		return 0
	}
}

func main() {
	input := helpers.ReadFile("2015/day12/input.txt")

	fmt.Printf("part1: %v\n", part1(input))
	fmt.Printf("part2: %v\n", part2(input))
}
