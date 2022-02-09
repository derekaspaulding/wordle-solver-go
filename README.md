# Wordle Sover

Solves the game [Wordle](https://www.powerlanguage.co.uk/wordle/).

**Note:** The implementation is tied to some internal implementation of the game found by digging through it's JS files.
Specifically, it is tied to the list of words. If the game is updated in the future to include more words than the ones
in `resources/words.txt`, this program will fail on those words.

## Usage

Assuming you have [Go](https://go.dev) installed. Tested on Go 1.17

1. Clone the repo
2. Run `go build .` in the repo root directory
    - Currently relies on the words list being at `./resources/words.txt`. You can use `go install` and run from
      anywhere that that is true
3. Run `./wordle-solver-go` to print options

### Options

- `-i` - Interactive Mode. Will walk you through making guesses in the game and asking for the results the game provides
  to solve the word.
- `-w [word]` - gives the suggested solution for a specific word. Good for testing a single word's solution
- `-t` - Runs the solver on every word it knows about and gathers some statistics on the current solution. 


