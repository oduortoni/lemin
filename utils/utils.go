package utils

import (
	"log"
	"strings"

	"lem-in/vars"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

// SliceInSlices checks if any element from the second slice (arr2) is present in any of the sub-slices of the first slice (arr1).
// The function iterates over each element of arr2 and, for each element, checks if it exists in any sub-slice of arr1
// using the helper function SliceContainsString. If a match is found, the function returns true.
// If no elements from arr2 are found in any sub-slice of arr1, it returns false.
func SliceInSlices(arr1 [][]string, arr2 []string) bool {
	for _, v := range arr2 {
		for _, w := range arr1 {
			if SliceContainsString(w, v) {
				return true
			}
		}
	}
	return false
}

// ProcessLine processes a single line of input, updating global variables based on its content.
// The function performs the following steps:
// - If the line represents a valid room connection, it updates the colony's room connectivity (graph).
// - The graph is in form of a map where the keys are the nodes, and values, all its neighbors.
// Each step also includes error handling for potential issues encountered during processing.
func ProcessLine(line string) {
	if vars.FirstLine {
		errNumberOfAnts := ProcessNumberOfAnts(line)
		HandleError(errNumberOfAnts)
		vars.FirstLine = false
	} else if vars.IsStartNode {
		start, _, _, errRoom := GetRoom(line)
		HandleError(errRoom)
		vars.StartRoom = start
		vars.IsStartNode = false
	} else if vars.IsEndNode {
		end, _, _, errRoom := GetRoom(line)
		HandleError(errRoom)
		vars.EndRoom = end
		vars.IsEndNode = false
	} else if ValidRoomConnection(line) {
		rooms := strings.Split(line, "-")
		vars.Colony[rooms[0]] = append(vars.Colony[rooms[0]], rooms[1])
		vars.Colony[rooms[1]] = append(vars.Colony[rooms[1]], rooms[0])
	}
}
