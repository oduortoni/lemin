package utils

import (
	"testing"

	"lem-in/models"
	"lem-in/vars"
)

func TestStoreRoom(t *testing.T) {
	restore := vars.Rooms
	vars.Rooms = []models.Room{}

	roomName := "rAll"

	table := []struct {
		Rooms    []models.Room
		Input    Coord
		Expected int
	}{
		{
			Rooms:    []models.Room{{X: 2, Y: 3}, {X: 3, Y: 5}, {X: 2, Y: 3}, {X: 9, Y: 0}},
			Expected: 4,
		},
		{
			Rooms:    []models.Room{{X: 2, Y: 3}, {X: 3, Y: 5}},
			Expected: 2,
		},
		{
			Rooms:    []models.Room{},
			Expected: 0,
		},
	}

	for _, entry := range table {
		for _, item := range entry.Rooms {
			StoreRoom(roomName, item.X, item.Y)
		}
		n := len(vars.Rooms)
		if n != entry.Expected {
			t.Errorf("expected saved rooms to be %d but got %d\n", entry.Expected, n)
		}
		vars.Rooms = []models.Room{}
	}

	vars.Rooms = restore
}
