package REPL

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	IB.Buffer = str
	fmt.Println(IB.Buffer)
	if err != nil {
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
