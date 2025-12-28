package gameinstance

import (
	"fmt"
)

type ApplyDamageEffectNotification struct {
	InstigatorID  string `json:"instigatorId"`
	TargetID      string `json:"targetId"`
	DamageApplied int    `json:"damageApplied"`
	HpLeft        int    `json:"hpLeft"`
}

func ProcessApplyDamageEffect(effect *GameStateEffect, t TickContext, gi *GameInstance) {
	var payload = effect.Value.(ApplyDamageEffectPayload)

	var damageApplied int

	if payload.TargetCharacter.Stats.Health < payload.DamageAmount {
		damageApplied = payload.TargetCharacter.Stats.Health
	} else {
		damageApplied = payload.DamageAmount
	}

	payload.TargetCharacter.Stats.Health -= damageApplied

	gi.EventNotifier.Notify(&EventNotification{
		Type: "applyDamage",
		Tick: t.CurrentTick,
		Payload: ApplyDamageEffectNotification{
			InstigatorID:  string(payload.InstigatorCharacter.RuntimeId),
			TargetID:      string(payload.TargetCharacter.RuntimeId),
			DamageApplied: damageApplied,
			HpLeft:        payload.TargetCharacter.Stats.Health,
		},
	})
}

func ProcessMoveEffect(effect *GameStateEffect, t TickContext, gi *GameInstance) {
	fmt.Println("process move effect")

	var payload = effect.Value.(MoveEffectPayload)

	var targetCharacter = payload.TargetCharacter
	var dv = payload.Dv

	targetCharacter.Position = *targetCharacter.Position.Add(&dv)

	fmt.Println(payload.Dv)
}

func ProcessTeleportEffect(effect *GameStateEffect, t TickContext, gi *GameInstance) {
	var payload = effect.Value.(TeleportEffectPayload)

	payload.TargetCharacter.Position = payload.NewPosition
}

func ProcessSpawnEffect(effect *GameStateEffect, t TickContext, gi *GameInstance) {
	var payload = effect.Value.(SpawnEffectPayload)

	var newCharacter = Character{
		Actor:                *NewActor(),
		CharacterDescription: payload.CharacterDescription,
		Position:             payload.Position,
		Stats:                payload.Stats,
	}

	gi.Characters = append(gi.Characters, &newCharacter)
}

type MoveStopNotification struct {
	Target *Character
}

func ProcessBashEffect(effect *GameStateEffect, t TickContext, gi *GameInstance) {
	var targetCharacter = effect.Value.(*Character)

	for _, task := range gi.gameplayTasks {
		var description = task.GetDescription()

		if description.ID == "move" && description.Metadata == targetCharacter.RuntimeId {
			gi.RemoveGameplayTask(task)
			gi.EventNotifier.Notify(&EventNotification{
				Type: "moveStop",
				Tick: t.CurrentTick,
				Payload: MoveStopNotification{
					Target: targetCharacter,
				},
			})
		}
	}
}
