package lexer

import (
	"github.com/salty-max/lars/src/token"
)

// Lexer represents a lexer.
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	col          int  // current column in line
	line         int  // current line
}

// New creates a new lexer.
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
	l.readChar()
	return l
}

// NextToken returns the next token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.line, l.col)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.line, l.col)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.line, l.col)
	case ':':
		tok = newToken(token.COLON, l.ch, l.line, l.col)
	case '.':
		if l.peekChar() == '.' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.DOTDOT, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.DOT, l.ch, l.line, l.col)
		}
	case '+':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.PLUS_EQ, Literal: literal, Col: l.col, Line: l.line}
		} else if l.peekChar() == '+' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.INC, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.PLUS, l.ch, l.line, l.col)
		}
	case '-':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.MINUS_EQ, Literal: literal, Col: l.col, Line: l.line}
		} else if l.peekChar() == '-' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.DEC, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.MINUS, l.ch, l.line, l.col)
		}
	case '*':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.STAR_EQ, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.STAR, l.ch, l.line, l.col)
		}
	case '/':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.SLASH_EQ, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.SLASH, l.ch, l.line, l.col)
		}
	case '%':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.PERCENT_EQ, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.PERCENT, l.ch, l.line, l.col)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.BANG, l.ch, l.line, l.col)
		}
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.line, l.col)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.line, l.col)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.line, l.col)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.line, l.col)
	case '[':
		tok = newToken(token.LBRACKET, l.ch, l.line, l.col)
	case ']':
		tok = newToken(token.RBRACKET, l.ch, l.line, l.col)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LTE, Literal: literal, Col: l.col, Line: l.line}
		} else if l.peekChar() == '<' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LSHIFT, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.LT, l.ch, l.line, l.col)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GTE, Literal: literal, Col: l.col, Line: l.line}

		} else if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.RSHIFT, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.GT, l.ch, l.line, l.col)
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.AND, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.BIT_AND, l.ch, l.line, l.col)
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.OR, Literal: literal, Col: l.col, Line: l.line}
		} else {
			tok = newToken(token.BIT_OR, l.ch, l.line, l.col)
		}
	case '^':
		tok = newToken(token.BIT_XOR, l.ch, l.line, l.col)
	case '~':
		tok = newToken(token.BIT_NOT, l.ch, l.line, l.col)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {

			tok.Literal = l.readNumber()
			if isFloat(tok.Literal) {
				tok.Type = token.FLOAT
			} else {
				tok.Type = token.INT
			}

			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.line, l.col)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte, line int, col int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Col: col, Line: line}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line += 1
			l.col = 0
		}
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
	l.col += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readNumber() string {
	position := l.position
	hasDot := false

	for isDigit(l.ch) || (l.ch == '.' && !hasDot) {
		if l.ch == '.' {
			hasDot = true
		}
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isFloat(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return true
		}
	}
	return false
}
