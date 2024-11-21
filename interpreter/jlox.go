package interpreter

import (
	"log"
	"os"
)

type interpreter interface {
	RunFile(filePath string)
	RunPrompt(line string)
}

func RunFile(filePath string) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Panicf(err.Error())
	}
	Run(string(f))
}
func RunPrompt(line string) {}
