package parser

import (
	"testing"

	"ljos.app/interpreter/ast"
	"ljos.app/interpreter/lexer"
)

func TestReturnStatements(t *testing.T) {
	input := `
  return 5;
  return 10;
  return 730246;
  `
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain exactly 3 statements, got %d", len(program.Statements))
		return
	}
	for i, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement(%d). got %T", i, stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}

	}

}
func TestLetStatements(t *testing.T) {
	input := `
  let x = 5;
  let y = 10;
  let foobar = 75319246;
  `
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatements(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatements(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral() not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt not *ast.LetStatement. got=%T", stmt)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("stmt.Name.Value not %s. got=%T", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("stmt.Name not %s. got=%T", name, letStmt.Name)
		return false
	}
	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error %q\n", msg.Error)
		t.Errorf("\n%s", msg.Lines)
	}
	t.FailNow()
}

func TestIdentifiers(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("program should contain 1 statement, got %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.Expression, got %T", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("stmt is not *ast.Identifier, got %T", stmt)
	}
	if ident.Value != "foobar" {
		t.Fatalf("stmt.TokenLiteral() should be 'foobar', got %q", stmt.TokenLiteral())
	}

}
