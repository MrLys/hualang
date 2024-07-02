package ast

import (
	"bytes"

	"ljos.app/interpreter/token"
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
	Token token.Token // the token.Let token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.Value)
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token token.Token // the token.Identifier Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type IfStatement struct {
	Token     token.Token // the token.IfStatement Token
	Condition Expression
	Value     Expression
	ElseValue Expression
	ElseIf    *IfStatement
}

func (i *IfStatement) statementNode() {}
func (i *IfStatement) TokenLiteral() string {
	return i.Token.Literal
}

func (is *IfStatement) String() string {
	var out bytes.Buffer
	out.WriteString(is.TokenLiteral())
	out.WriteString(" ")
	if is.Condition != nil {
		out.WriteString(is.Condition.String())
	}
	out.WriteString(" {")
	if is.Value != nil {
		out.WriteString(is.Value.String())
	}
	out.WriteString("}")
	if is.ElseValue != nil {
		out.WriteString(" else {")
		out.WriteString(is.ElseValue.String())
		out.WriteString("}")
	}

	return out.String()
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (rs *ExpressionStatement) statementNode() {}
func (rs *ExpressionStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
