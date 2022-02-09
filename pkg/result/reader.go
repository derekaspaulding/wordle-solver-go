package result

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type ResultReader interface {
	SetOutput(output *os.File)
	ReadLetterResult(letter int32) LetterResult
	PrintInstructions()
}

type DefaultResultReader struct {
	reader io.Reader
	output *os.File
}

func (guessReader DefaultResultReader) PrintInstructions() {
	_, err := fmt.Fprintln(guessReader.output, "Enter C for correct (green), E for exists (yellow) or I for incorrect")
	if err != nil {
		panic(err)
	}
}

func (guessReader DefaultResultReader) ReadLetterResult(letter int32) LetterResult {
	scanner := bufio.NewScanner(guessReader.reader)
	fmt.Printf("%s - ", string(letter))
	scanner.Scan()

	switch strings.ToLower(scanner.Text()) {
	case "c":
		return Correct
	case "e":
		return Exists
	}

	return Incorrect
}

func (guessReader *DefaultResultReader) SetOutput(o *os.File) {
	guessReader.output = o
}

func BuildDefaultGuessReader() *DefaultResultReader {
	guessReader := new(DefaultResultReader)
	guessReader.reader = os.Stdin
	guessReader.output = os.Stdout

	return guessReader
}
