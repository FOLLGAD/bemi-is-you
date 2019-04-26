package main

type Pos struct {
	x int
	y int
}

type Object struct {
	pos   Pos
	color string
	id    string
}

type Level struct {
	objects []Object
	width   int
	height  int
}

func main() {

}
