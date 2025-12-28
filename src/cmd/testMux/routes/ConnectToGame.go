package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	appstate "strategy-test-back/src/cmd/testMux/appState"
	"strategy-test-back/src/cmd/testMux/helpers"
	agent "strategy-test-back/src/core"
	wsgameagent "strategy-test-back/src/core/wsGameAgent"

	"github.com/gorilla/mux"
)

func ConnectToGame(w http.ResponseWriter, r *http.Request) {
	type ConnectToGameResponse struct {
		ClientID string `json:"clientId"`
	}

	var gameId = mux.Vars(r)["id"]

	clientMetadata, err := helpers.ReadClientMetaData(r)

	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	var clientId = clientMetadata.ClientId

	r.Body.Close()

	fmt.Printf("client id: %v\ngame id: %v\n", clientId, gameId)

	var gameManager = appstate.FindGameById(gameId)

	if gameManager == nil {
		fmt.Println("Game not found")
		return
	}

	var agent agent.IAgent

	for _, v := range gameManager.Agents {
		if v.GetAgentDescription().ID == clientId {
			agent = v
			break
		}
	}

	if agent == nil {
		agent = wsgameagent.NewWsGameAgent(clientId, gameManager.Game)
		gameManager.AddAgent(agent)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(ConnectToGameResponse{
		ClientID: agent.GetAgentDescription().ID,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
