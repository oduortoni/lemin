package utils

import "strings"

// ValidRoomConnection checks whether a given line represents a valid room connection.
// The function splits the line by a hyphen ("-") and validates the following conditions:
// - The line contains exactly one hyphen separating two rooms.
// - Both rooms are non-empty and do not contain spaces.
// - The two rooms are not the same.
func ValidRoomConnection(line string) bool {
	rooms := strings.Split(line, "-")
	return len(rooms) == 2 &&
		strings.Contains(line, "-") &&
		rooms[0] != "" &&
		rooms[1] != "" &&
		!strings.Contains(rooms[0], " ") &&
		!strings.Contains(rooms[1], " ") &&
		rooms[0] != rooms[1]
}
