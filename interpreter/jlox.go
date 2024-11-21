package interpreter

import (
	"bufio"
	"fmt"
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
	fmt.Printf("> ")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("> ")
		lexer.Run(text)
		if text == "exit" {
			return nil
		}
	}
	return nil
}
