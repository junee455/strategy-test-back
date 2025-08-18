package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const ONE_SEC = 1_000_000_000

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(ctx *gin.Context) {
	w, r := ctx.Writer, ctx.Request
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	fmt.Println("socket open")
	defer c.Close()
	for {
		data := map[string]interface{}{
			"type": "greeting",
			"body": "Hello, world!",
		}

		c.WriteJSON(data)

		time.Sleep(ONE_SEC)
	}
}

var connectionCounter = 0

func connect(ctx *gin.Context) {
	connectionCounter += 1

	var websocket = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	w, r := ctx.Writer, ctx.Request
	c, err := websocket.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}

	fmt.Printf("socket open %v\n", connectionCounter)

	defer func() {
		connectionCounter -= 1
		fmt.Printf("connection close")
		c.Close()
	}()

	for {
		data := map[string]interface{}{
			"connectionCount": connectionCounter,
		}

		c.WriteJSON(data)

		time.Sleep(ONE_SEC)
	}
}

func main() {
	fmt.Println("Hi from rts game")

	var gameInstance = GameInstance{
		clients:   []GameClient{},
		gameState: GameState{},
		id:        "1",
	}

	fmt.Printf("%v", gameInstance)

	r := gin.Default()
	r.GET("/connect", connect)
	r.GET("/echo", echo)
	r.Run(":8080")
}
