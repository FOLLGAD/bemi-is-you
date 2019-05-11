package main

import (
	"errors"
)

type Meanings map[string]map[string]bool

func (union *Meanings) AddMeanings(meanings Meanings) {
	for affected := range meanings {
		for modifier := range meanings[affected] {
			if (*union)[affected] == nil {
				(*union)[affected] = map[string]bool{}
			}
			(*union)[affected][modifier] = true
		}
	}
}

func findSentences(objects ObjectList) Meanings {
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
	meanings := Meanings{}

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
				meanings.AddMeanings(m)
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
				meanings.AddMeanings(m)
			}
		}
	}
	return meanings
}

// Parse a correct sentence
func parseSentence(words ObjectList) Meanings {
	meaning := Meanings{}

	affected := []string{}

	hasEncounteredIs := false
	for i, word := range words {

		if word.Kind == Conj && word.Item == "is" {
			hasEncounteredIs = true
		}

		if i%2 == 0 { // Is not a conjunction
			if hasEncounteredIs {
				for _, e := range affected {
					if meaning[e] == nil {
						meaning[e] = map[string]bool{}
					}
					meaning[e][word.Item] = true
				}
			} else {
				affected = append(affected, word.Item)
			}
		}
	}

	return meaning
}
