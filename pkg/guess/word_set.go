package guess

type wordSet struct {
	word string
}

func (ws wordSet) getSet() map[rune]bool {
	set := make(map[rune]bool)

	for _, letter := range ws.word {
		set[letter] = true
	}

	return set
}
