package routes

import (
	"encoding/json"
	"net/http"
	appstate "strategy-test-back/src/cmd/testMux/appState"
	gameinstance "strategy-test-back/src/core/GameInstance"

	// clientevents "strategy-test-back/src/core/types/clientEvents"
	// gameEvents "strategy-test-back/src/core/types/gameEvents/base"
	// gameEventsPayloads "strategy-test-back/src/core/types/gameEvents/gameEventsPayloads"

	"github.com/gorilla/mux"
)

func GetFullGameState(w http.ResponseWriter, r *http.Request) {
	var gameId = mux.Vars(r)["id"]

	var gameInstance = appstate.FindGameById(gameId)

	if gameInstance == nil {
		http.Error(w, "game not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	type ClientState struct {
		ID    string `json:"id"`
		State string `json:"state"`
	}

	type GetGameStateResponse struct {
		ID               string                    `json:"id"`
		OwnerID          string                    `json:"ownerId"`
		ClientsConnected []ClientState             `json:"clients"`
		Characters       []*gameinstance.Character `json:"characters"`
	}

	var clientStates []ClientState = make([]ClientState, 0)

	for _, c := range gameInstance.Agents {
		agentDescripiton := c.GetAgentDescription()

		clientStates = append(clientStates, ClientState{
			ID:    agentDescripiton.ID,
			State: agentDescripiton.State,
		})
	}

	var response = GetGameStateResponse{
		ID:               gameInstance.ID,
		OwnerID:          gameInstance.OwnerID,
		ClientsConnected: clientStates,
		Characters:       gameInstance.Game.Characters,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
