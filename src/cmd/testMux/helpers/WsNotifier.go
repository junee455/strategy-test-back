package helpers

import (
	"encoding/json"
	"fmt"
	gameinstance "strategy-test-back/src/core/GameInstance"
	basicgameplaytasks "strategy-test-back/src/core/GameInstance/BasicGameplayTasks"

	"github.com/gorilla/websocket"
)

type WsNotifier struct {
	connections []*websocket.Conn
}

func (w *WsNotifier) AddConnection(c *websocket.Conn) {
	w.connections = append(w.connections, c)
}

func (w *WsNotifier) NotifyAll(notification *WsNotification) {
	// var msgJson str

	var jsonMsg, err = json.Marshal(notification)

	if err != nil {
		return
	}

	for _, c := range w.connections {
		c.WriteMessage(websocket.TextMessage, jsonMsg)
	}
}

func NewWsNotifier() *WsNotifier {
	return &WsNotifier{}
}

func (n *WsNotifier) Notify(e *gameinstance.EventNotification) {
	var httpNotification = &WsNotification{
		Type: e.Type,
		Tick: e.Tick,
	}

	switch e.Type {
	case "move":
		{
			var payload, castOk = e.Payload.(basicgameplaytasks.MoveTaskNotification)
			if !castOk {
				return
			}

			if payload.Phase == basicgameplaytasks.End {
				return
			}

			httpNotification.Payload = &WsMovePayload{
				ActorID: string(payload.Target.RuntimeId),
				From:    payload.From,
				To:      payload.To,
				Speed:   payload.Speed,
			}
		}
	case "applyDamage":
		{
			fmt.Println("APPLY DAMAGE NOTIFY")

			var payload, castOk = e.Payload.(gameinstance.ApplyDamageEffectNotification)
			if !castOk {
				return
			}

			httpNotification.Payload = &WsApplyDamagePayload{
				InstigatorID:  payload.InstigatorID,
				TargetID:      payload.TargetID,
				DamageApplied: payload.DamageApplied,
				HpLeft:        payload.HpLeft,
			}

		}
	case "castProjectile":
		{
			var payload, castOk = e.Payload.(basicgameplaytasks.CastProjectileNotification)
			if !castOk {
				return
			}

			var casterId string = ""
			var targetId string = ""

			if payload.Caster != nil {
				casterId = string(payload.Caster.RuntimeId)
			}

			if payload.Target != nil {
				targetId = string(payload.Target.RuntimeId)
			}

			httpNotification.Payload = &WsCastProjectilePayload{
				Type:         payload.ProjectileType,
				InstigatorID: casterId,
				TargetID:     targetId,
				From:         payload.From,
				To:           payload.To,
				Speed:        payload.Speed,
			}

		}
	case "effect":
		{
			var payload, castOk = e.Payload.(basicgameplaytasks.BasicEffectNotification)

			if !castOk {
				return
			}

			if payload.Target == nil {
				return
			}

			httpNotification.Payload = &WsApplyEffectPayload{
				EffectType: int(payload.EffectType),
				TargetID:   string(payload.Target.RuntimeId),
				Remove:     payload.Remove,
			}
		}
	case "moveStop":
		{
			var payload, castOk = e.Payload.(gameinstance.MoveStopNotification)

			if !castOk {
				return
			}

			if payload.Target == nil {
				return
			}

			httpNotification.Payload = &WsMoveStopPayload{
				ActorID: string(payload.Target.RuntimeId),
			}
		}
	}

	n.NotifyAll(httpNotification)
}
