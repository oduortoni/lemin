package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// GetRoom parses a line of input representing a room's details and returns the room's name and coordinates (x, y).
// The function expects the line to contain exactly three space-separated values:
// - The room name (string)
// - The x-coordinate (integer)
// - The y-coordinate (integer)
func GetRoom(line string) (string, int, int, error) {
	room := strings.Split(line, " ")
	if len(room) != 3 {
		return "", 0, 0, fmt.Errorf("invalid room details, %s", line)
	}
	name := room[0]

	x, err := strconv.Atoi(room[1])
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid x coordinate: %v", err)
	}

	y, err := strconv.Atoi(room[2])
	if err != nil {
		return "", 0, 0, fmt.Errorf("invalid y coordinate: %v", err)
	}

	return name, x, y, nil
}
