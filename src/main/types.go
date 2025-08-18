package main

type EffectType int

const (
	area EffectType = iota
	target
)

type Effect struct {
	effectType EffectType
	duration   float32
}

type GameItem struct {
	id          string
	name        string
	description string
}

type GameCharacter struct {
	id   string
	name string

	health int
	mana   int

	items []GameItem
}

type GameClient struct {
	id         string
	connection interface{}
}

type GameState struct{}

type GameInstance struct {
	clients   []GameClient
	gameState GameState
	id        string
}
