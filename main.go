package main

import (
	"fmt"
	"golox/interpreter"
	"os"
	"strings"
)

var args = os.Args

func main() {

	fmt.Println("Hello lox language")
	if !strings.Contains(args[1], ".lox") {
		fmt.Printf("Unknown file `%s`\n", args[1])
		fmt.Println("Usage: jlox [script]")
		os.Exit(1)
	} else if len(args) == 2 {
		fmt.Println("Eval source file")
		interpreter.RunFile(args[1])
	} else {
		fmt.Println("REPL")
	}
}
