package main

import (
	"github.com/lithammer/shortuuid/v3"
)

var Rooms = make(map[string]*Room)

func createUser(name, offer string) *User {
	u := &User{
		Name:  name,
		Offer: offer,
	}
	return u
}

func createRoom() *Room {
	room := &Room{
		ID: shortuuid.New(),
	}
	Rooms[room.ID] = room
	return room
}
