package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	appstate "strategy-test-back/src/cmd/testMux/appState"
	"strategy-test-back/src/cmd/testMux/helpers"
	silencer "strategy-test-back/src/core/Characters/silencer"
	skeleton "strategy-test-back/src/core/Characters/skeleton"
	gameinstance "strategy-test-back/src/core/GameInstance"
	vector "strategy-test-back/src/core/Vector"
)

func NewGameSession(w http.ResponseWriter, r *http.Request) {
	clientMetaData, err := helpers.ReadClientMetaData(r)

	fmt.Println("start new game request")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var newGame = appstate.StartNewGame(clientMetaData.ClientId)
	var gameId = newGame.ID

	var gameInstance = newGame.Game

	gameInstance.AddImmediateEffect(&gameinstance.GameStateEffect{
		Type: gameinstance.Spawn,
		Value: gameinstance.SpawnEffectPayload{
			CharacterDescription: silencer.GetDefaults(),
			Position:             *vector.New([2]float64{0, 0}),
			Stats: func() gameinstance.Stats {
				var stats = silencer.GetDefaults().InitialStats
				stats.Health -= 20

				return stats
			}(),
		},
	})

	gameInstance.AddImmediateEffect(&gameinstance.GameStateEffect{
		Type: gameinstance.Spawn,
		Value: gameinstance.SpawnEffectPayload{
			CharacterDescription: skeleton.GetDefaults(),
			Position:             *vector.New([2]float64{0, 3}),
			Stats: func() gameinstance.Stats {
				var stats = skeleton.GetDefaults().InitialStats
				stats.Health -= 20

				return stats
			}(),
		},
	})

	fmt.Println("game created")

	var response = struct {
		ID      string `json:"id"`
		OwnerID string `json:"ownerId"`
	}{
		ID:      gameId,
		OwnerID: clientMetaData.ClientId,
	}

	go newGame.Game.Start()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
