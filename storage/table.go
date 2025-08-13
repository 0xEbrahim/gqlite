package storage

const (
	PAGE_SIZE          uint = 4096
	MAX_PAGES_IN_TABLE uint = 100
)

var (
	ROWS_PER_PAGE  uint = PAGE_SIZE / ROW_SIZE
	TABLE_MAX_ROWS uint = ROWS_PER_PAGE * MAX_PAGES_IN_TABLE
)

type Table struct {
	RowsNum uint
	Pages   [MAX_PAGES_IN_TABLE][]byte
}

func (tbl *Table) RowSlot(rowNum uint) []byte {
	pageNum := rowNum / ROWS_PER_PAGE
	page := tbl.Pages[pageNum]
	if page == nil {
		page = make([]byte, PAGE_SIZE)
		tbl.Pages[pageNum] = page
	}
	rowOffset := rowNum % ROWS_PER_PAGE
	byteOffset := rowOffset * ROW_SIZE
	return page[byteOffset : byteOffset+ROW_SIZE]
}

func NewTable() *Table {
	var pages [MAX_PAGES_IN_TABLE][]byte
	table := &Table{RowsNum: 0, Pages: pages}
	return table
}
