package main

import (
	"testing"
)

func TestCollision(t *testing.T) {
	var firstLevel = Level{}
	firstLevel.Height = 20
	firstLevel.Width = 25

	sami := addObject(Pos{2, 6}, Char, "sami")
	movedObject := addObject(Pos{2, 5}, Adj, "1")
	firstLevel.Objects = []*Object{
		addObject(Pos{10, 10}, Char, "bemi"),
		sami,
		addObject(Pos{0, 0}, Noun, "bemi"),
		addObject(Pos{1, 0}, Conj, "is"),
		addObject(Pos{2, 0}, Adj, "1"),
		addObject(Pos{0, 5}, Noun, "sami"),
		addObject(Pos{1, 5}, Conj, "is"),
		movedObject,
		addObject(Pos{0, 6}, Conj, "is"),
		addObject(Pos{0, 7}, Adj, "2"),
	}

	ch := make(chan Tick)
	game := Game{firstLevel, firstLevel.Objects, Timeline{}, ch}
	tick := &Tick{}
	meanings := Meanings{
		"sami": {"1": true},
		"bemi": {"2": true},
	}

	go func() {
		for {
			<-ch
		}
	}()

	game.CheckCollision(Pos{0, -1}, sami, meanings, tick)

	game.AddTick(*tick)

	expectedPos := Pos{2, 4}

	if movedObject.Pos != expectedPos {
		t.Error("Expected:", expectedPos, "- Actual:", movedObject.Pos)
	}
}

func TestSortTilesRight(t *testing.T) {
	objectlist := ObjectList{
		addObject(Pos{3, 4}, Char, "bemi"),
		addObject(Pos{3, 5}, Char, "sami"),
	}
	sorted := objectlist.SortTiles(Pos{0, 1})
	if sorted[0].Item != "sami" {
		t.Error("Expected first index to be sami, got", sorted[0].Item)
	} else if sorted[1].Item != "bemi" {
		t.Error("Expected second index to be bemi, got", sorted[1].Item)
	}
}
