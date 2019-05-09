package main

import (
	"fmt"
	"testing"
)

func TestSingleIsSingleNounSentence(t *testing.T) {
	bemiBlock := Object{
		Pos:  Pos{10, 10},
		Kind: Noun,
		Item: "bemi",
		Id:   666,
	}
	isBlock := Object{
		Pos:  Pos{10, 15},
		Kind: Conj,
		Item: "is",
		Id:   676,
	}
	wallBlock := Object{
		Pos:  Pos{13, 10},
		Kind: Noun,
		Item: "wall",
		Id:   667,
	}

	m := Meaning{
		affected: []string{bemiBlock.Item},
		modifier: []string{wallBlock.Item},
	}

	output := parseSentence([]Object{bemiBlock, isBlock, wallBlock})

	if len(output.affected) == len(m.affected) && len(output.modifier) == len(m.modifier) && output.affected[0] == m.affected[0] && output.modifier[0] == m.modifier[0] {
	} else {
		t.Errorf("Parse is incorrect")
	}
}

func TestParseSingleIsSingleNounSentence(t *testing.T) {
	bemiBlock := Object{
		Pos:  Pos{10, 10},
		Kind: Noun,
		Item: "bemi",
		Id:   5,
	}
	isBlock := Object{
		Pos:  Pos{11, 10},
		Kind: Conj,
		Item: "is",
		Id:   6,
	}
	wallBlock := Object{
		Pos:  Pos{12, 10},
		Kind: Noun,
		Item: "wall",
		Id:   7,
	}

	m := Meaning{
		affected: []string{bemiBlock.Item},
		modifier: []string{wallBlock.Item},
	}

	output := newWording([]Object{bemiBlock, isBlock, wallBlock})
	fmt.Println(output)
	if len(output) != 1 {
		t.Errorf("Wording returns incorrect amount")
	} else if len(output[0].affected) != len(m.affected) {
		t.Errorf("Wrong affected length")
	}
}
