package wsgameagent

import (
	"encoding/json"
	"fmt"
	gameinstance "strategy-test-back/src/core/GameInstance"
	basicgameplaytasks "strategy-test-back/src/core/GameInstance/BasicGameplayTasks"
	vector "strategy-test-back/src/core/Vector"
	"time"
	// basicgameplaytasks "strategy-test-back/src/core/GameInstance/BasicGameplayTasks"
	// vector "strategy-test-back/src/core/Vector"
)

/*
			cast at point:
			{
	    "actorId": "1",
	    "ability": "arcaneCurse",
	    "target": "point",
	    "point": [
	        5.888035964535192,
	        3.105207037093084
	    ]
			}
			cast at character:
			{
	    "actorId": "1",
	    "ability": "arcaneCurse",
	    "target": "2"
			}
*/

type WsUseAbilityEventPayload struct {
	ActorID string     `json:"actorId"`
	Ability string     `json:"ability"`
	Target  string     `json:"target,omitempty"`
	Point   [2]float64 `json:"point,omitempty"`
}

type WsUseAbilityEvent = WsMessage[WsUseAbilityEventPayload]

func ProcessUseAbilityEvent(gi *gameinstance.GameInstance, msg []byte) {
	var useAbilityEvent WsUseAbilityEvent

	var payload WsUseAbilityEventPayload

	json.Unmarshal([]byte(msg), &useAbilityEvent)

	payload = useAbilityEvent.Payload

	fmt.Printf("Use ability event: %v\n", payload)

	var casterCharacter = gi.FindCharacterByID(gameinstance.ActorID(payload.ActorID))

	if casterCharacter.CharacterID != "silencer" {
		fmt.Println("only silencer abilities are implemented...")
		return
	}

	var targetCharacter *gameinstance.Character
	var castPoint *vector.Vector2D

	if casterCharacter == nil {
		return
	}

	switch payload.Target {
	case "point":
		{
			castPoint = vector.New(payload.Point)
		}
	case "self":
		{
			targetCharacter = casterCharacter
		}
	default:
		{
			var tryTargetID = gameinstance.ActorID(payload.Target)
			targetCharacter = gi.FindCharacterByID(tryTargetID)

			if targetCharacter == nil {
				return
			}
		}
	}

	var onReach = func() {
		fmt.Println("silencer projectile hit!!!")
		if targetCharacter == nil {
			return
		}

		var endChan = make(chan struct{}, 1)

		go func() {
			time.Sleep(time.Second * 3)
			endChan <- struct{}{}
		}()

		gi.AddGameplayTask(
			basicgameplaytasks.NewBasicEffectTask(
				casterCharacter,
				targetCharacter,
				gameinstance.Bash,
				endChan,
			))
	}

	gi.AddGameplayTask(
		basicgameplaytasks.NewCastProjectileTask(
			casterCharacter,
			nil,
			targetCharacter,
			castPoint,
			payload.Ability,
			casterCharacter.ProjectileSpeed,
			&onReach,
		))

	// var moveToPosition = vector.New(payload.To)

	// var getPositionFunc = func() *vector.Vector2D {
	// 	return moveToPosition
	// }

	// gi.AddGameplayTask(
	// 	basicgameplaytasks.NewMoveTask(
	// 		casterCharacter,
	// 		getPositionFunc,
	// 		false,
	// 	))
}
