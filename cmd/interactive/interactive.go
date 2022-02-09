package interactive

import (
	"fmt"
	"github.com/derekaspaulding/wordle-solver-go/pkg/guess"
	"github.com/derekaspaulding/wordle-solver-go/pkg/result"
	"github.com/derekaspaulding/wordle-solver-go/pkg/words"
	"os"
)

func Run() {
	wordReader := words.BuildDefaultWordReader("./resources/words.txt")
	guessReader := result.BuildDefaultGuessReader()
	guesser := guess.BuildGuesser(wordReader.ReadWords())

	guesses := []string{guesser.MakeGuess()}
	for len(guesses) < 6 {
		guessResult := result.GuessResultFromUser(guesses[len(guesses)-1], guessReader, os.Stdout)

		if isCorrectGuess(guessResult) {
			break
		}

		guesser.UpdateCandidatesFromResult(guessResult)
		guesses = append(guesses, guesser.MakeGuess())
	}

	fmt.Println(guesses)
}

func isCorrectGuess(guessResult result.GuessResult) bool {
	for _, letterResult := range guessResult.Result {
		if letterResult != result.Correct {
			return false
		}
	}
	return true
}
