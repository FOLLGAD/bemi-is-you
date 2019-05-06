package main

func wording(words []Object) {
	if len(words) < 3 {
		// Too short
		return
	}

	// Check sentence validity
	hasEncounteredIs := false
	for i, w := range words {
		if w.Kind == Conj && w.Item == "is" {
			if hasEncounteredIs {
				return
			}
			hasEncounteredIs = true
		}
		shouldBeConjug := i%2 == 1
		isConjug := w.Kind == Conj
		if shouldBeConjug != isConjug {
			return // Not a correct sentence
		}
	}

	affectedWords := []Object{}
	for i, w := range words {

	}
}
