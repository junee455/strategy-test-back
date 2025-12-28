package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strategy-test-back/src/cmd/testMux/routes"
)

func main() {
	fmt.Println("~~ Test mux ~~")

	r := mux.NewRouter()

	r.HandleFunc("/test/login", routes.TestLoginListener).Methods("POST")

	r.HandleFunc("/game/list", routes.ListGamesHandler)
	r.HandleFunc("/game/newGameSession", routes.NewGameSession).Methods("POST")
	r.HandleFunc("/game/{id}/getFullGameState", routes.GetFullGameState).Methods("GET")
	r.HandleFunc("/game/{id}/connect", routes.ConnectToGame).Methods("GET")
	r.HandleFunc("/game/{id}/wsConnect", routes.WsConnectToGame)

	http.ListenAndServe(":8000", r)
}
