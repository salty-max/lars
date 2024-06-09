package lexer

import (
	"fmt"
	"strings"
	"testing"

	"github.com/salty-max/lars/src/token"
)

func TestNextToken(t *testing.T) {
	input := strings.Join([]string{
		`let five = 5;`,
		`let pi = 3.14;`,
		`let add = fn(x, y) {`,
		`  x + y;`,
		`};`,
		`let result = add(five, ten);`,
	}, "\n")

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		line            int
		col             int
	}{
		{token.LET, "let", 1, 1},
		{token.IDENT, "five", 1, 5},
		{token.ASSIGN, "=", 1, 10},
		{token.INT, "5", 1, 12},
		{token.SEMICOLON, ";", 1, 13},
		{token.LET, "let", 2, 1},
		{token.IDENT, "pi", 2, 5},
		{token.ASSIGN, "=", 2, 8},
		{token.FLOAT, "3.14", 2, 11},
		{token.SEMICOLON, ";", 2, 15},
		{token.LET, "let", 3, 1},
		{token.IDENT, "add", 3, 5},
		{token.ASSIGN, "=", 3, 9},
		{token.FUNCTION, "fn", 3, 12},
		{token.LPAREN, "(", 3, 14},
		{token.IDENT, "x", 3, 15},
		{token.COMMA, ",", 3, 16},
		{token.IDENT, "y", 3, 18},
		{token.RPAREN, ")", 3, 19},
		{token.LBRACE, "{", 3, 21},
		{token.IDENT, "x", 4, 3},
		{token.PLUS, "+", 4, 5},
		{token.IDENT, "y", 4, 7},
		{token.SEMICOLON, ";", 4, 8},
		{token.RBRACE, "}", 5, 1},
		{token.SEMICOLON, ";", 5, 2},
		{token.LET, "let", 6, 1},
		{token.IDENT, "result", 6, 5},
		{token.ASSIGN, "=", 6, 12},
		{token.IDENT, "add", 6, 14},
		{token.LPAREN, "(", 6, 17},
		{token.IDENT, "five", 6, 18},
		{token.COMMA, ",", 6, 22},
		{token.IDENT, "ten", 6, 26},
		{token.RPAREN, ")", 6, 29},
		{token.SEMICOLON, ";", 6, 30},
		{token.EOF, "", 6, 31},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		fmt.Printf("tests[%d] - tokentype: %q, literal: %q\n", i, tok.Type, tok.Literal)

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - tokentype wrong. expected=%q, got=%q",
				i,
				tt.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong. expected=%q, got=%q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
