package main

import (
	"flag"
	"github.com/derekaspaulding/wordle-solver-go/cmd/interactive"
	"github.com/derekaspaulding/wordle-solver-go/cmd/test_all_words"
	"github.com/derekaspaulding/wordle-solver-go/cmd/test_single_word"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	configureLogger()
	flagValues := getFlags()

	if flagValues.isInteractive {
		interactive.Run()
	} else if flagValues.testWord != "" {
		test_single_word.Run(flagValues.testWord)
	} else if flagValues.testAll {
		test_all_words.Run()
	} else {
		flag.Usage()
	}
}

func configureLogger() {
	prodEncoderConfig := zap.NewProductionEncoderConfig()
	prodEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(prodEncoderConfig)

	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.InfoLevel)
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
}

type flags struct {
	isInteractive bool
	testWord      string
	testAll       bool
}

func getFlags() flags {
	isInteractive := flag.Bool("i", false, "run in interactive mode")
	testWord := flag.String("w", "", "get solution for given word")
	testAll := flag.Bool("t", false, "gather statistics on all solutions for all words")
	flag.Parse()

	return flags{
		isInteractive: *isInteractive,
		testWord:      *testWord,
		testAll:       *testAll,
	}
}
