package utils

import (
	"strings"

	"lem-in/vars"
)

// StoreConnectedRooms processes a line representing a connection between two rooms,
// and stores the unique room names in the global vars.ConnectedRooms slice.
// The function splits the line by a hyphen ("-"), iterates over the resulting room names,
// and appends each room to vars.ConnectedRooms only if it is not already present in the slice.
func StoreConnectedRooms(line string) {
	rooms := strings.Split(line, "-")
	for _, v := range rooms {
		if !SliceContainsString(vars.ConnectedRooms, v) {
			vars.ConnectedRooms = append(vars.ConnectedRooms, v)
		}
	}
}
