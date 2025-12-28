package basicgameplaytasks

import (
	. "strategy-test-back/src/core/GameInstance"
	vector "strategy-test-back/src/core/Vector"
)

type TeleportTask struct {
	GameplayTask
	Target         *Character
	TargetPosition *vector.Vector2D
	Delay          float64
	startTime      float64
}

const TeleportTaskType GameplayTaskType = "teleport"

func NewTeleportTask(
	Target *Character,
	TargetPosition *vector.Vector2D,
	Delay float64,
) *TeleportTask {
	return &TeleportTask{
		GameplayTask: GameplayTask{
			ID:       TeleportTaskType,
			Priority: 1,
		},
		Target:         Target,
		TargetPosition: TargetPosition,
		Delay:          Delay,
	}
}

func (t *TeleportTask) OnAdd(
	tickContext TickContext,
	gi *GameInstance,
) {
	t.startTime = tickContext.Time
}

func (t *TeleportTask) OnTasksTick(
	tickContext TickContext,
	gi *GameInstance,
) {
	if tickContext.Time-t.startTime >= t.Delay {
		gi.AddImmediateEffect(&GameStateEffect{
			Type: TeleportEffect,
			Value: TeleportEffectPayload{
				TargetCharacter: t.Target,
				NewPosition:     *t.TargetPosition,
			},
		})

		gi.RemoveGameplayTask(t)
	}
}
