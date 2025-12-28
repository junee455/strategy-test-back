package gameinstance

import "strconv"

type ActorID string

var lastId = 0

type Actor struct {
	RuntimeId ActorID `json:"runtimeId"`
}

type IActor interface{}

func NewActor() *Actor {
	lastId++

	return &Actor{
		RuntimeId: ActorID(strconv.Itoa(lastId)),
	}
}
