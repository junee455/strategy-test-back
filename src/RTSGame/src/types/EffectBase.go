package RTSGame

type EffectID string

type CharacterEffect struct {
	ID         EffectID
	Instigator CharacterID
	Priority   int32
}

type ICharacterEffect interface {
}
