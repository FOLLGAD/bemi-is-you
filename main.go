package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Data    interface{} `json:"data"`
	MsgType int         `json:"msgType"`
	Room    int32       `json:"room"`
}

type ReceivedMessage struct {
	Data    interface{} `json:"data"`
	MsgType int         `json:"msgType"`
	player  int
}

var sessions = map[int32]Session{}

type Session struct {
	game    *Game
	room    int32
	players map[int]Player
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

func main() {
	loadTileSet()

	rand.Seed(time.Now().UnixNano())

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		room, err := strconv.Atoi(r.URL.Query().Get("room"))
		if err != nil {
			w.Write([]byte("Invalid room"))
			return
		}
		session, ok := sessions[int32(room)]
		if !ok {
			w.Write([]byte("Room doesn't exist"))
			return
		}
		session.join(w, r)
	})
	http.HandleFunc("/sessions", func(w http.ResponseWriter, r *http.Request) {
		// Normal http-get request for fetching an array of available sessions
		fmt.Println("Fetching sessions")

		type str struct {
			Room    int32 `json:"room"`
			Players int   `json:"players"`
		}

		// Format sessions into a more json-compatible format
		array := []str{}
		for room, ses := range sessions {
			array = append(array, str{Room: room, Players: len(ses.players)})
		}

		v, _ := json.Marshal(array)
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow CORS
		w.Write(v)
	})
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		session := createSession()
		session.join(w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	log.Fatal(http.ListenAndServe(":8123", nil))
}

func (session Session) join(w http.ResponseWriter, r *http.Request) {
	// Websocket logic
	conn, _ := upgrader.Upgrade(w, r, nil)

	playerNum := 0

	i := 1
	for playerNum == 0 {
		_, ok := session.players[i]
		if !ok {
			playerNum = i
		}
		i++
	}
	conn.WriteJSON(&Message{playerNum, 2, session.room})

	newPlayer := Player{playerNum, conn}
	session.players[playerNum] = newPlayer
	fmt.Println("New player", newPlayer.number)

	conn.WriteJSON(&Message{session.game, 0, session.room})
	go func() {
		for {
			_, p, err := conn.ReadMessage()

			// Connection broken; remove player from room
			if err != nil {
				fmt.Println(err)
				conn.Close()

				for i := range session.players {
					if session.players[i].number == playerNum {
						delete(session.players, playerNum)
						if len(session.players) == 0 {
							delete(sessions, session.room)
						}
						return
					}
				}
			}

			var message ReceivedMessage
			message.player = newPlayer.number
			json.Unmarshal(p, &message) // Parse message from json

			session.game.ReceiveData(message)
		}
	}()
}

func createSession() (session Session) {
	updateChan := make(chan Message)

	go func() {
		// Read incoming tick updates and them broadcast to all players
		for {
			message := <-updateChan
			for _, p := range session.players {
				p.conn.WriteJSON(&message)
			}
		}
	}()

	game := MakeGame(getLevel(0), updateChan)

	// Generate a random room number, and make sure it's unique
	var roomNum int32
	ok := true
	for ok {
		roomNum = rand.Int31()
		_, ok = sessions[roomNum]
	}

	fmt.Println("Created new room with number", roomNum)

	sessions[roomNum] = Session{
		game:    game,
		players: map[int]Player{},
		room:    roomNum,
	}
	session = sessions[roomNum]
	return session
}
