package token

// TokenType represents a type of token.
type TokenType string

// Token represents a token in the source code.
type Token struct {
	Type    TokenType
	Literal string
	Col     int
	Line    int
}

// LookupIdent checks if the given identifier is a keyword.
// If it is, it returns the keyword's TokenType.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"
	CHAR   = "CHAR"

	// Operators
	ASSIGN     = "="
	PLUS       = "+"
	MINUS      = "-"
	STAR       = "*"
	SLASH      = "/"
	PERCENT    = "%"
	INC        = "++"
	DEC        = "--"
	PLUS_EQ    = "+="
	MINUS_EQ   = "-="
	STAR_EQ    = "*="
	SLASH_EQ   = "/="
	PERCENT_EQ = "%="

	// Logical operators
	BANG    = "!"
	AND     = "&&"
	OR      = "||"
	EQ      = "=="
	NOT_EQ  = "!="
	LT      = "<"
	GT      = ">"
	LTE     = "<="
	GTE     = ">="
	BIT_AND = "&"
	BIT_OR  = "|"
	BIT_XOR = "^"
	BIT_NOT = "~"
	LSHIFT  = "<<"
	RSHIFT  = ">>"

	// Delimiters
	COMMA     = ","
	DOT       = "."
	DOTDOT    = ".."
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONST    = "CONST"
	CLASS    = "CLASS"
	STRUCT   = "STRUCT"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELIF     = "ELIF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FOR      = "FOR"
	IN       = "IN"
	WHILE    = "WHILE"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	NULL     = "NULL"
	TYPEOF   = "TYPEOF"
)

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"let":      LET,
	"const":    CONST,
	"class":    CLASS,
	"struct":   STRUCT,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"elif":     ELIF,
	"else":     ELSE,
	"return":   RETURN,
	"for":      FOR,
	"in":       IN,
	"while":    WHILE,
	"break":    BREAK,
	"continue": CONTINUE,
	"null":     NULL,
	"typeof":   TYPEOF,
}
