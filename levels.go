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
				addObject(Pos{2, 1}, Adj, "2"),

				addObject(Pos{0, 5}, Noun, "sami"),
				addObject(Pos{1, 5}, Conj, "is"),
				addObject(Pos{2, 6}, Adj, "1"),

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

				addObject(Pos{0, 1}, Noun, "fish"),
				addObject(Pos{1, 1}, Conj, "is"),
				addObject(Pos{2, 1}, Adj, "1"),

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
				addObject(Pos{4, 8}, Adj, "2"),

				addObject(Pos{0, 6}, Noun, "sami"),
				addObject(Pos{1, 6}, Conj, "is"),
				addObject(Pos{2, 6}, Adj, "1"),

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

				addObject(Pos{1, 7}, Noun, "fish"),
				addObject(Pos{1, 8}, Conj, "is"),
				addObject(Pos{0, 9}, Adj, "1"),

				addObject(Pos{0, 5}, Noun, "sami"),
				addObject(Pos{1, 5}, Conj, "is"),
				addObject(Pos{2, 5}, Adj, "2"),

				addObject(Pos{8, 7}, Adj, "win"),
			},
		}
	} else if level == 6 {
		return Level{
			Width:  18,
			Height: 12,
			Number: 6,
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

				addObject(Pos{0, 1}, Noun, "wall"),
				addObject(Pos{1, 1}, Conj, "is"),
				addObject(Pos{2, 1}, Adj, "stop"),

				addObject(Pos{2, 2}, Adj, "push"),

				addObject(Pos{14, 7}, Noun, "bemi"),
				addObject(Pos{14, 8}, Conj, "is"),
				addObject(Pos{14, 6}, Adj, "2"),

				addObject(Pos{0, 7}, Noun, "sami"),
				addObject(Pos{1, 7}, Conj, "is"),
				addObject(Pos{2, 7}, Adj, "1"),

				addObject(Pos{11, 0}, Char, "wall"),
				addObject(Pos{11, 1}, Char, "wall"),
				addObject(Pos{10, 1}, Char, "wall"),

				addObject(Pos{10, 0}, Char, "star"),
				addObject(Pos{17, 0}, Noun, "star"),
				addObject(Pos{17, 1}, Conj, "is"),
				addObject(Pos{17, 2}, Adj, "win"),
			},
		}
	} else if level == 7 {
		return Level{
			Width:  10,
			Height: 10,
			Number: 7,
			Objects: []Object{

				addObject(Pos{7, 2}, Adj, "win"),
				addObject(Pos{8, 8}, Char, "sami"),

				addObject(Pos{0, 1}, Noun, "fish"),
				addObject(Pos{1, 1}, Conj, "is"),
				addObject(Pos{2, 1}, Adj, "stop"),

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
				addObject(Pos{8, 6}, Conj, "is"),
			},
		}
	} else if level == 8 {
		return Level{
			Width:  18,
			Height: 12,
			Number: 8,
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

				addObject(Pos{10, 4}, Char, "wall"),
				addObject(Pos{11, 4}, Char, "wall"),
				addObject(Pos{12, 4}, Char, "wall"),
				addObject(Pos{13, 4}, Char, "wall"),
				addObject(Pos{14, 4}, Char, "wall"),
				addObject(Pos{15, 4}, Char, "wall"),
				addObject(Pos{16, 4}, Char, "wall"),
				addObject(Pos{17, 4}, Char, "wall"),

				addObject(Pos{0, 0}, Noun, "wall"),
				addObject(Pos{1, 0}, Conj, "is"),
				addObject(Pos{2, 0}, Adj, "defeat"),

				addObject(Pos{14, 7}, Noun, "bemi"),
				addObject(Pos{15, 8}, Conj, "is"),
				addObject(Pos{14, 6}, Adj, "2"),

				addObject(Pos{12, 6}, Adj, "push"),
				addObject(Pos{11, 6}, Conj, "is"),

				addObject(Pos{0, 7}, Noun, "sami"),
				addObject(Pos{0, 8}, Conj, "is"),
				addObject(Pos{0, 9}, Adj, "1"),

				addObject(Pos{10, 0}, Char, "star"),
				addObject(Pos{17, 0}, Noun, "star"),
				addObject(Pos{17, 1}, Conj, "is"),
				addObject(Pos{17, 2}, Adj, "win"),
			},
		}
	} else if level == 9 {
		return Level{
			Width:  18,
			Height: 15,
			Number: 9,
			Objects: []Object{

				addObject(Pos{11, 8}, Adj, "win"),

				addObject(Pos{10, 10}, Noun, "soup"),

				addObject(Pos{10, 1}, Noun, "star"),
				addObject(Pos{11, 1}, Conj, "is"),
				addObject(Pos{12, 1}, Adj, "defeat"),

				addObject(Pos{0, 1}, Noun, "wall"),
				addObject(Pos{6, 5}, Noun, "wall"),
				addObject(Pos{1, 1}, Conj, "is"),
				addObject(Pos{2, 1}, Adj, "stop"),

				addObject(Pos{5, 1}, Noun, "bemi"),
				addObject(Pos{6, 1}, Conj, "is"),
				addObject(Pos{7, 1}, Adj, "1"),
				addObject(Pos{5, 6}, Char, "bemi"),

				addObject(Pos{2, 4}, Char, "wall"),
				addObject(Pos{3, 4}, Char, "wall"),
				addObject(Pos{4, 4}, Char, "wall"),
				addObject(Pos{5, 4}, Char, "wall"),
				addObject(Pos{6, 4}, Char, "wall"),
				addObject(Pos{7, 4}, Char, "wall"),

				addObject(Pos{2, 5}, Char, "wall"),
				addObject(Pos{7, 5}, Char, "wall"),

				addObject(Pos{2, 6}, Char, "wall"),
				addObject(Pos{7, 6}, Char, "wall"),
				addObject(Pos{8, 6}, Char, "wall"),
				addObject(Pos{9, 6}, Char, "wall"),
				addObject(Pos{10, 6}, Char, "wall"),
				addObject(Pos{11, 6}, Char, "wall"),
				addObject(Pos{12, 6}, Char, "wall"),
				addObject(Pos{13, 6}, Char, "wall"),
				addObject(Pos{14, 6}, Char, "wall"),
				addObject(Pos{15, 6}, Char, "wall"),

				addObject(Pos{2, 7}, Char, "wall"),
				addObject(Pos{7, 7}, Char, "wall"),
				addObject(Pos{15, 7}, Char, "wall"),

				addObject(Pos{2, 8}, Char, "wall"),
				addObject(Pos{7, 8}, Char, "wall"),
				addObject(Pos{15, 8}, Char, "wall"),
				addObject(Pos{15, 9}, Char, "wall"),
				addObject(Pos{2, 9}, Char, "wall"),
				addObject(Pos{7, 9}, Char, "wall"),

				addObject(Pos{15, 10}, Char, "wall"),
				addObject(Pos{2, 10}, Char, "wall"),
				addObject(Pos{7, 10}, Char, "wall"),

				addObject(Pos{15, 11}, Char, "wall"),
				addObject(Pos{2, 11}, Char, "wall"),
				addObject(Pos{7, 11}, Char, "wall"),

				addObject(Pos{15, 12}, Char, "wall"),
				addObject(Pos{2, 12}, Char, "wall"),
				addObject(Pos{7, 12}, Char, "star"),

				addObject(Pos{2, 13}, Char, "wall"),
				addObject(Pos{2, 14}, Char, "wall"),
				addObject(Pos{3, 14}, Char, "wall"),
				addObject(Pos{4, 14}, Char, "wall"),
				addObject(Pos{5, 14}, Char, "wall"),
				addObject(Pos{6, 14}, Char, "wall"),
				addObject(Pos{7, 14}, Char, "wall"),
				addObject(Pos{7, 13}, Char, "wall"),
				addObject(Pos{8, 13}, Char, "wall"),
				addObject(Pos{9, 13}, Char, "wall"),
				addObject(Pos{10, 13}, Char, "wall"),
				addObject(Pos{11, 13}, Char, "wall"),
				addObject(Pos{12, 13}, Char, "wall"),
				addObject(Pos{13, 13}, Char, "wall"),
				addObject(Pos{14, 13}, Char, "wall"),
				addObject(Pos{15, 13}, Char, "wall"),

				addObject(Pos{14, 8}, Noun, "sami"),
				addObject(Pos{14, 9}, Conj, "is"),
				addObject(Pos{14, 10}, Adj, "2"),
				addObject(Pos{12, 9}, Char, "sami"),

				addObject(Pos{6, 6}, Conj, "is"),
			},
		}
	} else if level == 10 {
		return Level{
			Width:  18,
			Height: 15,
			Number: 10,
			Objects: []Object{

				addObject(Pos{11, 8}, Adj, "push"),
				addObject(Pos{8, 5}, Adj, "win"),

				addObject(Pos{10, 1}, Noun, "star"),
				addObject(Pos{11, 1}, Conj, "is"),
				addObject(Pos{12, 1}, Adj, "defeat"),

				addObject(Pos{0, 1}, Noun, "wall"),
				addObject(Pos{4, 5}, Noun, "wall"),
				addObject(Pos{1, 1}, Conj, "is"),
				addObject(Pos{2, 1}, Adj, "stop"),
				addObject(Pos{5, 5}, Conj, "is"),

				addObject(Pos{5, 1}, Noun, "bemi"),
				addObject(Pos{6, 1}, Conj, "is"),
				addObject(Pos{7, 1}, Adj, "1"),
				addObject(Pos{5, 6}, Char, "bemi"),

				addObject(Pos{2, 4}, Char, "wall"),
				addObject(Pos{3, 4}, Char, "wall"),
				addObject(Pos{4, 4}, Char, "wall"),
				addObject(Pos{5, 4}, Char, "wall"),
				addObject(Pos{6, 4}, Char, "wall"),
				addObject(Pos{7, 4}, Char, "wall"),

				addObject(Pos{2, 5}, Char, "wall"),
				addObject(Pos{7, 5}, Char, "wall"),

				addObject(Pos{2, 6}, Char, "wall"),
				addObject(Pos{7, 6}, Char, "wall"),
				addObject(Pos{8, 6}, Char, "wall"),
				addObject(Pos{9, 6}, Char, "wall"),
				addObject(Pos{10, 6}, Char, "wall"),
				addObject(Pos{11, 6}, Char, "wall"),
				addObject(Pos{12, 6}, Char, "wall"),
				addObject(Pos{13, 6}, Char, "wall"),
				addObject(Pos{14, 6}, Char, "wall"),
				addObject(Pos{15, 6}, Char, "wall"),

				addObject(Pos{2, 7}, Char, "wall"),
				addObject(Pos{7, 7}, Char, "wall"),
				addObject(Pos{15, 7}, Char, "wall"),

				addObject(Pos{2, 8}, Char, "wall"),
				addObject(Pos{7, 8}, Char, "wall"),
				addObject(Pos{15, 8}, Char, "wall"),
				addObject(Pos{15, 9}, Char, "wall"),
				addObject(Pos{2, 9}, Char, "wall"),
				addObject(Pos{7, 9}, Char, "wall"),

				addObject(Pos{15, 10}, Char, "wall"),
				addObject(Pos{2, 10}, Char, "wall"),
				addObject(Pos{7, 10}, Char, "wall"),

				addObject(Pos{15, 11}, Char, "wall"),
				addObject(Pos{2, 11}, Char, "wall"),
				addObject(Pos{7, 11}, Char, "wall"),

				addObject(Pos{15, 12}, Char, "wall"),
				addObject(Pos{2, 12}, Char, "wall"),
				addObject(Pos{7, 12}, Char, "star"),

				addObject(Pos{2, 13}, Char, "wall"),
				addObject(Pos{2, 14}, Char, "wall"),
				addObject(Pos{3, 14}, Char, "wall"),
				addObject(Pos{4, 14}, Char, "wall"),
				addObject(Pos{5, 14}, Char, "wall"),
				addObject(Pos{6, 14}, Char, "wall"),
				addObject(Pos{7, 14}, Char, "wall"),
				addObject(Pos{7, 13}, Char, "wall"),
				addObject(Pos{8, 13}, Char, "wall"),
				addObject(Pos{9, 13}, Char, "wall"),
				addObject(Pos{10, 13}, Char, "wall"),
				addObject(Pos{11, 13}, Char, "wall"),
				addObject(Pos{12, 13}, Char, "wall"),
				addObject(Pos{13, 13}, Char, "wall"),
				addObject(Pos{14, 13}, Char, "wall"),
				addObject(Pos{15, 13}, Char, "wall"),

				addObject(Pos{14, 8}, Noun, "sami"),
				addObject(Pos{14, 9}, Conj, "is"),
				addObject(Pos{14, 10}, Adj, "2"),
				addObject(Pos{12, 9}, Char, "sami"),
			},
		}
	} else if level == 11 {
		return Level{
			Width:  18,
			Height: 15,
			Number: 11,
			Objects: []Object{

				addObject(Pos{11, 3}, Adj, "push"),
				addObject(Pos{11, 4}, Conj, "is"),
				addObject(Pos{8, 14}, Adj, "win"),

				addObject(Pos{0, 2}, Char, "fish"),

				addObject(Pos{10, 1}, Noun, "star"),
				addObject(Pos{11, 1}, Conj, "is"),
				addObject(Pos{12, 1}, Adj, "defeat"),

				addObject(Pos{0, 1}, Noun, "wall"),
				addObject(Pos{4, 7}, Noun, "fish"),
				addObject(Pos{1, 1}, Conj, "is"),
				addObject(Pos{2, 1}, Adj, "stop"),

				addObject(Pos{5, 1}, Noun, "bemi"),
				addObject(Pos{6, 1}, Conj, "is"),
				addObject(Pos{7, 1}, Adj, "1"),
				addObject(Pos{5, 6}, Char, "bemi"),

				addObject(Pos{0, 0}, Char, "wall"),
				addObject(Pos{1, 0}, Char, "wall"),
				addObject(Pos{2, 0}, Char, "wall"),
				addObject(Pos{3, 0}, Char, "wall"),
				addObject(Pos{4, 0}, Char, "wall"),
				addObject(Pos{5, 0}, Char, "wall"),
				addObject(Pos{6, 0}, Char, "wall"),
				addObject(Pos{7, 0}, Char, "wall"),
				addObject(Pos{8, 0}, Char, "wall"),
				addObject(Pos{9, 0}, Char, "wall"),
				addObject(Pos{10, 0}, Char, "wall"),
				addObject(Pos{11, 0}, Char, "wall"),
				addObject(Pos{12, 0}, Char, "wall"),
				addObject(Pos{13, 0}, Char, "wall"),
				addObject(Pos{13, 1}, Char, "wall"),

				addObject(Pos{2, 4}, Char, "wall"),
				addObject(Pos{3, 4}, Char, "wall"),
				addObject(Pos{4, 4}, Char, "wall"),
				addObject(Pos{5, 4}, Char, "wall"),
				addObject(Pos{6, 4}, Char, "wall"),
				addObject(Pos{7, 4}, Char, "wall"),

				addObject(Pos{2, 5}, Char, "wall"),
				addObject(Pos{7, 5}, Char, "wall"),

				addObject(Pos{2, 6}, Char, "wall"),
				addObject(Pos{7, 6}, Char, "wall"),
				addObject(Pos{8, 6}, Char, "wall"),
				addObject(Pos{9, 6}, Char, "wall"),
				addObject(Pos{10, 6}, Char, "wall"),
				addObject(Pos{11, 6}, Char, "wall"),
				addObject(Pos{12, 6}, Char, "wall"),
				addObject(Pos{13, 6}, Char, "wall"),
				addObject(Pos{14, 6}, Char, "wall"),
				addObject(Pos{15, 6}, Char, "wall"),

				addObject(Pos{2, 7}, Char, "wall"),
				addObject(Pos{7, 7}, Char, "wall"),
				addObject(Pos{15, 7}, Char, "wall"),

				addObject(Pos{2, 8}, Char, "wall"),
				addObject(Pos{7, 8}, Char, "wall"),
				addObject(Pos{15, 8}, Char, "wall"),
				addObject(Pos{15, 9}, Char, "wall"),
				addObject(Pos{2, 9}, Char, "wall"),
				addObject(Pos{7, 9}, Char, "wall"),

				addObject(Pos{15, 10}, Char, "wall"),
				addObject(Pos{2, 10}, Char, "wall"),
				addObject(Pos{7, 10}, Char, "wall"),

				addObject(Pos{15, 11}, Char, "wall"),
				addObject(Pos{2, 11}, Char, "wall"),
				addObject(Pos{7, 11}, Char, "wall"),

				addObject(Pos{15, 12}, Char, "wall"),
				addObject(Pos{2, 12}, Char, "wall"),
				addObject(Pos{7, 12}, Char, "star"),

				addObject(Pos{2, 13}, Char, "wall"),
				addObject(Pos{3, 13}, Char, "wall"),
				addObject(Pos{5, 13}, Char, "wall"),
				addObject(Pos{6, 13}, Char, "wall"),
				addObject(Pos{3, 14}, Char, "wall"),
				addObject(Pos{4, 14}, Char, "wall"),
				addObject(Pos{5, 14}, Char, "wall"),
				addObject(Pos{7, 13}, Char, "wall"),
				addObject(Pos{8, 13}, Char, "wall"),
				addObject(Pos{9, 13}, Char, "wall"),
				addObject(Pos{10, 13}, Char, "wall"),
				addObject(Pos{11, 13}, Char, "wall"),
				addObject(Pos{12, 13}, Char, "wall"),
				addObject(Pos{13, 13}, Char, "wall"),
				addObject(Pos{14, 13}, Char, "wall"),
				addObject(Pos{15, 13}, Char, "wall"),

				addObject(Pos{11, 12}, Noun, "sami"),
				addObject(Pos{12, 12}, Conj, "is"),
				addObject(Pos{13, 12}, Adj, "2"),
				addObject(Pos{12, 9}, Char, "sami"),
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
