package utils

import "lem-in/models"

// MaxTurns calculates the nuber of turn in each selected paths
func MaxTurns(paths []models.Path) int {
	maxTurns := 1
	for _, path := range paths {
		rooms := path.Rooms[1 : len(path.Rooms)-1]
		ants := path.Ants
		turns := len(rooms) + len(ants)

		if turns > maxTurns {
			maxTurns = turns
		}
	}
	return maxTurns
}
