package repl

import (
	"fmt"
	"io"

	"github.com/chzyer/readline"

	"github.com/wmolicki/go-monkey/lexer"
	"github.com/wmolicki/go-monkey/token"

)

const PROMPT = ">> "


func Start(in io.Reader, out io.Writer) {
	scanner, err := readline.New(PROMPT)
	if err != nil {
		panic(err)
	}

	for {
		line, err := scanner.Readline()
		if err != nil || line == "\\q" {
			return
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}