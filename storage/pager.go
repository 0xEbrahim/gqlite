package storage

import (
	"io"
	"log"
	"os"
)

type Pager struct {
	File       *os.File
	FileLength uint
	Pages      [MAX_PAGES_IN_TABLE][]byte
}

func PagerOpen(fileName string) *Pager {
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 644)
	if err != nil {
		log.Fatal("Unable to open database file")
	}
	info, err := fd.Stat()
	sz := info.Size()
	var pages [MAX_PAGES_IN_TABLE][]byte
	for i := 0; uint(i) < MAX_PAGES_IN_TABLE; i++ {
		pages[i] = nil
	}
	pager := &Pager{File: fd, FileLength: uint(sz), Pages: pages}
	return pager
}

func (pager *Pager) getPage(pageNum uint) []byte {
	if pageNum >= MAX_PAGES_IN_TABLE {
		log.Fatal("You are trying to access out of bound page.")
	}
	if pager.Pages[pageNum] == nil {
		page := make([]byte, PAGE_SIZE)
		numberOfPages := pager.FileLength / PAGE_SIZE
		if pager.FileLength%PAGE_SIZE != 0 {
			numberOfPages += 1
		}
		if pageNum <= numberOfPages {
			_, err := pager.File.ReadAt(page, int64(pageNum*PAGE_SIZE))
			if err != io.EOF {
				log.Fatal("Error while reading the page")
			}
		}
		pager.Pages[pageNum] = page
	}
	return pager.Pages[pageNum]
}

func (pager *Pager) PagerFlush(pageNum uint, size uint) {
	if pager.Pages[pageNum] == nil {
		log.Fatal("Cannot flush a null page")
	}
	_, err := pager.File.WriteAt(pager.Pages[pageNum][:size], int64(PAGE_SIZE*pageNum))
	if err != nil {
		log.Fatal("Error while flushing the page.")
	}
}
