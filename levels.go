package main

import "fmt"

func getLevel(level int) Level {
	if level == 0 {
		return Level{
			Width:  10,
			Height: 10,
			Number: 0,
			Objects: []Object{
				addObject(Pos{3, 8}, Char, "bemi"),
				addObject(Pos{7, 8}, Char, "sami"),

				addObject(Pos{0, 0}, Noun, "bemi"),
				addObject(Pos{1, 0}, Conj, "is"),
				addObject(Pos{2, 0}, Adj, "1"),

				addObject(Pos{0, 1}, Noun, "sami"),
				addObject(Pos{1, 1}, Conj, "is"),
				addObject(Pos{2, 1}, Adj, "2"),

				addObject(Pos{0, 5}, Noun, "star"),
				addObject(Pos{0, 6}, Conj, "is"),
				addObject(Pos{0, 7}, Adj, "win"),
				addObject(Pos{8, 1}, Char, "star"),
			},
		}
	} else if level == 1 {
		return Level{
			Width:  10,
			Height: 10,
			Number: 1,
			Objects: []Object{
				addObject(Pos{3, 8}, Char, "bemi"),
				addObject(Pos{7, 8}, Char, "sami"),

				addObject(Pos{0, 3}, Char, "wall"),
				addObject(Pos{1, 3}, Char, "wall"),
				addObject(Pos{2, 3}, Char, "wall"),
				addObject(Pos{3, 3}, Char, "wall"),
				addObject(Pos{4, 3}, Char, "wall"),
				addObject(Pos{5, 3}, Char, "wall"),
				addObject(Pos{6, 3}, Char, "wall"),
				addObject(Pos{7, 3}, Char, "wall"),
				addObject(Pos{8, 3}, Char, "wall"),
				addObject(Pos{9, 3}, Char, "wall"),

				addObject(Pos{0, 0}, Noun, "bemi"),
				addObject(Pos{1, 0}, Conj, "is"),
				addObject(Pos{2, 0}, Adj, "1"),

				addObject(Pos{0, 1}, Noun, "sami"),
				addObject(Pos{1, 1}, Conj, "is"),
				addObject(Pos{2, 1}, Adj, "2"),

				addObject(Pos{0, 5}, Noun, "star"),
				addObject(Pos{0, 6}, Conj, "is"),
				addObject(Pos{0, 7}, Adj, "win"),
				addObject(Pos{8, 1}, Char, "star"),
			},
		}
	} else if level == 2 {
		return Level{
			Width:  18,
			Height: 12,
			Number: 2,
			Objects: []Object{
				addObject(Pos{2, 10}, Char, "bemi"),
				addObject(Pos{16, 10}, Char, "sami"),

				addObject(Pos{9, 0}, Char, "wall"),
				addObject(Pos{9, 1}, Char, "wall"),
				addObject(Pos{9, 2}, Char, "wall"),
				addObject(Pos{9, 3}, Char, "wall"),
				addObject(Pos{9, 4}, Char, "wall"),
				addObject(Pos{9, 5}, Char, "wall"),
				addObject(Pos{9, 6}, Char, "wall"),
				addObject(Pos{9, 7}, Char, "wall"),
				addObject(Pos{9, 8}, Char, "wall"),
				addObject(Pos{9, 9}, Char, "wall"),
				addObject(Pos{9, 10}, Char, "wall"),
				addObject(Pos{9, 11}, Char, "wall"),

				addObject(Pos{0, 0}, Noun, "wall"),
				addObject(Pos{1, 0}, Conj, "is"),
				addObject(Pos{2, 0}, Adj, "stop"),

				addObject(Pos{0, 1}, Noun, "bemi"),
				addObject(Pos{1, 1}, Conj, "is"),
				addObject(Pos{2, 1}, Adj, "1"),

				addObject(Pos{0, 5}, Noun, "sami"),
				addObject(Pos{1, 5}, Conj, "is"),
				addObject(Pos{3, 5}, Adj, "2"),

				addObject(Pos{15, 0}, Noun, "star"),
				addObject(Pos{16, 0}, Conj, "is"),
				addObject(Pos{17, 0}, Adj, "win"),

				addObject(Pos{11, 1}, Char, "star"),
				addObject(Pos{17, 5}, Noun, "star"),
				addObject(Pos{17, 6}, Conj, "is"),
				addObject(Pos{17, 7}, Adj, "win"),
			},
		}
	} else if level == 3 {
		return Level{
			Width:  18,
			Height: 12,
			Number: 3,
			Objects: []Object{
				addObject(Pos{2, 10}, Char, "fish"),
				addObject(Pos{16, 10}, Char, "soup"),

				addObject(Pos{9, 0}, Char, "bemi"),
				addObject(Pos{9, 1}, Char, "bemi"),
				addObject(Pos{9, 2}, Char, "bemi"),
				addObject(Pos{9, 3}, Char, "bemi"),
				addObject(Pos{9, 4}, Char, "bemi"),
				addObject(Pos{9, 5}, Char, "bemi"),
				addObject(Pos{9, 6}, Char, "bemi"),
				addObject(Pos{9, 7}, Char, "bemi"),
				addObject(Pos{9, 8}, Char, "bemi"),
				addObject(Pos{9, 9}, Char, "bemi"),
				addObject(Pos{9, 10}, Char, "bemi"),
				addObject(Pos{9, 11}, Char, "bemi"),

				addObject(Pos{0, 0}, Noun, "bemi"),
				addObject(Pos{1, 0}, Conj, "is"),
				addObject(Pos{2, 0}, Adj, "stop"),

				addObject(Pos{0, 2}, Noun, "fish"),
				addObject(Pos{0, 3}, Conj, "is"),
				addObject(Pos{0, 4}, Adj, "1"),

				addObject(Pos{0, 6}, Noun, "soup"),
				addObject(Pos{1, 6}, Conj, "is"),
				addObject(Pos{3, 5}, Adj, "2"),

				addObject(Pos{11, 1}, Char, "sami"),
				addObject(Pos{17, 5}, Noun, "sami"),
				addObject(Pos{17, 6}, Conj, "is"),
				addObject(Pos{17, 7}, Adj, "win"),
			},
		}

	} else if level == 4 {
		return Level{
			Width:  16,
			Height: 12,
			Number: 4,
			Objects: []Object{
				addObject(Pos{7, 2}, Char, "bemi"),
				addObject(Pos{12, 10}, Char, "sami"),

				addObject(Pos{0, 2}, Noun, "wall"),
				addObject(Pos{1, 2}, Conj, "is"),
				addObject(Pos{2, 2}, Adj, "stop"),

				addObject(Pos{0, 0}, Noun, "fish"),
				addObject(Pos{1, 0}, Conj, "is"),
				addObject(Pos{2, 0}, Adj, "stop"),

				addObject(Pos{6, 1}, Char, "wall"),
				addObject(Pos{7, 1}, Char, "wall"),
				addObject(Pos{8, 1}, Char, "wall"),
				addObject(Pos{8, 2}, Char, "wall"),
				addObject(Pos{8, 3}, Char, "wall"),
				addObject(Pos{7, 3}, Char, "wall"),
				addObject(Pos{6, 3}, Char, "wall"),
				addObject(Pos{6, 2}, Char, "wall"),

				addObject(Pos{5, 0}, Char, "fish"),
				addObject(Pos{6, 0}, Char, "fish"),
				addObject(Pos{7, 0}, Char, "fish"),
				addObject(Pos{8, 0}, Char, "fish"),
				addObject(Pos{9, 0}, Char, "fish"),
				addObject(Pos{9, 1}, Char, "fish"),
				addObject(Pos{9, 2}, Char, "fish"),
				addObject(Pos{9, 3}, Char, "fish"),
				addObject(Pos{9, 4}, Char, "fish"),
				addObject(Pos{8, 4}, Char, "fish"),
				addObject(Pos{7, 4}, Char, "fish"),
				addObject(Pos{6, 4}, Char, "fish"),
				addObject(Pos{5, 4}, Char, "fish"),
				addObject(Pos{5, 3}, Char, "fish"),
				addObject(Pos{5, 2}, Char, "fish"),
				addObject(Pos{5, 1}, Char, "fish"),

				addObject(Pos{2, 8}, Noun, "bemi"),
				addObject(Pos{3, 8}, Conj, "is"),
				addObject(Pos{4, 8}, Adj, "1"),

				addObject(Pos{0, 6}, Noun, "sami"),
				addObject(Pos{1, 6}, Conj, "is"),
				addObject(Pos{2, 6}, Adj, "2"),

				addObject(Pos{12, 2}, Noun, "star"),
				addObject(Pos{12, 3}, Conj, "is"),
				addObject(Pos{12, 4}, Adj, "win"),
			},
		}
	} else if level == 5 {
		return Level{
			Width:  10,
			Height: 10,
			Number: 5,
			Objects: []Object{

				addObject(Pos{7, 2}, Char, "wall"),
				addObject(Pos{8, 8}, Char, "sami"),

				addObject(Pos{0, 2}, Noun, "wall"),
				addObject(Pos{1, 2}, Conj, "is"),
				addObject(Pos{2, 2}, Adj, "stop"),

				addObject(Pos{0, 0}, Noun, "fish"),
				addObject(Pos{1, 0}, Conj, "is"),
				addObject(Pos{2, 0}, Adj, "stop"),

				addObject(Pos{5, 0}, Char, "fish"),
				addObject(Pos{6, 0}, Char, "fish"),
				addObject(Pos{7, 0}, Char, "fish"),
				addObject(Pos{8, 0}, Char, "fish"),
				addObject(Pos{9, 0}, Char, "fish"),
				addObject(Pos{9, 1}, Char, "fish"),
				addObject(Pos{9, 2}, Char, "fish"),
				addObject(Pos{9, 3}, Char, "fish"),
				addObject(Pos{9, 4}, Char, "fish"),
				addObject(Pos{8, 4}, Char, "fish"),
				addObject(Pos{7, 4}, Char, "fish"),
				addObject(Pos{6, 4}, Char, "fish"),
				addObject(Pos{5, 4}, Char, "fish"),
				addObject(Pos{5, 3}, Char, "fish"),
				addObject(Pos{5, 2}, Char, "fish"),
				addObject(Pos{5, 1}, Char, "fish"),

				addObject(Pos{2, 8}, Noun, "fish"),
				addObject(Pos{3, 8}, Conj, "is"),
				addObject(Pos{0, 9}, Adj, "1"),

				addObject(Pos{0, 6}, Noun, "sami"),
				addObject(Pos{1, 6}, Conj, "is"),
				addObject(Pos{2, 6}, Adj, "2"),

				addObject(Pos{8, 5}, Noun, "star"),
				addObject(Pos{8, 6}, Conj, "is"),
				addObject(Pos{8, 7}, Adj, "win"),
			},
		}
	} else if level == 6 {
		return Level{
			Width:  10,
			Height: 10,
			Number: 6,
			Objects: []Object{

				addObject(Pos{7, 2}, Adj, "win"),
				addObject(Pos{8, 8}, Char, "sami"),

				addObject(Pos{0, 2}, Noun, "fish"),
				addObject(Pos{1, 2}, Conj, "is"),
				addObject(Pos{2, 2}, Adj, "stop"),

				addObject(Pos{0, 0}, Noun, "wall"),
				addObject(Pos{1, 0}, Conj, "is"),
				addObject(Pos{2, 0}, Adj, "stop"),

				addObject(Pos{5, 0}, Char, "fish"),
				addObject(Pos{6, 0}, Char, "fish"),
				addObject(Pos{7, 0}, Char, "fish"),
				addObject(Pos{8, 0}, Char, "fish"),
				addObject(Pos{9, 0}, Char, "fish"),
				addObject(Pos{9, 1}, Char, "fish"),
				addObject(Pos{9, 2}, Char, "fish"),
				addObject(Pos{9, 3}, Char, "fish"),
				addObject(Pos{9, 4}, Char, "fish"),
				addObject(Pos{8, 4}, Char, "fish"),
				addObject(Pos{7, 4}, Char, "fish"),
				addObject(Pos{6, 4}, Char, "fish"),
				addObject(Pos{5, 4}, Char, "fish"),
				addObject(Pos{5, 3}, Char, "fish"),
				addObject(Pos{5, 2}, Char, "fish"),
				addObject(Pos{5, 1}, Char, "fish"),

				addObject(Pos{2, 8}, Noun, "bemi"),
				addObject(Pos{3, 8}, Conj, "is"),
				addObject(Pos{0, 9}, Adj, "2"),
				addObject(Pos{8, 9}, Char, "bemi"),

				addObject(Pos{0, 6}, Noun, "sami"),
				addObject(Pos{1, 6}, Conj, "is"),
				addObject(Pos{2, 6}, Adj, "1"),

				addObject(Pos{8, 5}, Adj, "push"),
			},
		}
	}

	//default level for testing and what not
	return Level{
		Width:  1,
		Height: 1,
		Objects: []Object{
			addObject(Pos{2, 10}, Char, "bemi"),
			addObject(Pos{16, 10}, Char, "sami"),

			addObject(Pos{9, 5}, Char, "wall"),

			addObject(Pos{0, 0}, Noun, "star"),
			addObject(Pos{1, 0}, Conj, "is"),
			addObject(Pos{2, 0}, Adj, "defeat"),

			addObject(Pos{0, 2}, Noun, "bemi"),
			addObject(Pos{1, 2}, Conj, "is"),
			addObject(Pos{2, 2}, Adj, "1"),

			addObject(Pos{0, 5}, Noun, "sami"),
			addObject(Pos{1, 5}, Conj, "is"),
			addObject(Pos{3, 5}, Adj, "2"),
			addObject(Pos{11, 1}, Char, "star")},
	}
}

func addObject(pos Pos, kind Kind, item string) Object {
	if kind == Conj && item != "is" && item != "and" {
		fmt.Println(item, "is not a conjugation.", pos)
	}
	obj := Object{
		pos,
		kind,
		item,
		newId(),
	}
	return obj
}

func newId() Id {
	oldId := idCounter
	idCounter++
	return oldId
}
