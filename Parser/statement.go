package Parser

import (
	"fmt"
	"gqlite/REPL"
	"gqlite/storage"
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

type StatementExecRes uint

const (
	EXECUTE_TABLE_FULL StatementExecRes = iota
	EXECUTED_SUCCESSFULLY
)

type Statement struct {
	SType StatementType
	Row   storage.Row
}

func trimSpaces(str string) string {
	return strings.TrimSpace(str)
}

func (statement *Statement) PrepareStatement(IB *REPL.InputBuffer) StatementType {
	if strings.HasPrefix(trimSpaces(IB.Buffer), "select") {
		statement.SType = SELECT_STATEMENT
		return STATEMENT_PREPARE_SUCCESS
	} else if strings.HasPrefix(trimSpaces(IB.Buffer), "insert") {
		statement.SType = INSERT_STATEMENT
		var username, email string
		n, err := fmt.Sscanf(IB.Buffer, "insert %d %s %s", &statement.Row.Id, &username, &email)
		if err != nil || n < 3 {
			return STATEMENT_PREPARE_ERROR
		}
		if len(username) > storage.COLUMN_USERNAME_SIZE {
			fmt.Println("Error: username is too long.")
			return STATEMENT_PREPARE_ERROR
		}
		if len(email) > storage.COLUMN_EMAIL_SIZE {
			fmt.Println("Error: email is too long.")
			return STATEMENT_PREPARE_ERROR
		}
		copy(statement.Row.Username[:], username)
		copy(statement.Row.Email[:], email)
		return STATEMENT_PREPARE_SUCCESS
	} else {
		return STATEMENT_UNRECOGNIZED
	}
}

func (statement *Statement) ExecInsert(table *storage.Table) StatementExecRes {
	if table.RowsNum >= storage.TABLE_MAX_ROWS {
		return EXECUTE_TABLE_FULL
	}
	rowToInsert := &statement.Row
	rowToInsert.Serialize(table.RowSlot(table.RowsNum))
	table.RowsNum += 1
	return EXECUTED_SUCCESSFULLY
}

func (statement *Statement) ExecSelect(table *storage.Table) StatementExecRes {
	var row storage.Row
	for i := 0; uint(i) < table.RowsNum; i++ {
		row.Deserialize(table.RowSlot(uint(i)))
		row.PrintRow()
	}
	return EXECUTED_SUCCESSFULLY
}

func (statement *Statement) ExecuteStatement(table *storage.Table) StatementExecRes {
	switch statement.SType {
	case SELECT_STATEMENT:
		return statement.ExecSelect(table)
	case INSERT_STATEMENT:
		return statement.ExecInsert(table)
	default:
		panic("unhandled default case")
	}
}
