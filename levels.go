package main

func getLevel(level int) Level {
	if level == 1 {
		return Level{
			Width:  18,
			Height: 12,
			Objects: []*Object{
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
				addObject(Pos{9, 12}, Char, "wall"),

				addObject(Pos{0, 0}, Noun, "wall"),
				addObject(Pos{1, 0}, Conj, "is"),
				addObject(Pos{2, 0}, Adj, "stop"),

				addObject(Pos{0, 1}, Noun, "bemi"),
				addObject(Pos{1, 1}, Conj, "is"),
				addObject(Pos{2, 1}, Conj, "1"),

				addObject(Pos{0, 5}, Noun, "sami"),
				addObject(Pos{1, 5}, Conj, "is"),
				addObject(Pos{3, 5}, Adj, "2"),
				addObject(Pos{11, 1}, Char, "star"),
			},
		}
	} else if level == 2 {
		return Level{
			Width:  18,
			Height: 12,
			Objects: []*Object{
				addObject(Pos{2, 10}, Char, "bemi"),
				addObject(Pos{16, 10}, Char, "sami"),

				addObject(Pos{9, 0}, Char, "wall"),
				addObject(Pos{9, 1}, Char, "wall"),
				addObject(Pos{9, 2}, Char, "wall"),
				addObject(Pos{9, 4}, Char, "wall"),
				addObject(Pos{9, 5}, Char, "wall"),
				addObject(Pos{9, 6}, Char, "wall"),
				addObject(Pos{9, 7}, Char, "wall"),
				addObject(Pos{9, 8}, Char, "wall"),
				addObject(Pos{9, 9}, Char, "wall"),
				addObject(Pos{9, 10}, Char, "wall"),
				addObject(Pos{9, 11}, Char, "wall"),
				addObject(Pos{9, 12}, Char, "wall"),

				addObject(Pos{0, 0}, Noun, "wall"),
				addObject(Pos{1, 0}, Conj, "is"),
				addObject(Pos{2, 0}, Adj, "stop"),

				addObject(Pos{0, 1}, Noun, "bemi"),
				addObject(Pos{1, 1}, Conj, "is"),
				addObject(Pos{2, 1}, Conj, "1"),

				addObject(Pos{0, 5}, Noun, "sami"),
				addObject(Pos{1, 5}, Conj, "is"),
				addObject(Pos{3, 5}, Adj, "2"),
				addObject(Pos{11, 1}, Char, "star"),
			},
		}
	}
	return Level{
		Width:   1,
		Height:  1,
		Objects: []*Object{},
	}
}

func addObject(pos Pos, kind Kind, item string) *Object {
	obj := &Object{
		pos,
		kind,
		item,
		idCounter,
	}
	idCounter++
	return obj
}
