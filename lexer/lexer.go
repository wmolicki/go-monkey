package lexer

import "github.com/wmolicki/go-monkey/token"

type Lexer struct {
	input string
	position int // points to current ch
	readPosition int // current reading position (after ch)
	ch byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			t = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			t = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '-':
		t = newToken(token.MINUS, l.ch)
	case '*':
		t = newToken(token.ASTERISK, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			t = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			t = newToken(token.BANG, l.ch)
		}
	case '/':
		t = newToken(token.SLASH, l.ch)
	case '>':
		t = newToken(token.GT, l.ch)
	case '<':
		t = newToken(token.LT, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdentifier(t.Literal)
			return t
		} else if isDigit(l.ch) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch) }
}


// readIdentifier reads an identifier and advances Lexer.position
// until it encounters non-letter char.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}