package RTSGame

type CharacterID string

type CharacterInstance struct {
	ID       CharacterID
	Position [2]float32
	Health   float32
	Mana     float32

	Effects []CharacterEffect
}

type ICharacterInstance interface {
}
