package lexer

import (
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"
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
	fmt.Println("Start Scan")
	fmt.Printf("Source length: %d\n", utf8.RuneCountInString(s.source))
	for {
		if s.isAtEnd() {
			break
		}
		// fmt.Printf("Current position %d\n", s.current)
		// Offsets where start is the first character of the lexeme
		// current is the current character being read
		s.start = s.current
		s.scanToken()

	}
	s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))

	return s.tokens
}

/*
Scans single lexeme character

If the current character in the lexeme matches any of the patterns in the switch statement create a new token based on the lexeme from the character stream.
*/
func (s *Scanner) scanToken() {
	c := s.advance()
	// fmt.Printf("Scanned Token %s\n", c)

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

		// Handling multi-character lexems
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

			Calling peek() checks next character if it is the EOF it returns 0 else next character
			Calling isAtEnd() which returns a bool if EOF TRUE else FALSE

		*/

		if s.match("/") {
			fmt.Println(c)
			// If Comment do
			for {
				if s.peek() != "\n" && !s.isAtEnd() {
					s.advance()
				}
				break
			}
		} else {
			s.addSingleToken(SLASH)
		}
		break

	case " ":
		break
	case "\r":
		break
	case "\t":
		break
	case "\n":
		s.line++
		break
		// Handling literals
	case `"`:
		s.string(`"`)
		break
	case "'":
		s.string("'")
		break

	case "0":
		fmt.Println("End of file")
		break
	default:
		if s.isDigit(c) {
			s.number()
			break
		}
		fmt.Printf("Unexpected character %s in line %d\n", c, s.line)
		break
	}
}

func (s *Scanner) string(punc string) {

	for {
		if s.peek() != punc && !s.isAtEnd() {
			if s.peek() == "\n" {
				s.line++
			}
			s.advance()
		} else {
			break
		}
	}

	if s.isAtEnd() {
		fmt.Printf("line %d Unterminated String\n", s.line)
	}
	s.advance()                                         // consume the last `"`
	s.addToken(STRING, s.source[s.start+1:s.current-1]) // trimming quotes
}

func (s *Scanner) number() {

	for {
		if s.peek() != "." && !s.isAtEnd() {
			s.advance()
		} else {
			break
		}
	}

	if s.peek() == "." && s.isDigit(s.peekNext()) {
		for {
			if s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			} else {
				break
			}
		}
	}

	fmt.Printf("Number %s", s.source[s.start:s.current])
	i, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		log.Println(err.Error())
	}
	s.addToken(NUMBER, i)
}

func (s *Scanner) isDigit(c string) bool {

	if c >= "0" && c <= "9" {
		return true
	}
	return false
}

// To Peek at the current character
func (s *Scanner) peek() string {
	if s.isAtEnd() {
		// cant parse "/0" escape sequence
		// gonna have to manually interpret this myself
		return "0"
	}
	return string(s.source[s.current])
}

// To Peek at the next character from the current character
func (s *Scanner) peekNext() string {
	if s.current+1 >= len(s.source)-1 {
		// cant parse "/0" escape sequence
		// gonna have to manually interpret this myself
		return "0"
	}
	return string(s.source[s.current+1])
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)-1
}

/*
increments the current position and returns the character
in the current position of the substring from the starting lexeme

Advance which "advances" the pointer to the next position to be read
*/
func (s *Scanner) advance() string {
	defer func() { s.current++ }()
	return string(s.source[s.current])
}

// Wrapper for single lexeme character
func (s *Scanner) addSingleToken(tokenType int) {
	s.addToken(tokenType, nil)
}

func (s *Scanner) addToken(tokenType int, literal any) {
	var byteText = s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tokenType, byteText, literal, s.line))
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
