package main

import (
	"fmt"
	"flag"
	"io"
	"os"

	"github.com/wmolicki/go-monkey/evaluator"
	"github.com/wmolicki/go-monkey/lexer"
	"github.com/wmolicki/go-monkey/object"
	"github.com/wmolicki/go-monkey/parser"
	"github.com/wmolicki/go-monkey/repl"
)

func main() {
	flag.Parse()
	files := flag.Args()

	switch len(files) {
	case 0:
		fmt.Println("Monke REPL!")
		repl.Start(os.Stdin, os.Stdout)
	case 1:
		filename := files[0]
		script, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("error reading script: %v\n", err)
			os.Exit(1)
		}

		l := lexer.New(string(script))
		p := parser.New(l)

		env := object.NewEnvironment()

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(os.Stdout, p.Errors())
			os.Exit(2)
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	default:
		fmt.Printf("unsupported amount of arguments: %d\n", len(files))
		os.Exit(1)
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Error interpreting program\n")
	io.WriteString(out, "  parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
