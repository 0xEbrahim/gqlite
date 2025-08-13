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
	num_rows uint
	pages    [MAX_PAGES_IN_TABLE][]byte
}
