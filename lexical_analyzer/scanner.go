package lexicalanalyzer

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) Scanner {
	return Scanner{source: source, line: 1}
}

func (s *Scanner) ScanTokens() []Token {
	// loop until end of file
	for {
		if s.isAtEnd() {
			break
		}
		// Offsets where start is the first character of the lexeme
		// current is the current character being read
		s.start = s.current
		s.scanToken()

	}
	s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))

	return []Token{}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

/*
increments the current position and returns the character
in the current position of the substring from the starting lexeme

Advance which "advances" the pointer to the next position to be read
*/
func (s *Scanner) advance() string {
	s.current++
	return string(s.source[s.current])
}

/*
Scans single lexeme character

If the current character in the lexeme matches any of the patterns in the switch statement create a new token based on the lexeme from the character stream.
*/
func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case "(":
		s.addSingleToken(LEFT_PAREN)
		break
	case ")":
		s.addSingleToken(RIGHT_PAREN)
		break
	case "{":
		s.addSingleToken(LEFT_BRACE)
		break
	case "}":
		s.addSingleToken(RIGHT_BRACE)
		break
	case ",":
		s.addSingleToken(COMMA)
		break
	case ".":
		s.addSingleToken(DOT)
		break
	case "-":
		s.addSingleToken(MINUS)
		break
	case "+":
		s.addSingleToken(PLUS)
		break
	case ";":
		s.addSingleToken(SEMICOLON)
		break
	case "*":
		s.addSingleToken(STAR)
		break
	}
}

// Wrapper for single lexeme character
func (s *Scanner) addSingleToken(tokenType int) {
	s.addToken(tokenType, nil)
}

func (s *Scanner) addToken(tokenType int, literal any) {
	byteText := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tokenType, string(byteText), literal, s.line))

}
