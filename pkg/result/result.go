package result

import (
	"fmt"
	"os"
	"strings"
)

type LetterResult string

const (
	Correct   LetterResult = "Correct"
	Incorrect              = "Incorrect"
	Exists                 = "Exists"
)

type GuessResult struct {
	Word   string
	Result []LetterResult
}

// TODO: create a GuessAnalyzer object that can hold the output filend the result reader
func GuessResultFromUser(guessWord string, reader ResultReader, output *os.File) GuessResult {
	guess := new(GuessResult)
	guess.Word = guessWord
	reader.SetOutput(output)

	_, err := fmt.Fprintf(output, "The word you guessed was: \"%s\"\n", guess.Word)
	if err != nil {
		panic(err)
	}

	reader.PrintInstructions()
	for _, letter := range guess.Word {
		letterResult := reader.ReadLetterResult(letter)
		guess.Result = append(guess.Result, letterResult)
	}

	return *guess
}

func GuessResultFromKnownWord(word string, guessWord string) GuessResult {
	var result []LetterResult
	for i, letter := range guessWord {
		if letter == []rune(word)[i] {
			result = append(result, Correct)
		} else if strings.Contains(word, string(letter)) {
			result = append(result, Exists)
		} else {
			result = append(result, Incorrect)
		}
	}

	return GuessResult{
		Word:   guessWord,
		Result: result,
	}
}
