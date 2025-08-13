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
	row   storage.Row
}

func strToLower(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

func (statement *Statement) PrepareStatement(IB *REPL.InputBuffer) StatementType {
	if strings.Compare(strToLower(IB.Buffer), "select") == 0 {
		statement.SType = SELECT_STATEMENT
		return STATEMENT_PREPARE_SUCCESS
	} else if strings.Compare(strToLower(IB.Buffer), "insert") == 0 {
		statement.SType = INSERT_STATEMENT
		n, err := fmt.Sscanf(IB.Buffer, "insert %d %s %s", &statement.row.Id, &statement.row.Username, &statement.row.Email)
		if err != nil || n < 3 {
			return STATEMENT_PREPARE_ERROR
		}
		return STATEMENT_PREPARE_SUCCESS
	} else {
		return STATEMENT_UNRECOGNIZED
	}
}

func (statement *Statement) ExecInsert(table *storage.Table) StatementExecRes {
	if table.RowsNum >= storage.TABLE_MAX_ROWS {
		return EXECUTE_TABLE_FULL
	}
	rowToInsert := &statement.row
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

func (statement *Statement) ExecuteStatement() {

}
