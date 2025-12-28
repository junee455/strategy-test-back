package routes

import (
	"encoding/json"
	"net/http"
	appstate "strategy-test-back/src/cmd/testMux/appState"
)

func ListGamesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var idsList = appstate.ListGameIds()

	if err := json.NewEncoder(w).Encode(idsList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
