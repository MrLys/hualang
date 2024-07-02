package parser

import (
	"fmt"
	"strings"

	"ljos.app/interpreter/ast"
	"ljos.app/interpreter/lexer"
	"ljos.app/interpreter/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type ParserError struct {
	Error string
	Lines string
}
type Parser struct {
	l      *lexer.Lexer
	errors []ParserError

	curToken  token.Token
	peekToken token.Token

	currentLines   []string
	currentLineIdx int

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func New(l *lexer.Lexer) *Parser {
	p := Parser{
		l:      l,
		errors: []ParserError{},
	}
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	// fill curToken and peekToken
	p.nextToken()
	p.nextToken()
	return &p
}
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
	p.currentLines = p.l.GetCurrentLines()
	p.currentLineIdx = p.l.CurrentLineIdx
}

func (p *Parser) Errors() []ParserError {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	lines := strings.Join(p.currentLines, "\n")
	msg := fmt.Sprintf("expected next token to be %s, got %s instead (on line %d)",
		t, p.peekToken.Type, p.currentLineIdx)
	p.errors = append(p.errors, ParserError{
		Error: msg,
		Lines: lines,
	})
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		statement := p.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {

	var statement ast.Statement

	if p.curToken.Type == token.LET {
		// fmt.Printf("parseStatement curToken type is %s\n", p.curToken.Type)
		statement = p.parseLetStatement()
	} else if p.curToken.Type == token.RETURN {
		// fmt.Printf("parseStatement curToken type is %s\n", p.curToken.Type)
		statement = p.parseReturnStatement()
	} else if p.curToken.Type == token.IF {
		// fmt.Printf("parseStatement curToken type is %s\n", p.curToken.Type)
		statement = p.parseIfStatement()
	} else {
		// fmt.Printf("parseStatement curToken type is %s\n", p.curToken.Type)
		return p.parseExpressionStatement()
	}
	return statement

}

func (p *Parser) expectedToken(expectedToken token.TokenType) bool {
	if p.peekTokenIs(expectedToken) {
		p.nextToken()
		return true
	}
	p.peekError(expectedToken)
	return false
}

func (p *Parser) curTokenIs(expectedToken token.TokenType) bool {
	return p.curToken.Type == expectedToken
}

func (p *Parser) peekTokenIs(expectedToken token.TokenType) bool {
	return p.peekToken.Type == expectedToken
}
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectedToken(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectedToken(token.ASSIGN) {
		return nil
	}
	// skip until semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	var expr ast.IfStatement
	return &expr
}

func (p *Parser) parseError(msg string, args ...any) {
	fmt.Printf(msg, args...)
}
