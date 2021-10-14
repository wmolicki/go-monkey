package ast

import "github.com/wmolicki/go-monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program is a root node of AST.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token // token.LET
	Name *Identifier
	Value *Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// interesting interface thing - cannot assign to concrete value, must be pointer type
// as methods on LetStatement have pointer receivers
var _ Statement = &LetStatement{}

type Identifier struct {
	token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

var _ Expression = &Identifier{}


type ReturnStatement struct {
	token.Token // token.RETURN
	ReturnValue *Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

var _ Statement = &ReturnStatement{}