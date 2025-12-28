package gameinstance

import vector "strategy-test-back/src/core/Vector"

type AttackType int

const (
	Melee AttackType = iota
	Range
)

type CharacterDescription struct {
	CharacterID string `json:"characterId"`

	Name        string `json:"name"`
	Description string `json:"description"`

	MoveSpeed float64 `json:"moveSpeed"`

	AttackType      AttackType `json:"attackType"`
	AttackRange     float64    `json:"attackRange"`
	AttackSpeed     float64    `json:"attackSpeed"`
	ProjectileSpeed float64    `json:"projectileSpeed"`

	InitialStats Stats `json:"initialStats"`
}

type Character struct {
	Actor
	CharacterDescription

	Position vector.Vector2D `json:"position"`
	Stats    Stats           `json:"stats"`
}

type ICharacter interface {
}
