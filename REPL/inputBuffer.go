package REPL

import (
	"cmp"
	"fmt"
	"log"
	"unicode/utf8"
)

type InputBuffer struct {
	Buffer       string
	InputLength  int
	BufferLength int
}

func NewInputBuffer() *InputBuffer {
	buffer := &InputBuffer{
		Buffer:       "",
		InputLength:  0,
		BufferLength: 0,
	}
	return buffer
}

// ReadInput IB -> Input Buffer
func (IB *InputBuffer) ReadInput() {
	n, err := fmt.Scanln(&IB.Buffer)
	if cmp.Or(err != nil, n <= 0) {
		log.Fatal("Error reading input \n")
		return
	}
	IB.BufferLength = len(IB.Buffer)
	IB.InputLength = utf8.RuneCountInString(IB.Buffer)
	return
}

func PrintPrompt() {
	fmt.Print("gqlite> ")
}
