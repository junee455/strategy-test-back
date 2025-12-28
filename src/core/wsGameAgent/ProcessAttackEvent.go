package wsgameagent

import (
	"encoding/json"
	"fmt"
	gameinstance "strategy-test-back/src/core/GameInstance"
	basicgameplaytasks "strategy-test-back/src/core/GameInstance/BasicGameplayTasks"
)

type WsAttackEventPayload struct {
	ActorID  string `json:"actorId"`
	TargetID string `json:"targetId"`
	Follow   bool   `json:"follow"`
	Single   bool   `json:"single"`
}

type WsAttackEvent = WsMessage[WsAttackEventPayload]

func ProcessAttackEvent(gi *gameinstance.GameInstance, msg []byte) {
	var attackEvent WsAttackEvent
	var payload WsAttackEventPayload

	json.Unmarshal([]byte(msg), &attackEvent)

	payload = attackEvent.Payload

	fmt.Printf("Attack event: %v\n", payload)

	var casterCharacter = gi.FindCharacterByID(gameinstance.ActorID(payload.ActorID))
	var targetCharacter = gi.FindCharacterByID(gameinstance.ActorID(payload.TargetID))

	if casterCharacter == nil || targetCharacter == nil {
		return
	}

	var onReach = func() {
		gi.AddImmediateEffect(&gameinstance.GameStateEffect{
			Type: gameinstance.ApplyDamageEffect,
			Value: gameinstance.ApplyDamageEffectPayload{
				InstigatorCharacter: casterCharacter,
				TargetCharacter:     targetCharacter,
				DamageAmount:        10,
				DamageType:          gameinstance.Physical,
			},
		})
	}

	gi.AddGameplayTask(
		basicgameplaytasks.NewCastProjectileTask(
			casterCharacter,
			nil,
			targetCharacter,
			nil,
			"attack",
			casterCharacter.ProjectileSpeed,
			&onReach,
		))
}
