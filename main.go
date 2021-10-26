package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/wmolicki/go-monkey/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Monke REPL!\n", u.Name)
	repl.Start(os.Stdin, os.Stdout)
}
