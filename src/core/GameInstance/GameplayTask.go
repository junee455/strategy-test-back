package gameinstance

type GameplayTaskType string

type GameplayTask struct {
	ID           GameplayTaskType
	Priority     int
	shouldRemove bool
	Ticked       bool
}

type GameStateEffect struct {
	Type  EffectType
	Value any
}

type GameplayTaskDescription struct {
	ID       GameplayTaskType
	Priority int
	Metadata any
}

type IGameplayTask interface {
	OnAdd(tickContext TickContext, gi *GameInstance)
	OnFirstTick(tickContext TickContext, gi *GameInstance)
	OnTasksTick(tickContext TickContext, gi *GameInstance)
	GetDescription() GameplayTaskDescription
	SetShouldRemove()
	GetShouldRemove() bool
}

func (gt *GameplayTask) GetDescription() GameplayTaskDescription {
	return GameplayTaskDescription{
		ID:       gt.ID,
		Priority: gt.Priority,
	}
}

func (gt *GameplayTask) OnFirstTick(tickContext TickContext, gi *GameInstance) {}

func (gt *GameplayTask) OnAdd(tickContext TickContext, gi *GameInstance) {}

func (gt *GameplayTask) OnTasksTick(tickContext TickContext, gi *GameInstance) {}

func (gt *GameplayTask) SetShouldRemove() {
	gt.shouldRemove = true
}

func (gt *GameplayTask) GetShouldRemove() bool {
	return gt.shouldRemove
}
