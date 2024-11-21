package lexer

import (
	"log"
	"math"
)

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
	case "!":
		ok := s.match("=")
		var tType = EQUAL_EQUAL
		if ok {
			tType = BANG_EQUAL
		}
		s.addSingleToken(tType)
		break
	case "=":
		ok := s.match("=")
		var tType = EQUAL
		if ok {
			tType = EQUAL_EQUAL
		}
		s.addSingleToken(tType)
		break
	case "<":
		ok := s.match("=")
		var tType = LESS
		if ok {
			tType = LESS_EQUAL
		}
		s.addSingleToken(tType)
		break
	case ">":
		ok := s.match("=")
		var tType = GREATER
		if ok {
			tType = GREATER_EQUAL
		}
		s.addSingleToken(tType)
		break

	case "/":
		/*
			This section handles "Comments" until it is EOF and is New line escape sequence

			In match(), if the next character is "/" which signifies as a comment, then we advance the current pointer in the source file

			Calling peek() checks if it is the EOF it returns 0 else next character
			Calling isAtEnd() which returns a bool if EOF TRUE else FALSE

			if peek() == new line (\n) and reading is EOF
			then we break else we keep advancing until we reach EOF and new line

		*/

		if s.match("/") {
			// If Comment do
			for {
				if s.peek() != "\n" && !s.isAtEnd() {
					s.advance()
				}
			}
		} else {
			s.addSingleToken(SLASH)
		}
		break
	default:
		log.Panicf("Unexpected character %s in line %d", c, s.line)
		break
	}
}

// To Peek one character at a time
func (s *Scanner) peek() string {
	if s.isAtEnd() {
		// cant parse "/0" escape sequence
		// gonna have to manually interpret this myself
		return "0"
	}
	return string(s.source[s.current])
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

// Wrapper for single lexeme character
func (s *Scanner) addSingleToken(tokenType int) {
	s.addToken(tokenType, nil)
}

func (s *Scanner) addToken(tokenType int, literal any) {
	byteText := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tokenType, string(byteText), literal, s.line))
}

/*
Checks the second character if it matches the expected character then advances
if it does match which means its a single lexeme
*/
func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if string(s.source[s.current+1]) != expected {
		return false
	}
	// since the current character is actually a match with the next expected
	// character which makes it a single lexeme or is atomic eg: !=, >=, <=
	// then we can just go to the next character
	s.current++
	return true
}
