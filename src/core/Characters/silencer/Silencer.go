package silencer

import gameinstance "strategy-test-back/src/core/GameInstance"

func GetDefaults() gameinstance.CharacterDescription {
	return gameinstance.CharacterDescription{
		CharacterID: "silencer",

		Name:        "Silencer",
		Description: "",

		MoveSpeed: 3,

		AttackType:      gameinstance.Range,
		AttackRange:     10.0,
		AttackSpeed:     0.3,
		ProjectileSpeed: 10,

		InitialStats: gameinstance.Stats{
			Health: 100,
			Mana:   100,

			Armor:        5,
			Strength:     10,
			Intelligence: 15,
			Agility:      10,
		},
	}
}
