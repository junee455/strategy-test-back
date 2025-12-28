package wsgameagent

import (
	"encoding/json"
	"fmt"
	gameinstance "strategy-test-back/src/core/GameInstance"
	basicgameplaytasks "strategy-test-back/src/core/GameInstance/BasicGameplayTasks"
	vector "strategy-test-back/src/core/Vector"
)

type WsMoveEventPayload struct {
	ActorID string     `json:"actorId"`
	To      [2]float64 `json:"to"`
}

type WsMoveEvent = WsMessage[WsMoveEventPayload]

func ProcessMoveEvent(gi *gameinstance.GameInstance, msg []byte) {
	var moveEvent WsMoveEvent

	var payload WsMoveEventPayload

	json.Unmarshal([]byte(msg), &moveEvent)

	payload = moveEvent.Payload

	fmt.Printf("Move event: %v\n", payload)

	var targetCharacter = gi.FindCharacterByID(gameinstance.ActorID(payload.ActorID))

	if targetCharacter == nil {
		return
	}

	var moveToPosition = vector.New(payload.To)

	var getPositionFunc = func() *vector.Vector2D {
		return moveToPosition
	}

	gi.AddGameplayTask(
		basicgameplaytasks.NewMoveTask(
			targetCharacter,
			getPositionFunc,
			false,
		))
}
