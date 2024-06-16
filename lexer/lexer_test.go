package lexer

import (
	"testing"

	"ljos.app/interpreter/token"
)

type expectedToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func runTestNextToken(input string, tests []expectedToken, t *testing.T) {

	lexer := New(input)

	for i, tt := range tests {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=(%q, %q), got=(%q, %q)",
				i, tt.expectedType, tt.expectedLiteral, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

	}

}

func TestNextTokenSimple(t *testing.T) {
	input := `=+(){}[],;.*`
	tests := []expectedToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.DOT, "."},
		{token.ASTERISK, "*"},
		{token.EOF, ""},
	}
	runTestNextToken(input, tests, t)
}

func TestNextTokenSimpleFn(t *testing.T) {
	input := `
  let five = 5;
  let ten = 10;

  let add = fn(x, y) {
    x + y;
  };

  let result = add(five, ten);
  `
	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	runTestNextToken(input, tests, t)
}

func TestNextTokenMoreSymbolsAndWhiteSpace(t *testing.T) {
	input := `=+(){}[],;.*
  !-/*5;
  5 < 10 > 5;
  `
	tests := []expectedToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.DOT, "."},
		{token.ASTERISK, "*"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LESS_THAN, "<"},
		{token.INT, "10"},
		{token.GREATER_THAN, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	runTestNextToken(input, tests, t)
}

func TestNextTokenWithMultiCharOperators(t *testing.T) {
	input := `
  let five = 5;
  let ten = 10;
  5 <= 10 >= 5;
  five *= 2;
  ten == five;
  five -= 2;
  five != 10;
  `

	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LESS_THAN_EQ, "<="},
		{token.INT, "10"},
		{token.GRTR_THAN_EQ, ">="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "five"},
		{token.MUL_ASSIGN, "*="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "ten"},
		{token.EQUAL, "=="},
		{token.IDENTIFIER, "five"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "five"},
		{token.MINUS_ASSIGN, "-="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.IDENTIFIER, "five"},
		{token.NOT_EQUAL, "!="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	runTestNextToken(input, tests, t)
}

func TestSimpleLambda(t *testing.T) {
	input := `let test = (a, b) => return a + b;`
	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENTIFIER, "test"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "a"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "b"},
		{token.RPAREN, ")"},
		{token.LAMBDA, "=>"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "a"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "b"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	runTestNextToken(input, tests, t)
}

func TestSimpleIfElseTrueFalse(t *testing.T) {
	input := `
  if 5 > 10 {
    return true;
  } else {
   return false;
  }
  `
	tests := []expectedToken{
		{token.IF, "if"},
		{token.INT, "5"},
		{token.GREATER_THAN, ">"},
		{token.INT, "10"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}
	runTestNextToken(input, tests, t)
}
