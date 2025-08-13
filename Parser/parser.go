package Parser

import (
	"gqlite/REPL"
	"os"
	"strings"
	"unsafe"
)

type MetaCommandResult int

const (
	META_SUCCESSFUL MetaCommandResult = iota
	META_UNRECOGNIZED
)

func ExecMetaCommand(IB *REPL.InputBuffer) MetaCommandResult {
	if strings.Compare(strToLower(IB.Buffer), ".exit") == 0 {
		os.Exit(0)
		return META_SUCCESSFUL
	} else {
		return META_UNRECOGNIZED
	}
}
