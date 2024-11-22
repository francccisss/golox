package interpreter

import (
	"bufio"
	"golox/lexer"
	"os"
)

type interpreter interface {
	RunFile(filePath string)
	RunPrompt(line string)
}

func RunFile(filePath string) error {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	lexer.Run(string(f))
	return nil
}
func RunPrompt() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			return nil
		}

		lexer.Run(text)
	}
	return nil
}
