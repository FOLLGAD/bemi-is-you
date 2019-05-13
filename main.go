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

var idCounter Id = 0

func addObject(pos Pos, kind Kind, item string) *Object {
	obj := &Object{
		pos,
		kind,
		item,
		idCounter,
	}
	idCounter++
	return obj
}

func main() {
	// LEVEL
	var firstLevel = Level{}
	firstLevel.Width = 18
	firstLevel.Height = 12

	firstLevel.Objects = []*Object{
		addObject(Pos{10, 10}, Char, "bemi"),
		addObject(Pos{8, 10}, Char, "sami"),
		addObject(Pos{0, 0}, Noun, "bemi"),
		addObject(Pos{1, 0}, Conj, "is"),
		addObject(Pos{2, 0}, Adj, "1"),
		addObject(Pos{0, 5}, Noun, "sami"),
		addObject(Pos{1, 5}, Conj, "is"),
		addObject(Pos{2, 5}, Adj, "1"),
		addObject(Pos{0, 6}, Conj, "is"),
		addObject(Pos{0, 7}, Adj, "2"),
	}
	// ******

	players := []Player{}

	updateChan := make(chan Tick)

	go func() {
		// Read incoming tick updates and them broadcast to all players
		for {
			tick := <-updateChan
			if len(tick) > 0 {
				js, _ := json.Marshal(Message{tick, 1})
				for _, p := range players {
					p.conn.WriteJSON(js)
				}
			}
		}
	}()

	game := Game{firstLevel, firstLevel.Objects, Timeline{}, updateChan}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		// Websocket logic
		conn, _ := upgrader.Upgrade(w, r, nil)
		js, _ := json.Marshal(Message{firstLevel, 0})

		playerNum := len(players) + 1
		newPlayer := Player{playerNum, conn}
		players = append(players, newPlayer)
		fmt.Println("New player", newPlayer.number)

		conn.WriteJSON(js)
		for {
			_, p, err := conn.ReadMessage()

			if err != nil {
				// Connetion timed out
				fmt.Println(err)
				conn.Close()

				for i := range players {
					if players[i].number == playerNum {
						players = append(players[:i], players[i+1:]...)
						return
					}
				}
			}

			var message ReceivedMessage
			message.player = newPlayer.number
			json.Unmarshal(p, &message)

			game.ReceiveData(message)
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://localhost:8080", http.StatusSeeOther)
	})
	http.ListenAndServe(":8081", nil)
}
