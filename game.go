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

func (pos Pos) Neg() Pos {
	return Pos{
		-pos.X,
		-pos.Y,
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
	Objects []Object `json:"objects"`
	Height  int      `json:"height"`
	Width   int      `json:"width"`
	Number  int
}

type Event int

const (
	Move  Event = 0
	Death Event = 1
	Spawn Event = 2
	// Transform: death+spawn
)

type Change struct {
	Event Event  `json:"event"`
	Id    Id     `json:"id"`
	Pos   Pos    `json:"pos"` // Move: deltapos
	Item  string `json:"item"`
	Kind  Kind   `json:"kind"`
}

type Tick []Change

type Timeline []Tick

type Game struct {
	level       Level
	Width       int        `json:"width"`
	Height      int        `json:"height"`
	ObjectState ObjectList `json:"objects"`
	timeline    Timeline
	updateChan  chan<- Message
}

func (objs ObjectList) FindCharsByItem(item string) ObjectList {
	outs := ObjectList{}
	for _, e := range objs {
		if e.Kind == Char && e.Item == item {
			outs = append(outs, e)
		}
	}
	return outs
}

func MakeGame(level Level, updateChan chan<- Message) *Game {
	game := &Game{Level{}, 0, 0, ObjectList{}, Timeline{}, updateChan}
	game.SetLevel(level)
	game.timeline = Timeline{}
	return game
}

func (game *Game) SetLevel(level Level) {
	game.ObjectState = ObjectList{}
	game.level = level
	game.Width = level.Width
	game.Height = level.Height
	game.timeline = Timeline{}
	for i := range level.Objects {
		newObj := level.Objects[i]
		game.ObjectState = append(game.ObjectState, &newObj)
	}

	game.EmitState()
}

func (game *Game) ReceiveData(msg ReceivedMessage) {
	meanings := findSentences(game.ObjectState)
	
	isSpectator := msg.player > 2
	
	if isSpectator {
		return
	}

	var tick Tick
	var delta Pos

	switch msg.Data {
	case "restart":
		game.SetLevel(game.level)
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

		objectsToMove := ObjectList{}

		somethingMoved := false
		for affectedsKey := range meanings {
			for modifiersKey := range meanings[affectedsKey] {
				if modifiersKey == strconv.Itoa(msg.player) {
					objectsToMove = append(objectsToMove, game.ObjectState.FindCharsByItem(affectedsKey)...)
				}
			}
		}

		// sort tiles based on delta
		sortedObjects := objectsToMove.SortTiles(delta)
		for _, toMoveObject := range sortedObjects {
			success := game.CheckCollision(delta, toMoveObject, meanings, &tick)
			if success {
				somethingMoved = true
			}
		}

		if somethingMoved {
			game.AddTick(tick)
		}
	case "undo":
		game.Undo()
	}
}

// Sort according to move direction
func (objects ObjectList) SortTiles(delta Pos) ObjectList {
	// 1, 0 = start looking from right
	// -1, 0 = start looking from left
	// 0, 1 = start looking from down
	// 0, -1 = start looking from up
	if delta.X > 0 {
		sort.Slice(objects[:], func(i, j int) bool { return objects[i].Pos.X > objects[j].Pos.X })
	} else if delta.X < 0 {
		sort.Slice(objects[:], func(i, j int) bool { return objects[i].Pos.X < objects[j].Pos.X })
	} else if delta.Y > 0 {
		sort.Slice(objects[:], func(i, j int) bool { return objects[i].Pos.Y > objects[j].Pos.Y })
	} else if delta.X < 0 {
		sort.Slice(objects[:], func(i, j int) bool { return objects[i].Pos.Y < objects[j].Pos.Y })
	}
	return objects
}

func (game *Game) CheckCollision(delta Pos, objectToMove *Object, meanings Meanings, tick *Tick) (success bool) {
	success = true

	if objectToMove.CheckOutOfBounds(game, delta) {
		// Would be out of bounds
		success = false
	} else {
		for _, atPos := range game.ObjectState.FindObjectsAtPos(objectToMove.Pos.Add(delta)) {
			meaningsMap := meanings[atPos.Item]

			switch {
			case atPos.Kind == Char && objectToMove.Kind == Char && ((meaningsMap["1"] && meanings[objectToMove.Item]["1"]) || (meaningsMap["2"] && meanings[objectToMove.Item]["2"])):
			// Do nothing
			case meaningsMap["push"] || atPos.Kind != Char: // Text blocks are by default "push"
				if !game.CheckCollision(delta, atPos, meanings, tick) {
					success = false
				}
			case meaningsMap["stop"]:
				success = false
			case meaningsMap["defeat"]:
				// Success true
			default:
			}
		}
	}

	if success {
		// Add a new tick to move!
		change := Change{
			Event: Move,
			Id:    objectToMove.Id,
			Pos:   delta,
		}
		*tick = append(*tick, change)
		game.DoChange(change)
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

func (objectList ObjectList) FindId(id Id) *Object {
	for _, e := range objectList {
		if e.Id == id {
			return e
		}
	}
	return &Object{} // Shouldn't happen
}

func (game *Game) RemoveById(id Id) {
	for i := range game.ObjectState {
		if game.ObjectState[i].Id == id {
			game.ObjectState = append(game.ObjectState[:i], game.ObjectState[i+1:]...)
			return
		}
	}
}

func (game *Game) AddObject(id Id, kind Kind, pos Pos, item string) {
	newObject := &Object{
		Id:   id,
		Pos:  pos,
		Kind: kind,
		Item: item,
	}
	game.ObjectState = append(game.ObjectState, newObject)
}

func (game *Game) DoChange(change Change) {
	switch change.Event {
	case Move:
		obj := game.ObjectState.FindId(change.Id)
		obj.Pos.X += change.Pos.X
		obj.Pos.Y += change.Pos.Y
	case Death:
		game.RemoveById(change.Id)
	case Spawn:
		game.AddObject(change.Id, change.Kind, change.Pos, change.Item)
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
	if len(tick) > 0 {
		game.updateChan <- Message{Data: tick, MsgType: 1}
	}
}

func (game *Game) EmitState() {
	game.updateChan <- Message{Data: game, MsgType: 0}
}

func (game *Game) CheckWins(tick *Tick) {
	blocks := []string{
		"bemi",
		"fish",
		"sami",
		"soup",
		"star",
		"wall",
	}
	isWinning := false
	meanings := findSentences(game.ObjectState)
	toCheck := []string{}
	for p := range meanings {
		if meanings[p]["1"] || meanings[p]["2"] {
			toCheck = append(toCheck, p)
		}
		for _, blck := range blocks {
			if meanings[p][blck] {
				blocksToChange := game.ObjectState.FindCharsByItem(p)
				for _, e := range blocksToChange {
					change := Change{
						Event: Death,
						Id:    e.Id,
						Pos:   e.Pos,
						Item:  p,
					}
					change2 := Change{
						Event: Spawn,
						Id:    e.Id,
						Pos:   e.Pos,
						Item:  blck,
					}
					game.DoChange(change)
					game.DoChange(change2)
					*tick = append(*tick, change, change2)
				}
			}
		}
	}
	for _, item := range toCheck {
		chars := game.ObjectState.FindCharsByItem(item)
		for _, char := range chars {
			colliding := game.ObjectState.FindObjectsAtPos(char.Pos)
			for _, colObj := range colliding {
				if meanings[colObj.Item]["win"] {
					isWinning = true
				}
				if meanings[colObj.Item]["defeat"] {
					change := Change{Event: Death, Id: char.Id, Pos: char.Pos, Item: char.Item}
					game.DoChange(change)
					*tick = append(*tick, change)
				}
			}
		}
	}

	if isWinning {
		game.SetLevel(getLevel(game.level.Number + 1))
		*tick = Tick{}
	}
}

// Not in use
func (game *Game) DoTick(tick Tick) {
	for _, change := range tick {
		game.DoChange(change)
	}
}

func (game *Game) AddTick(tick Tick) {
	game.CheckWins(&tick)
	game.timeline = append(game.timeline, tick)
	game.EmitDelta(tick)
}

func (game *Game) Undo() {
	length := len(game.timeline)
	if length > 0 {
		lastTick := game.timeline[length-1]
		game.timeline = game.timeline[:length-1] // Remove last tick

		for i, e := range lastTick {
			// Invert every tick to do the opposite
			switch e.Event {
			case Move:
				e.Pos = e.Pos.Neg()
			case Death:
				e.Event = Spawn
			case Spawn:
				e.Event = Death
			}
			lastTick[i] = e
		}

		// Invert lastTick
		for i, j := 0, len(lastTick)-1; i < j; i, j = i+1, j-1 {
			lastTick[i], lastTick[j] = lastTick[j], lastTick[i]
		}

		for _, t := range lastTick {
			game.DoChange(t)
		}

		game.EmitDelta(lastTick)
	}
}
