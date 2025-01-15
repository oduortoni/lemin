package utils

import (
	"testing"

	"lem-in/models"
	"lem-in/vars"
)

func TestValidCoordinates(t *testing.T) {
	table := []struct {
		Rooms    []models.Room
		Input    Coord
		Expected bool
	}{
		{
			Rooms:    []models.Room{{X: 2, Y: 3}, {X: 3, Y: 5}, {X: 2, Y: 3}, {X: 9, Y: 0}},
			Input:    Coord{34, 89},
			Expected: true,
		},
		{
			Rooms:    []models.Room{{X: 2, Y: 3}, {X: 3, Y: 5}, {X: 2, Y: 3}, {X: 9, Y: 0}},
			Input:    Coord{9, 0},
			Expected: false,
		},
	}

	restoreRoooms := vars.Rooms
	for _, entry := range table {
		vars.Rooms = entry.Rooms
		valid := ValidCoordinates(entry.Input.X, entry.Input.Y)
		if valid != entry.Expected {
			vars.Rooms = restoreRoooms
			t.Errorf("no two oordinates should overlap in a map")
		}
	}
	vars.Rooms = restoreRoooms
}
