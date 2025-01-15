package utils

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"

	"lem-in/vars"
)

// ValidColonyRooms validates room and connection data in the provided file.
// It checks for valid room connections, ensures no duplicate room names or coordinates, and stores valid rooms.
// After processing, it compares the number of connected rooms to the total room names, sorting and comparing them.
// The function returns false if any validation fails, and true if all checks pass.
// It uses file.Seek(0, 0) to reset the file pointer to the beginning for potential further operations.
func ValidColonyRooms(file *os.File) bool {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if ValidRoomConnection(line) {
			StoreConnectedRooms(line)
		} else if strings.Contains(line, "-") && !ValidRoomConnection(line) {
			log.Fatalf("ERROR: invalid data format, invalid room link %s", line)
		}

		if strings.Contains(line, " ") {
			name, x, y, errRoom := GetRoom(line)
			HandleError(errRoom)

			if SliceContainsString(vars.RoomNames, name) {
				log.Fatal("ERROR: invalid data format, room definition repeated")
			}

			if !ValidCoordinates(x, y) {
				log.Fatalf("ERROR: invalid data format, duplicated coordinates on %s", line)
			}

			StoreRoom(name, x, y)
			if !SliceContainsString(vars.RoomNames, name) {
				vars.RoomNames = append(vars.RoomNames, name)
			}
		}
	}

	if len(vars.ConnectedRooms) != len(vars.RoomNames) {
		return false
	}

	sort.Strings(vars.RoomNames)
	sort.Strings(vars.ConnectedRooms)

	for i := range vars.ConnectedRooms {
		if vars.ConnectedRooms[i] != vars.RoomNames[i] {
			return false
		}
	}
	file.Seek(0, 0)
	return true
}
