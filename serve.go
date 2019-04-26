package main

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

}
