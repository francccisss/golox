package main

import (
	"fmt"
	"golox/interpreter"
	"os"
	"strings"
)

var args = os.Args

// Describing a pattern through regular expression, and identifying that as a type of token

func main() {
	if len(args) > 1 && !strings.Contains(args[1], ".lox") {
		fmt.Printf("Unknown file `%s`\n", args[1])
		fmt.Println("Usage: jlox [script]")
		os.Exit(1)
	} else if len(args) == 2 {
		fmt.Println("Eval source file")
		f, _ := os.ReadFile(args[1])
		fmt.Println(string(f))
		err := interpreter.RunFile(args[1])
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("REPL")
		err := interpreter.RunPrompt()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
