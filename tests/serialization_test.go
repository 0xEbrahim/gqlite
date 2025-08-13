package tests

import (
	"gqlite/Parser"
	"gqlite/REPL"
	"gqlite/storage"
	"testing"
	"unicode/utf8"
)

func TestRowSerialization(t *testing.T) {
	cmd := []string{"insert 1 user user@email.com", "insert 2 khaled khaled@email.com", "insert 3 ebrahim ebrahim@email.com"}
	table := storage.NewTable()
	statement := &Parser.Statement{}
	for i := 0; i < len(cmd); i++ {
		buf := &REPL.InputBuffer{
			Buffer:       cmd[i],
			BufferLength: len(cmd[i]),
			InputLength:  utf8.RuneCountInString(cmd[i]),
		}
		res := statement.PrepareStatement(buf)
		if res != Parser.STATEMENT_PREPARE_SUCCESS {
			t.Errorf("error preparing statement")
		}
		res1 := statement.ExecuteStatement(table)
		if res1 != Parser.EXECUTED_SUCCESSFULLY {
			t.Errorf("Error executing insert query")
		}
	}
	cmdd := "select"
	buf := &REPL.InputBuffer{
		Buffer:       cmdd,
		BufferLength: len(cmdd),
		InputLength:  utf8.RuneCountInString(cmdd),
	}
	res := statement.PrepareStatement(buf)
	if res != Parser.STATEMENT_PREPARE_SUCCESS {
		t.Errorf("error preparing statement")
	}
	res1 := statement.ExecuteStatement(table)
	if res1 != Parser.EXECUTED_SUCCESSFULLY {
		t.Errorf("Error executing insert query")
	}
}
