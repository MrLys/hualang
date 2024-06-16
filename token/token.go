package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	ASSIGN       = "="
	MINUS        = "-"
	PLUS         = "+"
	ASTERISK     = "*"
	SEMICOLON    = ";"
	COMMA        = ","
	DOT          = "."
	SLASH        = "/"
	BANG         = "!"
	LESS_THAN    = "<"
	GREATER_THAN = ">"
	GRTR_THAN_EQ = ">="
	LESS_THAN_EQ = "<="
	LAMBDA       = "=>"
	EQUAL        = "=="
	NOT_EQUAL    = "!="
	MUL_ASSIGN   = "*="
	MINUS_ASSIGN = "-="
	PLUS_ASSIGN  = "+="

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"
	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)
