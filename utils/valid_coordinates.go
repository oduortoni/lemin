package utils

import "lem-in/vars"

// ValidCoordinates checks if the given coordinates (x, y) are valid by ensuring that no existing room
// in the global vars.Rooms slice has the same coordinates.
func ValidCoordinates(x, y int) bool {
	for _, v := range vars.Rooms {
		if v.X == x && v.Y == y {
			return false
		}
	}
	return true
}
