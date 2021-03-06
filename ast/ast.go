package ast

import (
	"bytes"
	"strings"

	"github.com/wmolicki/go-monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// interesting interface thing - cannot assign to concrete value, must be pointer type
// as methods on LetStatement have pointer receivers
var _ Statement = &LetStatement{}

type Identifier struct {
	token.Token // token.IDENT
	Value       string
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

var _ Expression = &Identifier{}

type ReturnStatement struct {
	token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

var _ Statement = &ReturnStatement{}

type ExpressionStatement struct {
	Token      token.Token // first token in expression
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) statementNode() {}

var _ Statement = &ExpressionStatement{}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) expressionNode() {}

var _ Expression = &IntegerLiteral{}

type PrefixExpression struct {
	Token    token.Token // prefix token, like "!"
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) expressionNode() {}

var _ Expression = &PrefixExpression{}

type InfixExpression struct {
	Token    token.Token // operator token, like + or -
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

func (ie *InfixExpression) expressionNode() {}

var _ Expression = &InfixExpression{}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

var _ Expression = &Boolean{}

func (b *Boolean) expressionNode() {}

type IfExpression struct {
	Token       token.Token // if token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

func (ie *IfExpression) expressionNode() {}

var _ Expression = &IfExpression{}

type BlockStatement struct {
	Token      token.Token // { token
	Statements []Statement
}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (bs *BlockStatement) statementNode() {}

var _ Statement = &BlockStatement{}

type FunctionLiteral struct {
	Token      token.Token // fn token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}

func (fl *FunctionLiteral) expressionNode() {}

var _ Expression = &FunctionLiteral{}

type CallExpression struct {
	Token     token.Token // '(' token (call exp is an "infix" expression with "(" as operator)
	Function  Expression  // identifier or function literal
	Arguments []Expression
}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, p := range ce.Arguments {
		args = append(args, p.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

func (ce *CallExpression) expressionNode() {}

var _ Expression = &CallExpression{}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

func (sl *StringLiteral) expressionNode() {}

var _ Expression = &StringLiteral{}

type ArrayLiteral struct {
	Token    token.Token // '[' token
	Elements []Expression
}

func (al *ArrayLiteral) TokenLiteral() string {
	panic("implement me")
}

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}

	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (al *ArrayLiteral) expressionNode() {}

var _ Expression = &ArrayLiteral{}

type IndexExpression struct {
	Token token.Token // '[' token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

func (ie *IndexExpression) expressionNode() {}

var _ Expression = &IndexExpression{}

type HashLiteral struct {
	Token token.Token // '{'
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }

func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for k, v := range hl.Pairs {
		pairs = append(pairs, k.String()+":"+v.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ","))
	out.WriteString("}")

	return out.String()
}

func (hl *HashLiteral) expressionNode() {}

var _ Expression = &HashLiteral{}

type ForExpression struct {
	Token       token.Token // for token
	Initializer Statement
	Condition   Expression
	Loop        Statement
	Body        *BlockStatement
}

func (fe *ForExpression) TokenLiteral() string {
	return fe.Token.Literal
}

func (fe *ForExpression) String() string {
	var out bytes.Buffer

	out.WriteString("for")
	out.WriteString("(")
	out.WriteString(fe.Initializer.String())
	out.WriteString(";")
	out.WriteString(fe.Condition.String())
	out.WriteString(";")
	out.WriteString(fe.Loop.String())
	out.WriteString(") ")
	out.WriteString(fe.Body.String())

	return out.String()
}

func (fe *ForExpression) expressionNode() {}

var _ Expression = &ForExpression{}
