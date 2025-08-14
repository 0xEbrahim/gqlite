package storage

type Cursor struct {
	Table      *Table
	RowNum     uint
	EndOfTable bool
}

func TopCursor(table *Table) *Cursor {
	cursor := &Cursor{Table: table, RowNum: 0, EndOfTable: table.RowsNum == 0}
	return cursor
}

func BottomCursor(table *Table) *Cursor {
	cursor := &Cursor{Table: table, RowNum: table.RowsNum, EndOfTable: true}
	return cursor
}

func (cursor *Cursor) Advance() {
	cursor.RowNum += 1
	if cursor.RowNum == cursor.Table.RowsNum {
		cursor.EndOfTable = true
	}
}

func (cursor *Cursor) CursorValue() []byte {
	rowNum := cursor.RowNum
	table := cursor.Table
	if rowNum >= TABLE_MAX_ROWS || rowNum >= table.RowsNum {
		return nil
	}
	pageNum := rowNum / ROWS_PER_PAGE
	page := table.Pager.getPage(pageNum)
	rowOffset := rowNum % ROWS_PER_PAGE
	byteOffset := rowOffset * ROW_SIZE
	return page[byteOffset : byteOffset+ROW_SIZE]
}
