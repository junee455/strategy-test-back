package basicgameplaytasks

import (
	. "strategy-test-back/src/core/GameInstance"
	vector "strategy-test-back/src/core/Vector"
)

const MoveTaskType GameplayTaskType = "move"

type MoveTask struct {
	GameplayTask
	Target         *Character
	TargetPosition func() *vector.Vector2D
	targetReached  bool
	shouldFollow   bool

	lastNotificationPos  *vector.Vector2D
	notificationCooldown int
}

type TaskPhase int

const (
	Start TaskPhase = iota
	End
)

type MoveTaskNotification struct {
	Target *Character
	Phase  TaskPhase
	From   [2]float64
	To     [2]float64
	Speed  float64
}

func NewMoveTask(
	Target *Character,
	TargetPosition func() *vector.Vector2D,
	shouldFollow bool,
) *MoveTask {
	return &MoveTask{
		GameplayTask: GameplayTask{
			ID:       MoveTaskType,
			Priority: 1,
		},
		Target:               Target,
		TargetPosition:       TargetPosition,
		targetReached:        false,
		shouldFollow:         shouldFollow,
		lastNotificationPos:  &Target.Position,
		notificationCooldown: 0,
	}
}

func (t *MoveTask) OnFirstTick(tickContext TickContext, gi *GameInstance) {
	if t.GameplayTask.Ticked {
		return
	}

	t.GameplayTask.Ticked = true

	gi.EventNotifier.Notify(&EventNotification{
		Type: "move",
		Tick: tickContext.CurrentTick,
		Payload: MoveTaskNotification{
			Target: t.Target,
			Phase:  Start,
			From:   t.Target.Position,
			To:     *t.TargetPosition(),
			Speed:  t.Target.MoveSpeed,
		},
	})
}

func (t *MoveTask) OnTasksTick(
	tickContext TickContext,
	gi *GameInstance,
) {
	// compute move diff
	var charPos = &t.Target.Position
	var targetPos = t.TargetPosition()

	var vectorToTarget = targetPos.Sub(charPos)

	var vectorToTargetNormalized = vectorToTarget.Normalized()

	var distanceToTarget = vectorToTarget.Len()

	var charSpeed = t.Target.MoveSpeed

	var dv *vector.Vector2D

	if distanceToTarget > tickContext.Dt*charSpeed {
		dv = vectorToTargetNormalized.MulScalar(tickContext.Dt * charSpeed)
		t.targetReached = false
	} else {
		dv = vectorToTarget
		t.targetReached = true
	}

	gi.AddImmediateEffect(&GameStateEffect{
		Type: MoveEffect,
		Value: MoveEffectPayload{
			TargetCharacter: t.Target,
			Dv:              *dv,
		},
	})

	if !t.shouldFollow && t.targetReached {

		gi.EventNotifier.Notify(&EventNotification{
			Type: "move",
			Tick: tickContext.CurrentTick,
			Payload: MoveTaskNotification{
				Phase: End,
				From:  t.Target.Position,
				To:    *t.TargetPosition(),
			},
		})

		gi.RemoveGameplayTask(t)
	}
}

func (t *MoveTask) GetDescription() GameplayTaskDescription {
	var defaultDescription = t.GameplayTask.GetDescription()
	defaultDescription.Metadata = t.Target.RuntimeId
	return defaultDescription
}
