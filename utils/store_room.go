package utils

import (
	"lem-in/models"
	"lem-in/vars"
)

// StoreRoom strores and struct room with its name and x and y coordinates
func StoreRoom(name string, x, y int) {
	room := models.Room{
		Name: name,
		X:    x,
		Y:    y,
	}
	vars.Rooms = append(vars.Rooms, room)
}
