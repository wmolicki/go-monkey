package parser

import (
	"testing"

	"github.com/wmolicki/go-monkey/ast"
	"github.com/wmolicki/go-monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
 	let monke = 83;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements should contain 3 statements, got %d",
			len(program.Statements))
	}

	tests := []struct{expectedIdent string}{
		{"x"},
		{"y"},
		{"monke"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if testLetStatement(t, statement, tt.expectedIdent) == false {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser had %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)

	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() should return let, got=%q", s.TokenLiteral())
		return false
	}

	// To test whether an interface value holds a specific type,
	// a type assertion can return two values:
	// the underlying value and a boolean value that reports whether
	// the assertion succeeded.
	//
	//t, ok := i.(T)
	letStatement, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not %s, got=%s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("s.Name not %s, got=%s", name, letStatement.Name)
		return false
	}
	return true
}


func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 666;
	return 5 + 6;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 4 {
		t.Fatalf("should contain 4 statements, got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt did not return *ast.ReturnStatement, got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral() did not return 'return', got=%T", stmt)
		}
	}
}