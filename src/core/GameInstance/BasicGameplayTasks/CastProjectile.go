package basicgameplaytasks

import (
	. "strategy-test-back/src/core/GameInstance"
	vector "strategy-test-back/src/core/Vector"
)

type CastProjectileTask struct {
	GameplayTask

	Caster          *Character
	TargetPosition  *vector.Vector2D
	TargetCharacter *Character
	CurrentPosition *vector.Vector2D

	ProjectileType string

	speed         float64
	targetReached bool
	onReach       *func()
}

const CastProjectileTaskType GameplayTaskType = "castProjectile"

func NewCastProjectileTask(
	Caster *Character,
	CurrentPosition *vector.Vector2D,

	TargetCharacter *Character,
	TargetPosition *vector.Vector2D,
	ProjectileType string,

	speed float64,
	onReach *func(),
) *CastProjectileTask {
	if CurrentPosition == nil && Caster == nil {
		return nil
	}

	if TargetPosition == nil && TargetCharacter == nil {
		return nil
	}

	if CurrentPosition == nil {
		CurrentPosition = &Caster.Position
	}

	if TargetPosition == nil {
		TargetPosition = &TargetCharacter.Position
	}

	return &CastProjectileTask{
		GameplayTask: GameplayTask{
			ID:       CastProjectileTaskType,
			Priority: 1,
		},
		Caster:          Caster,
		TargetPosition:  TargetPosition,
		CurrentPosition: CurrentPosition,
		TargetCharacter: TargetCharacter,
		ProjectileType:  ProjectileType,
		speed:           speed,
		targetReached:   false,
		onReach:         onReach,
	}
}

func (t *CastProjectileTask) OnTasksTick(
	tickContext TickContext,
	gi *GameInstance,
) {
	// compute move diff
	var projectilePos = t.CurrentPosition
	var targetPos *vector.Vector2D

	if t.TargetCharacter != nil {
		targetPos = &t.TargetCharacter.Position
	} else {
		targetPos = t.TargetPosition
	}

	var vectorToTarget = targetPos.Sub(projectilePos)

	var vectorToTargetNormalized = vectorToTarget.Normalized()

	var distanceToTarget = vectorToTarget.Len()

	var charSpeed = t.speed

	var dv *vector.Vector2D

	if distanceToTarget > tickContext.Dt*charSpeed {
		dv = vectorToTargetNormalized.MulScalar(tickContext.Dt * charSpeed)
	} else {
		dv = vectorToTarget
		t.targetReached = true
	}

	t.CurrentPosition = t.CurrentPosition.Add(dv)

	if t.targetReached {
		gi.RemoveGameplayTask(t)

		if t.onReach != nil {
			(*t.onReach)()
		}
	}
}

type CastProjectileNotification struct {
	Phase          TaskPhase
	Caster         *Character
	Target         *Character
	ProjectileType string
	From           [2]float64
	To             [2]float64
	Speed          float64
}

func (t *CastProjectileTask) OnFirstTick(
	tickContext TickContext,
	gi *GameInstance,
) {
	if t.GameplayTask.Ticked {
		return
	}

	t.GameplayTask.Ticked = true

	gi.EventNotifier.Notify(&EventNotification{
		Type: "castProjectile",
		Tick: tickContext.CurrentTick,
		Payload: CastProjectileNotification{
			Phase:          Start,
			Caster:         t.Caster,
			Target:         t.TargetCharacter,
			ProjectileType: t.ProjectileType,
			From:           *t.CurrentPosition,
			To:             *t.TargetPosition,
			Speed:          t.Caster.ProjectileSpeed,
		},
	})
}
