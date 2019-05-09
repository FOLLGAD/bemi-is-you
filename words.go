package main

import (
	"errors"
)

type Meaning struct {
	affected []string // Item
	modifier []string // The adjectives
}

func wording(words []Object) []Meaning {
	if len(words) < 3 {
		// Too short
		return []Meaning{}
	}

	currently := false
	currentHead := 0
	hasEncounteredIs := false

	// Check sentence validity
	for i := 0; i < len(words); i++ {
		word := words[i]

		shouldBeConj := (i-currentHead)%2 == 1

		if word.Kind == Conj {
			if currently == false {
				continue
			}
			if !shouldBeConj {
				parseSentence(words[currentHead : i-1])
				currently = false
			} else if word.Item == "and" {
				continue
			} else if hasEncounteredIs {
				parseSentence(words[currentHead : i-1])
				currently = false
			}
			if word.Item == "is" {
				hasEncounteredIs = true
			}
			continue // Should be conjuction, and is.
		} else if currently {
			if shouldBeConj {
				parseSentence(words[currentHead : i-1])
				currently = false
			} else if word.Kind == Noun {
				continue
			}
		} else if word.Kind == Noun {
			currentHead = i
			currently = true
		} else if word.Kind == Adj {

		}
	}
	return []Meaning{}
}

func newWording(words []Object) []Meaning {
	isWords := []Object{}
	for _, e := range words {
		if e.Kind == Conj && e.Item == "is" {
			isWords = append(isWords, e)
		}
	}
	findWordByPos := func(pos Pos) (Object, error) {
		for _, e := range words {
			if e.Pos == pos {
				return e, nil
			}
		}
		return Object{}, errors.New("Not found")
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
				m := parseSentence([]Object{objBack, e, objForw})
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
				m := parseSentence([]Object{objBack, e, objForw})
				meanings = append(meanings, m)
			}
		}
	}
	return meanings
}

// Parse a correct sentence
func parseSentence(words []Object) Meaning {
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
