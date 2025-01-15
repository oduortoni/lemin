package utils

import (
	"fmt"
	"strconv"

	"lem-in/vars"
)

// ProcessNumberOfAnts parses a line representing the number of ants and updates the global ants count.
func ProcessNumberOfAnts(line string) error {
	number, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("invalid number of ants: %v", err)
	}
	vars.AntsNumber = number
	return nil
}
