package interpreter

import "fmt"

type Token struct {
	tType   int
	lexeme  string
	literal any
	line    int
}

type TokenType int

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

var enumString = []string{
	"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACE", "RIGHT_BRACE", "COMMA", "DOT", "MINUS", "PLUS", "SEMICOLON", "SLASH", "STAR",
	"BANG", "BANG_EQUAL", "EQUAL", "EQUAL_EQUAL", "GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL",
	"IDENTIFIER", "STRING", "NUMBER",
	"AND", "CLASS", "ELSE", "FALSE", "FUN", "FOR", "IF", "NIL", "OR", "PRINT", "RETURN", "SUPER", "THIS", "TRUE", "VAR", "WHILE", "EOF",
}

func NewToken(index int, lexeme string, literal any, line int) Token {
	return Token{
		tType:   index,
		lexeme:  lexeme,
		literal: literal,
		line:    line}
}

func (t Token) toString() {
	fmt.Printf("%s %s %+v", enumString[t.tType], t.lexeme, t.literal)
}

func Run(source string) {
	fmt.Println(source)
}
