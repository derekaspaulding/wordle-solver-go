package guess

import (
	"github.com/derekaspaulding/wordle-solver-go/pkg/result"
	"go.uber.org/zap"
	"strings"
)

type Guesser struct {
	candidates        []string
	letterFrequencies map[rune]int
}

func (guesser *Guesser) MakeGuess() string {
	bestScore := 0
	var bestCandidates []string

	for _, word := range guesser.candidates {
		score := guesser.calculateWordScore(word)
		if score > bestScore {
			bestCandidates = []string{word}
			bestScore = score
		} else if score == bestScore {
			bestCandidates = append(bestCandidates, word)
		}
	}

	return bestCandidates[0]
}

func (guesser *Guesser) UpdateCandidatesFromResult(guessResult result.GuessResult) {
	var newCandidates []string
	for _, word := range guesser.candidates {
		if isWordStillValid(word, guessResult) {
			newCandidates = append(newCandidates, word)
		}
	}

	// TODO: implement a logger
	zap.S().Infof("Eliminate %d of %d words", len(guesser.candidates)-len(newCandidates), len(guesser.candidates))

	guesser.candidates = newCandidates
	guesser.updateLetterFrequencies()
}

func (guesser *Guesser) updateLetterFrequencies() {
	guesser.resetLetterFrequencies()
	for _, word := range guesser.candidates {
		runeSet := wordSet{word}.getSet()
		for letter := range runeSet {
			guesser.letterFrequencies[letter]++
		}
	}
}

func (guesser *Guesser) resetLetterFrequencies() {
	start := []rune("a")[0]
	end := []rune("z")[0]
	guesser.letterFrequencies = make(map[rune]int)

	for letter := start; letter <= end; letter++ {
		guesser.letterFrequencies[letter] = 0
	}
}

func (guesser Guesser) calculateWordScore(word string) int {
	score := 0
	runeSet := wordSet{word}.getSet()
	for letter := range runeSet {
		score += guesser.letterFrequencies[letter]
	}

	return score
}

func isWordStillValid(word string, guessResult result.GuessResult) bool {
	guessRunes := []rune(guessResult.Word)
	for i, letter := range word {
		lettersMatch := guessRunes[i] == letter

		switch guessResult.Result[i] {
		case result.Correct:
			if !lettersMatch {
				return false
			}
			break
		case result.Incorrect:
			if strings.Contains(word, string(guessRunes[i])) {
				return false
			}
			break
		case result.Exists:
			if lettersMatch || !strings.Contains(word, string(guessRunes[i])) {
				return false
			}
			break
		}
	}

	return true
}

func BuildGuesser(candidates []string) *Guesser {
	guesser := new(Guesser)
	guesser.candidates = candidates
	guesser.updateLetterFrequencies()

	return guesser
}
