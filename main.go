package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Data    interface{} `json:"data"`
	MsgType int         `json:"msgType"`
}

type ReceivedMessage struct {
	Data    interface{} `json:"data"`
	MsgType int         `json:"msgType"`
	player  int
}

type Player struct {
	number int
	conn   *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	// LEVEL
	var firstLevel = Level{}
	firstLevel.Height = 20
	firstLevel.Width = 25
	var bemi = &Object{Pos{10, 10}, Char, "bemi", 0}
	firstLevel.Objects = []*Object{bemi}
	// ******

	players := []Player{}

	updateChan := make(chan Tick)

	go func() {
		// Read incoming tick updates and them broadcast to all players
		for {
			tick := <-updateChan
			js, _ := json.Marshal(Message{tick, 1})
			for _, p := range players {
				p.conn.WriteJSON(js)
			}
		}
	}()

	game := Game{firstLevel, firstLevel.Objects, Timeline{}, updateChan}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		// Websocket logic
		conn, _ := upgrader.Upgrade(w, r, nil)
		js, _ := json.Marshal(Message{firstLevel, 0})

		playerPos := len(players)
		newPlayer := Player{playerPos, conn}
		players = append(players, newPlayer)
		fmt.Println("New player ", newPlayer.number)

		conn.WriteJSON(js)
		for {
			_, p, err := conn.ReadMessage()

			if err != nil {
				fmt.Println(err)
				conn.Close()

				players = append(players[:playerPos], players[playerPos+1:]...)

				return
			}

			var message ReceivedMessage
			message.player = newPlayer.number
			json.Unmarshal(p, &message)

			game.ReceiveData(message)
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string("hello"))
	})
	http.ListenAndServe(":8081", nil)
}
