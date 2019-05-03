package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Kind byte

// Kind enums
const (
	Char Kind = 0
	Noun Kind = 1
	Adj  Kind = 2
	Conj Kind = 3
)

type Pos struct {
	x int
	y int
}

type Id int // Randomized id?

type Object struct {
	pos  Pos
	kind Kind
	item string
	id   Id
}

type Level struct {
	objects []Object
	height  int
	width   int
}

type Event string // move,  death, spawn (transform = death+spawn)

type Change struct {
	event Event
	item  string
	id    Id
	pos   Pos
}
type Tick []Change
type Timeline []Tick

func main() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	func handler(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
