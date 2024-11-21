package interpreter

import (
	"bufio"
	"fmt"
	lexicalanalyzer "golox/lexical_analyzer"
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
	lexicalanalyzer.Run(string(f))
	return nil
}
func RunPrompt() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("> ")
		lexicalanalyzer.Run(text)
		if text == "exit" {
			return nil
		}
	}
	return nil
}
