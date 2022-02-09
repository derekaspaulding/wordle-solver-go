package test_all_words

import (
	"fmt"
	"github.com/derekaspaulding/wordle-solver-go/pkg/guess"
	"github.com/derekaspaulding/wordle-solver-go/pkg/result"
	"github.com/derekaspaulding/wordle-solver-go/pkg/words"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Run() {
	setupFileLogger()
	wordReader := words.BuildDefaultWordReader("./resources/words.txt")

	wordsList := wordReader.ReadWords()
	wordStats := stats{
		maxSolutionLength:     0,
		averageSolutionLength: 0,
		failedSolutions:       []string{},
	}

	for _, word := range wordsList {
		zap.S().Infof("Solving word: %s", word)
		guesser := guess.BuildGuesser(wordsList)
		guesses := []string{guesser.MakeGuess()}

		for guesses[len(guesses)-1] != word {
			guessResult := result.GuessResultFromKnownWord(word, guesses[len(guesses)-1])
			guesser.UpdateCandidatesFromResult(guessResult)

			nextGuess := guesser.MakeGuess()
			guesses = append(guesses, nextGuess)
		}

		guessLen := len(guesses)
		if guessLen > wordStats.maxSolutionLength {
			wordStats.maxSolutionLength = guessLen
		}

		if guessLen > 6 {
			wordStats.failedSolutions = append(wordStats.failedSolutions, word)
		}

		wordStats.averageSolutionLength += float32(guessLen) / float32(len(wordsList))

		zap.S().Infof("solution: %v\n\n", guesses)
	}

	fmt.Printf("Max Solution Length:     %d\n", wordStats.maxSolutionLength)
	fmt.Printf("Average Solution Length: %f\n", wordStats.averageSolutionLength)
	fmt.Printf("Number of Failures:      %d\n", len(wordStats.failedSolutions))
	fmt.Printf("Failures:                %v\n", wordStats.failedSolutions)
}

func setupFileLogger() {
	tmpFile, err := os.CreateTemp("", "wordle-solver-go-test")
	if err != nil {
		panic(err)
	}

	prodEncoderConfig := zap.NewProductionEncoderConfig()
	prodEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(prodEncoderConfig)

	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(tmpFile), zap.InfoLevel)
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)

	fmt.Printf("logs written to %s\n\n", tmpFile.Name())
}

type stats struct {
	maxSolutionLength     int
	averageSolutionLength float32
	failedSolutions       []string
}
