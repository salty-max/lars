package lexer

import (
	"strings"
	"testing"

	"github.com/salty-max/lars/src/token"
)

func TestNextToken(t *testing.T) {
	input := strings.Join([]string{
		`let five = 5;`,
		`let pi = 3.14;`,
		``,
		`let add = fn(x, y) {`,
		`  x + y;`,
		`};`,
		``,
		`let result = add(five, ten);`,
		`!-/*5;`,
		`5 < 10 > 5;`,
		``,
		`if (5 < 10) {`,
		`  return true;`,
		`} else {`,
		`  return false;`,
		`}`,
		``,
		`10 == 10;`,
		`10 != 9;`,
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
		{token.BANG, "!", 7, 1},
		{token.MINUS, "-", 7, 2},
		{token.SLASH, "/", 7, 3},
		{token.STAR, "*", 7, 4},
		{token.INT, "5", 7, 5},
		{token.SEMICOLON, ";", 7, 6},
		{token.INT, "5", 8, 1},
		{token.LT, "<", 8, 3},
		{token.INT, "10", 8, 5},
		{token.GT, ">", 8, 8},
		{token.INT, "5", 8, 10},
		{token.SEMICOLON, ";", 8, 11},
		{token.IF, "if", 10, 1},
		{token.LPAREN, "(", 10, 4},
		{token.INT, "5", 10, 5},
		{token.LT, "<", 10, 7},
		{token.INT, "10", 10, 9},
		{token.RPAREN, ")", 10, 11},
		{token.LBRACE, "{", 10, 13},
		{token.RETURN, "return", 11, 3},
		{token.TRUE, "true", 11, 10},
		{token.SEMICOLON, ";", 11, 14},
		{token.RBRACE, "}", 12, 1},
		{token.ELSE, "else", 12, 3},
		{token.LBRACE, "{", 12, 8},
		{token.RETURN, "return", 13, 3},
		{token.FALSE, "false", 13, 10},
		{token.SEMICOLON, ";", 13, 15},
		{token.RBRACE, "}", 14, 1},
		{token.INT, "10", 16, 1},
		{token.EQ, "==", 16, 4},
		{token.INT, "10", 16, 7},
		{token.SEMICOLON, ";", 16, 9},
		{token.INT, "10", 17, 1},
		{token.NOT_EQ, "!=", 17, 4},
		{token.INT, "9", 17, 7},
		{token.SEMICOLON, ";", 17, 8},
		{token.EOF, "", 18, 1},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] [%d:%d] - tokentype wrong. expected=%q, got=%q",
				i,
				tok.Line,
				tok.Col,
				tt.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] [%d: %d] - literal wrong. expected=%q, got=%q",
				i,
				tok.Line,
				tok.Col,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
