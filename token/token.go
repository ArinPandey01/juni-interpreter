package token

type TokenType int

const (
	// Single-character tokens
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	LEFT_BRACKET
	RIGHT_BRACKET
	SEMICOLON
	COMMA
	COLON
	DOT

	// Mathematical operators
	PLUS
	PLUS_EQUAL
	MINUS
	MINUS_EQUAL
	STAR
	STAR_EQUAL
	SLASH
	SLASH_EQUAL
	MODULO
	MODULO_EQUAL

	// Comparison / logical operators
	EQUAL
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals
	IDENTIFIER
	INTEGER
	FLOAT
	STRING

	// Keywords
	IF
	ELSE
	FOR
	FUNCTION
	RETURN
	VAR
	TRUE
	FALSE
	NULL

	EOF
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}
