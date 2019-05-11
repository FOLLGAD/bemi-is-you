package main

import (
	"strconv"
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
	X int `json:"x"`
	Y int `json:"y"`
}

func (pos Pos) Add(addPos Pos) Pos {
	return Pos{
		pos.X + addPos.X,
		pos.Y + addPos.Y,
	}
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

type ObjectList []*Object

type Game struct {
	level       Level
	objectState ObjectList
	timeline    Timeline
	updateChan  chan<- Tick
}

func (objs ObjectList) findCharByItem(item string) ObjectList {
	outs := []*Object{}
	for _, e := range objs {
		if e.Kind == Char && e.Item == item {
			outs = append(outs, e)
		}
	}
	return outs
}

func (game *Game) ReceiveData(msg ReceivedMessage) {
	meanings := findSentences(game.objectState)

	var tick Tick
	var delta Pos

	switch msg.Data {
	case "up":
		delta = Pos{0, -1}
		fallthrough
	case "down":
		if msg.Data == "down" {
			delta = Pos{0, 1}
		}
		fallthrough
	case "right":
		if msg.Data == "right" {
			delta = Pos{1, 0}
		}
		fallthrough
	case "left":
		if msg.Data == "left" {
			delta = Pos{-1, 0}
		}
		for _, e := range meanings {
			for _, mod := range e.modifier {
				if mod == strconv.Itoa(msg.player) {
					// move all e.affected's
					for _, toMoveItem := range e.affected {
						// toMoveItem is an "Item"
						objects := game.objectState.findCharByItem(toMoveItem)
						for _, toMoveObject := range objects {
							// Add a new tick to move!
							tick = append(tick, Change{
								Event: Move,
								Id:    toMoveObject.Id,
								Pos:   delta,
							})
						}
					}
				}
			}
		}

		game.AddTick(tick)
		break
	}
}

func (game *Game) FindId(id Id) *Object {
	for _, e := range game.objectState {
		if e.Id == id {
			return e
		}
	}
	return &Object{} // Shouldn't happen
}

func (game *Game) DoChange(change Change) {
	switch change.Event {
	case Move:
		obj := game.FindId(change.Id)
		obj.Pos.X += change.Pos.X
		obj.Pos.Y += change.Pos.Y
		break
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
