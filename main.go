package main

import (
	"fmt"
	"os"

	"github.com/wmolicki/go-monkey/repl"
)

func main() {
	fmt.Println("Monke REPL!")
	repl.Start(os.Stdin, os.Stdout)
}
