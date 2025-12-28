package basicgameplaytasks

import (
	"fmt"
	. "strategy-test-back/src/core/GameInstance"
)

type BasicEffectTask struct {
	GameplayTask

	InstigatorCharacter *Character
	TargetCharacter     *Character

	EffectType EffectType

	shouldEnd   chan struct{}
	effectValue *GameStateEffect
}

type BasicEffectNotification struct {
	EffectType EffectType
	Remove     bool
	Target     *Character
}

func NewBasicEffectTask(
	InstigatorCharacter *Character,
	TargetCharacter *Character,
	EffectType EffectType,
	shouldEnd chan struct{},
) *BasicEffectTask {
	return &BasicEffectTask{
		InstigatorCharacter: InstigatorCharacter,
		TargetCharacter:     TargetCharacter,
		EffectType:          EffectType,
		shouldEnd:           shouldEnd,
	}
}

func (t *BasicEffectTask) OnAdd(tickContext TickContext, gi *GameInstance) {
	fmt.Println("ADD: basic effect task")

	var effectValue = &GameStateEffect{
		Type:  t.EffectType,
		Value: t.TargetCharacter,
	}

	t.effectValue = effectValue

	gi.AddPersisentEffect(effectValue)

	gi.EventNotifier.Notify(&EventNotification{
		Type: "effect",
		Tick: tickContext.CurrentTick,
		Payload: BasicEffectNotification{
			EffectType: t.EffectType,
			Target:     t.TargetCharacter,
			Remove:     false,
		},
	})

	go func() {
		<-t.shouldEnd

		gi.EventNotifier.Notify(&EventNotification{
			Type: "effect",
			Tick: tickContext.CurrentTick,
			Payload: BasicEffectNotification{
				EffectType: t.EffectType,
				Target:     t.TargetCharacter,
				Remove:     true,
			},
		})

		gi.RemoveGameplayTask(t)
		gi.RemovePersisentEffect(t.effectValue)
	}()
}
