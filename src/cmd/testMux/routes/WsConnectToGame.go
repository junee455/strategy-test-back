package routes

import (
	"fmt"
	"net/http"
	appstate "strategy-test-back/src/cmd/testMux/appState"
	"strategy-test-back/src/cmd/testMux/helpers"
	agent "strategy-test-back/src/core"
	wsgameagent "strategy-test-back/src/core/wsGameAgent"

	// gameclient "strategy-test-back/src/core/game/gameClient"
	"github.com/gorilla/mux"
)

func WsConnectToGame(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connecting...")

	clientMetaData, err := helpers.ReadClientMetaData(r)

	if err != nil {
		fmt.Println("client metadata missing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var gameId = mux.Vars(r)["id"]

	// find game instance
	var gameInstance = appstate.FindGameById(gameId)

	if gameInstance == nil {
		fmt.Println("Game not found")
		return
	}

	var agent agent.IAgent

	for _, v := range gameInstance.Agents {
		if v.GetAgentDescription().ID == clientMetaData.ClientId {
			agent = v
		}
	}

	if agent == nil {
		fmt.Println("agent not found")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var wsAgent, castOk = agent.(*wsgameagent.WsGameAgent)

	if !castOk {
		fmt.Println("agent cast fail")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = wsAgent.EstablishConnection(w, r)

	if err != nil {
		return
	}

	gameInstance.Notifier.AddConnection(wsAgent.Connection)
}
