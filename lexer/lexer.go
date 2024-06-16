package lexer

import (
	token "ljos.app/interpreter/token"
)

type Lexer struct {
	input        string
	position     int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch           byte
}
type TokenLambda func() token.Token

func newToken(input string, tokenType token.TokenType) TokenLambda {
	return func() token.Token {
		return token.Token{Type: tokenType, Literal: input}
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isNumber(ch byte) bool {
	return ch <= '9' && ch >= '0'
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func getToken(ch byte) (TokenLambda, bool) {
	val, ok := tokenMap[string(ch)]
	return val, ok
}

func getTokenFromString(identifier string) (TokenLambda, bool) {
	val, ok := tokenMap[identifier]
	return val, ok
}

var multiCharTokenMap = map[byte]struct{}{
	'=': {},
	'<': {},
	'>': {},
	'!': {},
	'-': {},
	'*': {},
	'+': {},
	'/': {},
}
var tokenMap = map[string]TokenLambda{
	"=":      newToken("=", token.ASSIGN),
	";":      newToken(";", token.SEMICOLON),
	"(":      newToken("(", token.LPAREN),
	")":      newToken(")", token.RPAREN),
	"{":      newToken("{", token.LBRACE),
	"}":      newToken("}", token.RBRACE),
	"[":      newToken("[", token.LBRACK),
	"]":      newToken("]", token.RBRACK),
	",":      newToken(",", token.COMMA),
	".":      newToken(".", token.DOT),
	"+":      newToken("+", token.PLUS),
	"-":      newToken("-", token.MINUS),
	"!":      newToken("!", token.BANG),
	"<":      newToken("<", token.LESS_THAN),
	">":      newToken(">", token.GREATER_THAN),
	"/":      newToken("/", token.SLASH),
	"*":      newToken("*", token.ASTERISK),
	"+=":     newToken("+=", token.PLUS),
	"-=":     newToken("-=", token.MINUS_ASSIGN),
	"!=":     newToken("!=", token.NOT_EQUAL),
	"<=":     newToken("<=", token.LESS_THAN_EQ),
	">=":     newToken(">=", token.GRTR_THAN_EQ),
	"/=":     newToken("/=", token.SLASH),
	"*=":     newToken("*=", token.MUL_ASSIGN),
	"==":     newToken("==", token.EQUAL),
	"=>":     newToken("=>", token.LAMBDA),
	"let":    newToken("let", token.LET),
	"fn":     newToken("fn", token.FUNCTION),
	"return": newToken("return", token.RETURN),
	"if":     newToken("if", token.IF),
	"else":   newToken("else", token.ELSE),
	"true":   newToken("true", token.TRUE),
	"false":  newToken("false", token.FALSE),
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

	var tok token.Token
	l.skipWhiteSpace()

	if isNumber(l.ch) {
		tok.Literal = l.readNumber()
		tok.Type = token.INT
		return tok
	}

	if val, ok := getToken(l.ch); ok {
		t := val()
		defer l.readChar()
		// if its a normal token then we need to check if it is a multi char token like ==
		if _, ok := multiCharTokenMap[l.ch]; ok {
			multiCharToken := string(l.ch)
			nextChar := l.peekChar()
			if nextChar == 0 {
				return t
			}
			multiCharToken += string(nextChar)
			if val, ok = getTokenFromString(multiCharToken); ok {
				// read an additional char for multiCharToken
				// this only works if the multi char token i of length 2
				l.readChar()
				return val()
			}
		}
		return t
	}

	if isLetter(l.ch) {
		tok.Literal = l.readIdentifier()
		if v, ok := getTokenFromString(tok.Literal); !ok {
			tok.Type = token.IDENTIFIER
		} else {
			tok = v()
		}

		return tok
	}

	if l.ch == 0 {
		tok.Type = token.EOF
		tok.Literal = ""
		return tok
	}

	tok.Literal = string(l.ch)
	tok.Type = token.ILLEGAL

	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for isWhiteSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isNumber(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}
