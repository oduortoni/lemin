package utils

import (
	"bufio"
	"os"
	"strings"
)

// HasStartAndEnd check if the colony has a start and end room,
// It returns true if it does and false otherwise.
func HasStartAndEnd(file *os.File) bool {
	hasStart := false
	hasEnd := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "##start" {
			hasStart = true
		} else if line == "##end" {
			hasEnd = true
		}
		if hasStart && hasEnd {
			file.Seek(0, 0)
			return true
		}
	}
	return false
}
