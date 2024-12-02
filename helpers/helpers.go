package helpers

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadFile returns the content of the file as a string, trimming the line feed
// at the end and panicking on error.
func ReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	return strings.TrimRight(string(content), "\n")
}

// ReadLines returns the content of the file as a string, trimming the line feed
// at the end, splitting using the line feed character as separator, and
// panicking on error.
func ReadLines(path string) []string {
	return strings.Split(ReadFile(path), "\n")
}

// SliceAtoi applies strconv.Atoi to all the values of a slice and returns a
// new one with the values as integers. Panics on conversion errors.
func SliceAtoi(slice []string) []int {
	result := make([]int, 0, len(slice))
	for _, s := range slice {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}

		result = append(result, n)
	}

	return result
}
