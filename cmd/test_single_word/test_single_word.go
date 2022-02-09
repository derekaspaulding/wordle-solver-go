package test_single_word

import (
	"fmt"
	"github.com/derekaspaulding/wordle-solver-go/pkg/guess"
	"github.com/derekaspaulding/wordle-solver-go/pkg/result"
	"github.com/derekaspaulding/wordle-solver-go/pkg/words"
)

func Run(word string) {
	fmt.Printf("testing word %s\n", word)
	wordReader := words.BuildDefaultWordReader("./resources/words.txt")
	guesser := guess.BuildGuesser(wordReader.ReadWords())

	guesses := []string{guesser.MakeGuess()}

	for guesses[len(guesses)-1] != word {
		guessResult := result.GuessResultFromKnownWord(word, guesses[len(guesses)-1])
		guesser.UpdateCandidatesFromResult(guessResult)

		nextGuess := guesser.MakeGuess()
		guesses = append(guesses, nextGuess)
	}

	fmt.Println(guesses)
}
