package wsgameagent

import (
	"encoding/json"
	"fmt"
	"net/http"
	agent "strategy-test-back/src/core"
	gameinstance "strategy-test-back/src/core/GameInstance"

	"github.com/gorilla/websocket"
)

type ConnectionState int

const (
	Idle ConnectionState = iota
	Error
	Connected
)

type WsMessage[PayloadType any] struct {
	Type    string      `json:"type"`
	Payload PayloadType `json:"payload,omitempty"`
}

type WsGameAgent struct {
	ClientID        string
	websocket       websocket.Upgrader
	Connection      *websocket.Conn
	connectionState ConnectionState
	gameInstance    *gameinstance.GameInstance
}

func getMsgType(msg []byte) (string, error) {
	type WsMsgType struct {
		Type string `json:"type"`
	}

	var msgType WsMsgType

	var err = json.Unmarshal(msg, &msgType)

	if err != nil {
		return "", err
	}

	return msgType.Type, nil
}

func NewWsGameAgent(
	ClientID string,
	gameInstance *gameinstance.GameInstance,
) *WsGameAgent {
	var newAgent = WsGameAgent{
		ClientID:     ClientID,
		gameInstance: gameInstance,
		websocket: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	return &newAgent
}

func (ga *WsGameAgent) GetAgentDescription() agent.AgentDescription {
	var stateDescription string = "unknown"

	switch ga.connectionState {
	case Connected:
		stateDescription = "connected"
	case Error:
		stateDescription = "error"
	}

	return agent.AgentDescription{
		ID:    ga.ClientID,
		State: stateDescription,
	}
}

func (ga *WsGameAgent) EstablishConnection(w http.ResponseWriter, r *http.Request) error {
	var connection, err = ga.websocket.Upgrade(w, r, nil)

	if err != nil {
		ga.connectionState = Error
	} else {
		ga.connectionState = Connected
	}

	ga.Connection = connection

	// read msg loop
	go func() {
		for {
			defer connection.Close()

			_, msg, err := connection.ReadMessage()

			if err != nil {
				ga.connectionState = Error
				return
			}

			var msgType string
			msgType, err = getMsgType(msg)

			if err != nil {
				fmt.Printf("error          : %v\n", err)
				fmt.Printf("failed to parse: %v\n", string(msg))
				continue
			}

			if msgType == "ping" {
				fmt.Println(".")
				connection.WriteJSON(map[string]string{
					"type": "pong",
				})
				continue
			}

			ga.ProcessInputEvent(msgType, msg)
		}
	}()

	return err
}

func (ga *WsGameAgent) ProcessInputEvent(msgType string, msg []byte) {
	// var decoder = json.NewDecoder()

	var gameInstance = ga.gameInstance

	switch msgType {
	case "move":
		ProcessMoveEvent(gameInstance, msg)
	case "attack":
		ProcessAttackEvent(gameInstance, msg)
	case "useAbility":
		ProcessUseAbilityEvent(gameInstance, msg)
	}
}
