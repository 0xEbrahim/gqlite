package db

import (
	"gqlite/storage"
	"log"
)

func OpenDB(fileName string) *storage.Table {
	pager := storage.PagerOpen(fileName)
	numberOfRows := pager.FileLength / storage.ROW_SIZE
	table := &storage.Table{RowsNum: numberOfRows, Pager: pager}
	return table
}

func CloseDB(table *storage.Table) {
	pager := table.Pager
	numOfFullPages := table.RowsNum / storage.ROWS_PER_PAGE
	for i := 0; uint(i) < numOfFullPages; i++ {
		if pager.Pages[i] == nil {
			continue
		}
		pager.PagerFlush(uint(i), storage.PAGE_SIZE)
		pager.Pages[i] = nil
	}
	numOfAdditionalRows := table.RowsNum % storage.ROWS_PER_PAGE
	if numOfAdditionalRows > 0 {
		if pager.Pages[numOfFullPages] != nil {
			pager.PagerFlush(numOfFullPages, numOfAdditionalRows*storage.ROW_SIZE)
			pager.Pages[numOfFullPages] = nil
		}
	}
	err := pager.File.Close()
	if err != nil {
		log.Fatal("Error while close descriptor file")
	}
	for i := 0; uint(i) < storage.MAX_PAGES_IN_TABLE; i++ {
		pager.Pages[i] = nil
	}
}
