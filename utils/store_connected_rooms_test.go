package utils

import (
	"testing"

	"lem-in/vars"
)

func TestStoreConnectedRooms(t *testing.T) {
	restore := vars.ConnectedRooms
	vars.ConnectedRooms = []string{} // set it to empty just in case it was set by other tests

	table := []struct {
		Connections    []string
		Expect         bool
		NumRoomsInPath int
	}{
		{
			Connections:    []string{"1-2", "2-3", "3-4"},
			Expect:         true,
			NumRoomsInPath: 4,
		},
		{
			Connections:    []string{"1-2", "2-3", "3-4", "1-7", "4-1", "5-3"}, // added two new rooms. Even though they are connected to already existing rooms, they themselves have to be added in the path as new rooms, hence sum goes up
			Expect:         true,
			NumRoomsInPath: 6,
		},
	}

	for _, entry := range table {
		for _, connects := range entry.Connections {
			StoreConnectedRooms(connects)
		}
		n := len(vars.ConnectedRooms)
		if n != entry.NumRoomsInPath {
			t.Errorf("Expected total number of rooms in path to be %d but got %d\n", entry.NumRoomsInPath, n)
			vars.ConnectedRooms = restore
		}
	}

	vars.ConnectedRooms = restore
}
