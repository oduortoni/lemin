package utils

import (
	"testing"

	"lem-in/models"
)

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
