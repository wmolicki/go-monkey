package ast

import (
	"testing"

	"github.com/wmolicki/go-monkey/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "myVar"}, Value: "myVar"},
				Value: &Identifier{Token: token.Token{Type: token.IDENT, Literal: "anotherVal"}, Value: "anotherVal"},
			},
		},
	}
	if program.String() != "let myVar = anotherVal;" {
		t.Errorf("program.String() did not return correct value. got=%s", program.String())
	}
}
