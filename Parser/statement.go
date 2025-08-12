package Parser

import (
	"gqlite/REPL"
	"strings"
)

type StatementType int

const (
	INSERT_STATEMENT StatementType = iota
	SELECT_STATEMENT
	STATEMENT_PREPARE_SUCCESS
	STATEMENT_UNRECOGNIZED
)

type Statement struct {
	SType StatementType
}

func (statement *Statement) PrepareStatement(IB *REPL.InputBuffer) StatementType {
	if strings.Compare(strings.ToLower(IB.Buffer), "select") == 0 {
		statement.SType = SELECT_STATEMENT
		return STATEMENT_PREPARE_SUCCESS
	} else if strings.Compare(strings.ToLower(IB.Buffer), "insert") == 0 {
		statement.SType = INSERT_STATEMENT
		return STATEMENT_PREPARE_SUCCESS
	} else {
		return STATEMENT_UNRECOGNIZED
	}
}

func (statement *Statement) Exec() {

}
