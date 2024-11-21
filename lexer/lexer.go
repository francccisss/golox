package lexer

import "fmt"

// takes in a stream of characters, finds a pattern of a
// lexeme if it matches any of the token type defined in the
// TokenType enum, if so create a <token> out of that lexeme,
// proceed to the next character in the stream

const (
	// Single-character tokens.
	LEFT_PAREN = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF
)

func Run(source string) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens()
	for _, token := range tokens {
		token.toString()

	}
}

func (t Token) toString() {
	fmt.Printf("{ Token Type: %d, Lexeme: %s, Literal: %+v }\n", t.tType, t.lexeme, t.literal)
}
