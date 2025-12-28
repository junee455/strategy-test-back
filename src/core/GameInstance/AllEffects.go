package gameinstance

import (
	vector "strategy-test-back/src/core/Vector"
)

type EffectType int64

const (
	MoveEffect EffectType = iota
	TeleportEffect
	MoveProjectileEffect
	ApplyDamageEffect
	RestoreHealthEffect
	RestoreManaEffect
	Dispel
	Spawn
	// continuous effects
	Bash
	Hex
	Invisibility
	Silence
	ImmunityMagic
	ImmunityPhysical
	ImmunityEffects
	Disarm
	Ghost
)

// ~~~~~~~~~~~~~

type MoveEffectPayload struct {
	TargetCharacter *Character
	Dv              vector.Vector2D
}

type TeleportEffectPayload struct {
	TargetCharacter *Character
	NewPosition     vector.Vector2D
}

type MoveProjectileEffectPayload struct {
	Dv vector.Vector2D
}

type DamageType int

const (
	Physical DamageType = iota
	Magic
	Pure
)

type ApplyDamageEffectPayload struct {
	InstigatorCharacter *Character
	TargetCharacter     *Character
	DamageAmount        int
	DamageType          DamageType
}

type RestoreHealthEffectPayload struct{}

type RestoreManaEffectPayload struct{}

type DispelPayload struct{}

type SpawnEffectPayload struct {
	CharacterDescription CharacterDescription
	Position             vector.Vector2D
	Stats                Stats
}
