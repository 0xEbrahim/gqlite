package Parser

import (
	"gqlite/REPL"
	"gqlite/db"
	"gqlite/storage"
	"os"
	"strings"
)

type MetaCommandResult int

const (
	META_SUCCESSFUL MetaCommandResult = iota
	META_UNRECOGNIZED
)

func ExecMetaCommand(IB *REPL.InputBuffer, table *storage.Table) MetaCommandResult {
	if strings.Compare(trimSpaces(IB.Buffer), ".exit") == 0 {
		db.CloseDB(table)
		os.Exit(0)
		return META_SUCCESSFUL
	} else {
		return META_UNRECOGNIZED
	}
}
