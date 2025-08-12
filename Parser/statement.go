package Parser

import (
	"fmt"
	"gqlite/REPL"
	"strings"
)

type StatementType int

const (
	INSERT_STATEMENT StatementType = iota
	SELECT_STATEMENT
	STATEMENT_PREPARE_SUCCESS
	STATEMENT_UNRECOGNIZED
	STATEMENT_PREPARE_ERROR
)

type Statement struct {
	SType StatementType
	row   Row
}

func (statement *Statement) PrepareStatement(IB *REPL.InputBuffer) StatementType {
	if strings.Compare(strings.ToLower(IB.Buffer), "select") == 0 {
		statement.SType = SELECT_STATEMENT
		return STATEMENT_PREPARE_SUCCESS
	} else if strings.Compare(strings.ToLower(IB.Buffer), "insert") == 0 {
		statement.SType = INSERT_STATEMENT
		n, err := fmt.Sscan(IB.Buffer, "insert %d %s %s", &statement.row.id, &statement.row.username, &statement.row.email)
		if err != nil || n < 3 {
			return STATEMENT_PREPARE_ERROR
		}
		return STATEMENT_PREPARE_SUCCESS
	} else {
		return STATEMENT_UNRECOGNIZED
	}
}

func (statement *Statement) Exec() {
	switch statement.SType {
	case SELECT_STATEMENT:
		println("This is a select statement")
	case INSERT_STATEMENT:
		println("This is an insert statement")
	default:
		panic("unhandled default case")
	}
}
