package result_test

import (
	"github.com/derekaspaulding/wordle-solver-go/pkg/result"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGuessResultFromUser(t *testing.T) {
	testWord := "words"
	guessReaderMock := new(mockGuessReader)
	tempFile, err := os.CreateTemp("", "testGuessResultFromUser-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := tempFile.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	expectedResults := []result.LetterResult{
		result.Incorrect,
		result.Exists,
		result.Correct,
		result.Correct,
		result.Incorrect,
	}
	guessReaderMock.Result = expectedResults

	guessResult := result.GuessResultFromUser(testWord, guessReaderMock, tempFile)

	assert.Equal(t, result.GuessResult{Word: testWord, Result: expectedResults}, guessResult)
}

func TestGuessResultFromKnownWord(t *testing.T) {
	// TODO: Create test table
	testWord := "words"
	testGuess := "wrist"
	guessResult := result.GuessResultFromKnownWord(testWord, testGuess)

	assert.Equal(
		t,
		result.GuessResult{
			Word: testGuess,
			Result: []result.LetterResult{
				result.Correct,
				result.Exists,
				result.Incorrect,
				result.Exists,
				result.Incorrect,
			},
		},
		guessResult,
	)
}

type mockGuessReader struct {
	Result []result.LetterResult
}

func (m *mockGuessReader) SetOutput(_ *os.File) {}

func (m mockGuessReader) PrintInstructions() {}

func (m *mockGuessReader) ReadLetterResult(_ int32) result.LetterResult {
	letterResult := m.Result[0]
	m.Result = (m.Result)[1:]
	return letterResult
}
