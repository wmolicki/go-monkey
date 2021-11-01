package repl

import (
	"io"

	"github.com/chzyer/readline"

	"github.com/wmolicki/go-monkey/evaluator"
	"github.com/wmolicki/go-monkey/lexer"
	"github.com/wmolicki/go-monkey/object"
	"github.com/wmolicki/go-monkey/parser"
)

const PROMPT = ">> "

func Start(in io.ReadCloser, out io.Writer) {
	conf := &readline.Config{Prompt: PROMPT}
	scanner, err := readline.NewEx(conf)
	if err != nil {
		panic(err)
	}
	env := object.NewEnvironment()

	for {
		line, err := scanner.Readline()
		if err != nil || line == "\\q" {
			io.WriteString(out, "bye")
			return
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Error interpreting program\n")
	io.WriteString(out, "  parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
