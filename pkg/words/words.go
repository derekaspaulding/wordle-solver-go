package words

import (
	"bufio"
	"os"
)

type WordReader interface {
	ReadWords() []string
}

type DefaultWordReader struct {
	filePath string
}

func (d DefaultWordReader) ReadWords() []string {
	file, err := os.Open(d.filePath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func BuildDefaultWordReader(filePath string) *DefaultWordReader {
	wordReader := new(DefaultWordReader)
	wordReader.filePath = filePath

	return wordReader
}
