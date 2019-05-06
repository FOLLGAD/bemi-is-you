package main

import "fmt"

type Kind byte

// Kind enums
const (
	Char Kind = 0
	Noun Kind = 1
	Adj  Kind = 2
	Conj Kind = 3
)

type Pos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Id int // Randomized id?

type Object struct {
	Pos  Pos    `json:"pos"`
	Kind Kind   `json:"kind"`
	Item string `json:"item"`
	Id   Id     `json:"id"`
}

type Level struct {
	Objects []*Object `json:"objects"`
	Height  int       `json:"height"`
	Width   int       `json:"width"`
}

type Event int

const (
	Move  Event = 0
	Death Event = 1
	Spawn Event = 2
	// Transform: death+spawn
)

type Change struct {
	Event Event `json:"event"`
	Id    Id    `json:"id"`
	Pos   Pos   `json:"pos"` // Move: deltapos
}
type Tick []Change
type Timeline []Tick

// Game struct

type Game struct {
	level       Level
	objectState []*Object
	timeline    Timeline
	updateChan  chan<- Tick
}

func (game *Game) ReceiveData(msg ReceivedMessage) {
	switch msg.Data {
	case "up":
		id := game.objectState[0].Id
		moveChange := Change{
			Event: Move,
			Id:    id,
			Pos:   Pos{0, -1},
		}
		game.AddTick(Tick{moveChange})
	case "down":
		id := game.objectState[0].Id
		moveChange := Change{
			Event: Move,
			Id:    id,
			Pos:   Pos{0, 1},
		}
		game.AddTick(Tick{moveChange})
	case "right":
		id := game.objectState[0].Id
		moveChange := Change{
			Event: Move,
			Id:    id,
			Pos:   Pos{1, 0},
		}
		game.AddTick(Tick{moveChange})
	case "left":
		id := game.objectState[0].Id
		moveChange := Change{
			Event: Move,
			Id:    id,
			Pos:   Pos{-1, 0},
		}
		game.AddTick(Tick{moveChange})
	}
	for _, e := range game.objectState {
		fmt.Println(*e)
	}
}

func (game *Game) FindId(id Id) *Object {
	for _, e := range game.objectState {
		if e.Id == id {
			return e
		}
	}
	return &Object{} // Doesn't happen
}

func (game *Game) DoChange(change Change) {
	switch change.Event {
	case Move:
		obj := game.FindId(change.Id)
		obj.Pos.X += change.Pos.X
		obj.Pos.Y += change.Pos.Y
	}
}

func (game *Game) EmitDelta(tick Tick) {
	game.updateChan <- tick
}

func (game *Game) DoTick(tick Tick) {
	for _, change := range tick {
		game.DoChange(change)
	}
	game.EmitDelta(tick)
}

func (game *Game) AddTick(tick Tick) {
	game.DoTick(tick)
	game.timeline = append(game.timeline, tick)
}

func (game *Game) Undo() {
	length := len(game.timeline)
	lastTick := game.timeline[length-1]
	game.timeline = game.timeline[:len(game.timeline)-2] // Remove last tick

	for i := len(lastTick) - 1; i >= 0; i-- {
		game.DoChange(lastTick[i])
	}
}
