package utils

import (
	"os"
	"testing"

	"lem-in/models"
	"lem-in/vars"
)

func TestSliceContainsString(t *testing.T) {
	table := []struct {
		Slice  []string
		Str    string
		Expect bool
	}{
		{
			Slice:  []string{"hello", "tenesse", "charlote", "gws"},
			Str:    "charlote",
			Expect: true,
		},
		{
			Slice:  []string{"hello", "tenesse", "charlote", "gws"},
			Str:    "non-existent",
			Expect: false,
		},
		{
			Slice:  []string{},
			Str:    "anything",
			Expect: false,
		},
		{
			Slice:  []string{"hello", "tenesse", "charlote", "gws"},
			Str:    "",
			Expect: false,
		},
	}

	for _, entry := range table {
		got := SliceContainsString(entry.Slice, entry.Str)
		if got != entry.Expect {
			t.Errorf("expected %v, but got %v\n", entry.Expect, got)
		}
	}
}

func TestSliceContainsSlice(t *testing.T) {
	table := []struct {
		Haystack []string
		Kneedle  []string
		Expect   bool
	}{
		{
			Haystack: []string{"hello", "tenesse", "charlote", "gws"},
			Kneedle:  []string{"charlote"},
			Expect:   true,
		},
		{
			Haystack: []string{"hello", "tenesse", "charlote", "gws"},
			Kneedle:  []string{"non-existent"},
			Expect:   false,
		},
		{
			Haystack: []string{},
			Kneedle:  []string{"anything"},
			Expect:   false,
		},
		{
			Haystack: []string{"hello", "tenesse", "charlote", "gws"},
			Kneedle:  []string{""},
			Expect:   false,
		},
	}

	for _, entry := range table {
		got := SliceContainsSlice(entry.Haystack, entry.Kneedle)
		if got != entry.Expect {
			t.Errorf("expected %v, but got %v\n", entry.Expect, got)
		}
	}
}

type Coord struct {
	X int
	Y int
}

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

func TestValidRoomConnection(t *testing.T) {
	table := []struct {
		Input  string
		Expect bool
	}{
		{
			Input:  "4-7",
			Expect: true,
		},
		{
			Input:  "4 7",
			Expect: false,
		},
		{
			Input:  "start-7",
			Expect: true,
		},
		{
			Input:  "4-end",
			Expect: true,
		},
		{
			Input:  "4-",
			Expect: false,
		},
		{
			Input:  "jakarta-kisumu",
			Expect: true,
		},
	}

	for _, entry := range table {
		got := ValidRoomConnection(entry.Input)
		if got != entry.Expect {
			t.Errorf("All valid room connections must have a '-' between them, but got %s\n", entry.Input)
		}
	}
}

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

func TestProcessNumberOfAnts(t *testing.T) {
	restore := vars.AntsNumber

	table := []struct {
		Line    string
		Expect  func(error) bool
		Message string
	}{
		{
			Line: "90",
			Expect: func(err error) bool {
				return err == nil
			},
			Message: "Expected a nil error",
		},
		{
			Line: "Hello",
			Expect: func(err error) bool {
				return err != nil
			},
			Message: "Expected a non-nil error",
		},
	}

	for _, entry := range table {
		err := ProcessNumberOfAnts(entry.Line)
		if !entry.Expect(err) {
			t.Errorf("%s\n", entry.Message)
			vars.AntsNumber = restore
		}
	}
	vars.AntsNumber = restore
}

func TestGetRoom(t *testing.T) {
	table := []struct {
		Line    string
		Expect  func(int, int, error) bool
		Message string
	}{
		{
			Line: "90 2 27",
			Expect: func(x, y int, err error) bool {
				return err == nil && x == 2 && y == 27
			},
			Message: "Expected a nil error",
		},
		{
			Line: "Hello",
			Expect: func(x, y int, err error) bool {
				return err != nil
			},
			Message: "Expected a non-nil error",
		},
	}

	for _, entry := range table {
		_, x, y, err := GetRoom(entry.Line)
		if !entry.Expect(x, y, err) {
			t.Errorf("%s\n", entry.Message)
		}
	}
}

func TestStoreRoom(t *testing.T) {
	//
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

func TestValidColonyRooms(t *testing.T) {
	file, err := os.CreateTemp("", "temp_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name()) // Clean up after test

	_, err = file.WriteString("4\n##start\n0 0 3\n2 2 5\n3 4 0\n##end\n1 8 3\n0-2\n2-3\n3-1\n")
	if err != nil {
		t.Fatalf("Failed to write to file: %v", err)
	}

	// open/close to test if file can handle the open/read cycles
	err = file.Close() // Close the file after writing
	if err != nil {
		t.Fatalf("Failed to close file: %v", err)
	}

	file, err = os.Open(file.Name()) // reopen the file for reading
	if err != nil {
		t.Fatalf("Failed to open file for testing: %v", err)
	}
	defer file.Close()

	isValid := ValidColonyRooms(file)
	if !isValid {
		t.Errorf("Expected the file to be valid, but got false")
	}
}

func TestHasStartAndEnd(t *testing.T) {
	file, err := os.CreateTemp("", "temp_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name()) // Clean up after test

	_, err = file.WriteString("4\n##start\n0 0 3\n2 2 5\n3 4 0\n##end\n1 8 3\n0-2\n2-3\n3-1\n")
	if err != nil {
		t.Fatalf("Failed to write to file: %v", err)
	}

	// open/close to test if file can handle the open/read cycles
	err = file.Close() // Close the file after writing
	if err != nil {
		t.Fatalf("Failed to close file: %v", err)
	}

	file, err = os.Open(file.Name()) // reopen the file for reading
	if err != nil {
		t.Fatalf("Failed to open file for testing: %v", err)
	}
	defer file.Close()

	isValid := HasStartAndEnd(file)
	if !isValid {
		t.Errorf("Expected the file to be valid, but got false")
	}
}

func TestMaxTurns(t *testing.T) {
	table := []struct {
		Paths    []models.Path
		Expected int
	}{
		{
			Paths: []models.Path{
				{Rooms: []string{"start", "2", "3", "end"}, Ants: []int{1, 2, 3}}, // 5
				{Rooms: []string{"start", "2", "end"}, Ants: []int{23, 4}},        // 3
			},
			Expected: 5,
		},
		{
			Paths: []models.Path{
				{Rooms: []string{"start", "end"}, Ants: []int{1, 2, 3}},
				{Rooms: []string{"start", "2", "end"}, Ants: []int{23, 4}},
			},
			Expected: 3,
		},
	}

	for _, entry := range table {
		maxTurns := MaxTurns(entry.Paths)
		if maxTurns != entry.Expected {
			t.Errorf("Expected %d max turns but got %d\n", entry.Expected, maxTurns)
		}
	}
}
