package main

import (
	"fmt"
	"gqlite/Parser"
	"gqlite/REPL"
	"gqlite/storage"
)

func main() {
	table := storage.NewTable()
	inputBuffer := REPL.NewInputBuffer()
	for {
		REPL.PrintPrompt()
		inputBuffer.ReadInput()
		if inputBuffer.Buffer[0] == '.' {
			switch Parser.ExecMetaCommand(inputBuffer) {
			case Parser.META_SUCCESSFUL:
				continue
			case Parser.META_UNRECOGNIZED:
				fmt.Println("Unrecognized meta command: " + inputBuffer.Buffer)
				continue
			}
		}
		statement := &Parser.Statement{}
		switch statement.PrepareStatement(inputBuffer) {
		case Parser.STATEMENT_PREPARE_SUCCESS:
			break
		case Parser.STATEMENT_PREPARE_ERROR:
			fmt.Println("Error while preparing the query, please check the query again")
			continue
		case Parser.STATEMENT_UNRECOGNIZED:
			fmt.Println("Unrecognized keyword at the start of: " + inputBuffer.Buffer)
			continue
		default:
			panic("unhandled default case")
		}
		switch statement.ExecuteStatement(table) {
		case Parser.EXECUTED_SUCCESSFULLY:
			println("Executed.")
			continue
		case Parser.EXECUTE_TABLE_FULL:
			fmt.Println("cannot insert on this table because it is full")
			continue
		}

	}

}
