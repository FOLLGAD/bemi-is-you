package main

import (
	"sort"
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

type ObjectList []*Object

type Level struct {
	Objects ObjectList `json:"objects"`
	Height  int        `json:"height"`
	Width   int        `json:"width"`
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
type Game struct {
	level       Level
	objectState ObjectList
	timeline    Timeline
	updateChan  chan<- Tick
}

func (objs ObjectList) FindCharByItem(item string) ObjectList {
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

	case "restart":
		game.objectState = game.level.Objects //bruh?
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
		for affectedsKey := range meanings {
			for modifiersKey := range meanings[affectedsKey] {
				if modifiersKey == strconv.Itoa(msg.player) {
					objects := game.objectState.FindCharByItem(affectedsKey)
					// sort tiles based on delta
					sortedObjects := objects.SortTiles(delta)
					for _, toMoveObject := range sortedObjects {
						game.CheckCollision(delta, toMoveObject, meanings, &tick)
					}
				}
			}
		}

		game.AddTick(tick)
		break
	}
}

// Sort according to move direction
func (objects ObjectList) SortTiles(delta Pos) ObjectList {
	// 1, 0 = start looking from right
	// -1, 0 = start looking from left
	// 0, 1 = start looking from down
	// 0, -1 = start looking from up
	if delta.X == 1 {
		sort.Slice(objects, func(i, j int) bool { return objects[i].Pos.X > objects[j].Pos.X })
	} else if delta.X == -1 {
		sort.Slice(objects, func(i, j int) bool { return objects[i].Pos.X < objects[j].Pos.X })
	} else if delta.Y == 1 {
		sort.Slice(objects, func(i, j int) bool { return objects[i].Pos.Y > objects[j].Pos.Y })
	} else if delta.X == -1 {
		sort.Slice(objects, func(i, j int) bool { return objects[i].Pos.Y < objects[j].Pos.Y })
	}
	return objects
}

func (game *Game) CheckCollision(delta Pos, objectToMove *Object, meanings Meanings, tick *Tick) (success bool) {
	success = true

	for _, atPos := range game.objectState.FindObjectsAtPos(objectToMove.Pos.Add(delta)) {
		meaningsMap := meanings[atPos.Item]

		switch {
		case meaningsMap["push"]:
			if !game.CheckCollision(delta, atPos, meanings, tick) {
				success = false
			}
		case meaningsMap["stop"]:

		case meaningsMap["defeat"]:

		default:
		}
	}

	if success {
		// Add a new tick to move!
		change := Change{
			Event: Move,
			Id:    objectToMove.Id,
			Pos:   delta,
		}

		if !objectToMove.CheckOutOfBounds(game, delta) {
			*tick = append(*tick, change)
			game.DoChange(change)
		} else {
			success = false
		}
	}

	return success
}

func (objects ObjectList) FindObjectsAtPos(pos Pos) ObjectList {
	objectsAtPos := ObjectList{}
	for _, e := range objects {
		if e.Pos == pos {
			objectsAtPos = append(objectsAtPos, e)
		}
	}
	return objectsAtPos
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

func (obj Object) CheckOutOfBounds(game *Game, delta Pos) bool {
	newPos := obj.Pos.Add(delta)
	if newPos.X >= game.level.Width || newPos.X < 0 || newPos.Y >= game.level.Height || newPos.Y < 0 {
		return true
	}
	return false
}
func (game *Game) EmitDelta(tick Tick) {
	game.updateChan <- tick
}

// Not in use
func (game *Game) DoTick(tick Tick) {
	for _, change := range tick {
		game.DoChange(change)
	}
}

func (game *Game) AddTick(tick Tick) {
	game.timeline = append(game.timeline, tick)
	game.EmitDelta(tick)
}

func (game *Game) Undo() {
	length := len(game.timeline)
	lastTick := game.timeline[length-1]
	game.timeline = game.timeline[:len(game.timeline)-2] // Remove last tick

	for i := len(lastTick) - 1; i >= 0; i-- {
		game.DoChange(lastTick[i])
	}
}
