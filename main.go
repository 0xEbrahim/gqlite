package main

import (
	"fmt"
	"gqlite/REPL"
	"os"
	"strings"
)

func main() {
	inputBuffer := REPL.NewInputBuffer()
	for {
		REPL.PrintPrompt()
		inputBuffer.ReadInput()
		if strings.Compare(inputBuffer.Buffer, ".exit") == 0 {
			os.Exit(0)
		} else {
			fmt.Println("Unrecognized command: " + inputBuffer.Buffer)
		}
	}
}
