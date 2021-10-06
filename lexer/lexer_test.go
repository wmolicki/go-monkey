package lexer

import "testing"

import "github.com/wmolicki/go-monkey/token/token"

func TestNextToken(t *testing.T) {
	input := `=+)(}{,;`

	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		token.
	}
}