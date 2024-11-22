package lexer

type Token struct {
	tType   int
	lexeme  string
	literal any
	line    int
}

type TokenType int

var Keywords = map[string]int{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"fun":    FUN,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
	"eof":    EOF,
}

var enumString = []string{
	"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACE", "RIGHT_BRACE", "COMMA", "DOT", "MINUS", "PLUS", "SEMICOLON", "SLASH", "STAR",
	"BANG", "BANG_EQUAL", "EQUAL", "EQUAL_EQUAL", "GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL",
	"IDENTIFIER", "STRING", "NUMBER",
	"AND", "CLASS", "ELSE", "FALSE", "FUN", "FOR", "IF", "NIL", "OR", "PRINT", "RETURN", "SUPER", "THIS", "TRUE", "VAR", "WHILE", "EOF",
}

func NewToken(tokenType int, lexeme string, literal any, line int) Token {
	return Token{
		tType:   tokenType,
		lexeme:  lexeme,
		literal: literal,
		line:    line}
}

func (t Token) toString() {
	fmt.Printf("{ Token Type: %s, Lexeme: %s, Literal: %+v }\n", enumString[t.tType], t.lexeme, t.literal)
}
