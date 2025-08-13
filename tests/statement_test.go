package tests

import (
	"gqlite/Parser"
	"gqlite/REPL"
	"testing"
	"unicode/utf8"
)

func TestPrepareStatement(t *testing.T) {
	commands := []string{
		"insert 1 user1 user1@email.com",
		"select",
		".play",
	}
	expected := []any{
		Parser.STATEMENT_PREPARE_SUCCESS,
		Parser.STATEMENT_PREPARE_SUCCESS,
		Parser.META_UNRECOGNIZED,
	}

	statement := &Parser.Statement{}

	for i, cmd := range commands {
		buf := &REPL.InputBuffer{
			Buffer:       cmd,
			BufferLength: len(cmd),
			InputLength:  utf8.RuneCountInString(cmd),
		}

		var res any
		if i <= 1 {
			res = statement.PrepareStatement(buf)
		} else {
			res = Parser.ExecMetaCommand(buf)
		}

		if res != expected[i] {
			t.Errorf("command %q: got %v, want %v", cmd, res, expected[i])
		}
	}
}

func TestStatementType(t *testing.T) {
	
}
