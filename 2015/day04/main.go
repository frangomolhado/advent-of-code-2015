package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

const input = "iwrupvqb"

func part1() int {
	for i := 1; ; i++ {
		sum := md5.Sum([]byte(input + strconv.Itoa(i)))
		if hex.EncodeToString(sum[:])[:5] == "00000" {
			return i
		}
	}
}

func part2() int {
	for i := 1; ; i++ {
		sum := md5.Sum([]byte(input + strconv.Itoa(i)))
		if hex.EncodeToString(sum[:])[:6] == "000000" {
			return i
		}
	}
}

func main() {
	fmt.Printf("part1: %d\n", part1())
	fmt.Printf("part2: %d\n", part2())
}
