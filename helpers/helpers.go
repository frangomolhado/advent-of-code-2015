package helpers

import (
	"log"
	"os"
	"strings"
)

// Readfile returns the content of the file as a string, trimming the line feed
// at the end and panicking on error.
func ReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	return strings.TrimRight(string(content), "\n")
}
