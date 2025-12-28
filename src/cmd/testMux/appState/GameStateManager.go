package appstate

import (
	"strategy-test-back/src/cmd/testMux/helpers"
	agent "strategy-test-back/src/core"
	gameinstance "strategy-test-back/src/core/GameInstance"
	"strconv"
)

var lastGameID = 0

type GameStateManager struct {
	ID       string `json:"id"`
	OwnerID  string `json:"ownerId"`
	Notifier *helpers.WsNotifier
	Agents   []agent.IAgent
	Game     *gameinstance.GameInstance
}

func NewGameStateManager(ownerId string) *GameStateManager {
	lastGameID++

	var newGame = gameinstance.NewGameInstance()
	var wsNotifier = helpers.NewWsNotifier()
	newGame.EventNotifier = wsNotifier

	return &GameStateManager{
		ID:       strconv.Itoa(lastGameID),
		OwnerID:  ownerId,
		Agents:   make([]agent.IAgent, 0),
		Game:     newGame,
		Notifier: wsNotifier,
	}
}

func (gm *GameStateManager) AddAgent(agent agent.IAgent) {
	gm.Agents = append(gm.Agents, agent)
}
