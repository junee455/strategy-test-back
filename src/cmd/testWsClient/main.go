package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func WsConnectHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	connection, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "test",
		})
		return
	}

	var closeChan = make(chan struct{}, 1)

	go func() {
		for {
			_, msg, err := connection.ReadMessage()
			if err != nil {
				fmt.Println("msg read error")
				closeChan <- struct{}{}
				connection.Close()
				break
			}

			fmt.Println(string(msg))
		}
	}()

	go func() {
		for {
			select {
			case <-closeChan:
				return
			default:
			}

			connection.WriteMessage(websocket.TextMessage, []byte("test"))
			time.Sleep(time.Second)
		}
	}()

}

func main() {
	fmt.Println("test ws client")

	r := mux.NewRouter()

	r.HandleFunc("/api/wsConnect", WsConnectHandler)

	http.ListenAndServe(":8000", r)
}
