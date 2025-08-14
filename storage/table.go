package storage

const (
	PAGE_SIZE          uint = 4096
	MAX_PAGES_IN_TABLE uint = 1000
)

var (
	ROWS_PER_PAGE  uint = PAGE_SIZE / ROW_SIZE
	TABLE_MAX_ROWS uint = ROWS_PER_PAGE * MAX_PAGES_IN_TABLE
)

type Table struct {
	RowsNum uint
	Pager   *Pager
}

func (tbl *Table) RowSlot(rowNum uint) []byte {
	if rowNum >= TABLE_MAX_ROWS {
		return nil
	}
	pageNum := rowNum / ROWS_PER_PAGE
	page := tbl.Pager.getPage(pageNum)
	rowOffset := rowNum % ROWS_PER_PAGE
	byteOffset := rowOffset * ROW_SIZE
	return page[byteOffset : byteOffset+ROW_SIZE]
}
