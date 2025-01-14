package ants

import (
	"testing"

	"lem-in/models"
)

func TestAssignAnts(t *testing.T) {
	paths := []models.Path{
		{
			Rooms: []string{"r1", "r2", "r10"},
			Ants:  []int{},
		},
		{
			Rooms: []string{"r1", "r3", "r5", "r7", "r8", "r10"},
			Ants:  []int{},
		},
		{
			Rooms: []string{"r1", "r4", "r6", "r10"},
			Ants:  []int{},
		},
		{
			Rooms: []string{"r1", "r9", "r10"},
			Ants:  []int{},
		},
	}

	expected := []int{8, 4, 4, 4}

	numAnts := 20

	AssignAnts(paths, numAnts) // populates the ants in place
	// fmt.Println(paths)

	for index, path := range paths {
		got := len(path.Ants) // number of ants assigned to current path

		expect := expected[index] // expected number of ants in current path
		if expect != got {
			t.Errorf("Expected path at index %d to have %d ants assigned, but got %d instead\n", index, expect, got)
		}
	}
}
