package parser

import (
	"fmt"

	"ljos.app/interpreter/ast"
	"ljos.app/interpreter/lexer"
	"ljos.app/interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := Parser{
		l: l,
	}
	// fill curToken and peekToken
	p.nextToken()
	p.nextToken()
	return &p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		statement := p.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return &program
}

func (p *Parser) parseStatement() ast.Statement {

	var statement ast.Statement

	if p.curToken.Type == token.LET {
		statement = p.parseLetStatement()
	} else if p.curToken.Type == token.RETURN {
		statement = p.parseReturnStatement()
	} else if p.curToken.Type == token.IF {
		statement = p.parseIfStatement()
	}
	return statement

}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := ast.LetStatement{Token: p.curToken}
	if p.peekToken.Type != token.IDENTIFIER {
		p.parseError("Identifier expected after let")
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.peekToken, Value: p.peekToken.Literal}
	p.nextToken()
	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}
	//if p.curToken.Type != token.ASSIGN {
	//	p.parseError("Equal sign expected after %s in let expression", stmt.Name.TokenLiteral())
	//	return nil
	//}

	return &stmt
}

func (p *Parser) parseExpression() *ast.Expression {
	var expr ast.Expression
	return &expr
}
func (p *Parser) parseReturnStatement() *ast.LetStatement {
	var expr ast.LetStatement
	return &expr
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	var expr ast.IfStatement
	return &expr
}

func (p *Parser) parseError(msg string, args ...any) {
	fmt.Printf(msg, args...)
}
