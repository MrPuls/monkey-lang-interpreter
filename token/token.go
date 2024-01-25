package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT" // add, foobar, x, y, ...
	INT       = "INT"   // 1234567890
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
