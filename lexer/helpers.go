package lexer

func isDigit(c string) bool {

	if c >= "0" && c <= "9" {
		return true
	}
	return false
}

func isAlphabet(c string) bool {
	if c >= "a" && c <= "z" || c >= "A" && c <= "Z" || c == "_" {
		return true
	}
	return false
}

func isAlphaNum(c string) bool {
	if isDigit(c) || isAlphabet(c) {
		return true
	}
	return false
}
