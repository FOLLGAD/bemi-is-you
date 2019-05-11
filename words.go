package main

import (
	"errors"
)

type Meaning struct {
	affected []string // Item
	modifier []string // The adjectives
}

func findSentences(objects ObjectList) []Meaning {
	isWords := ObjectList{}
	for _, e := range objects {
		if e.Kind == Conj && e.Item == "is" {
			isWords = append(isWords, e)
		}
	}
	findWordByPos := func(pos Pos) (*Object, error) {
		for _, e := range objects {
			if e.Pos == pos && e.Kind != Char {
				return e, nil
			}
		}
		return &Object{}, errors.New("Not found")
	}
	meanings := []Meaning{}

	for _, e := range isWords {
		// traverse X-range
		{
			posForw := Pos{
				e.Pos.X + 1,
				e.Pos.Y,
			}
			posBack := Pos{
				e.Pos.X - 1,
				e.Pos.Y,
			}
			objForw, errForw := findWordByPos(posForw)
			objBack, errBack := findWordByPos(posBack)
			if errForw == nil && errBack == nil && objBack.Kind == Noun && (objForw.Kind == Noun || objForw.Kind == Adj) {
				m := parseSentence(ObjectList{objBack, e, objForw})
				meanings = append(meanings, m)
			}
		}

		// traverse Y-range
		{
			posForw := Pos{
				e.Pos.X,
				e.Pos.Y + 1,
			}
			posBack := Pos{
				e.Pos.X,
				e.Pos.Y - 1,
			}
			objForw, errForw := findWordByPos(posForw)
			objBack, errBack := findWordByPos(posBack)
			if errForw == nil && errBack == nil && objBack.Kind == Noun && (objForw.Kind == Noun || objForw.Kind == Adj) {
				m := parseSentence(ObjectList{objBack, e, objForw})
				meanings = append(meanings, m)
			}
		}
	}
	return meanings
}

// Parse a correct sentence
func parseSentence(words ObjectList) Meaning {
	affected := []string{}
	modifier := []string{}

	hasEncounteredIs := false
	for i, word := range words {

		if word.Kind == Conj && word.Item == "is" {
			hasEncounteredIs = true
		}

		if i%2 == 0 { // Is not a conjunction
			if hasEncounteredIs {
				modifier = append(modifier, word.Item)
			} else {
				affected = append(affected, word.Item)
			}
		}
	}

	return Meaning{
		affected,
		modifier,
	}
}
