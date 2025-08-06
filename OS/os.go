package OS

import (
	"log"
	"os"
	"path/filepath"
	"syscall"
)

const (
	READ_ONLY       int = syscall.O_RDONLY
	WRITE_ONLY      int = syscall.O_WRONLY
	READ_WRITE_ONLY int = syscall.O_RDWR
	CREATE_         int = syscall.O_CREAT
)

type gqliteFilename = string

type XOS struct {
	MFileSz      int64
	MxFilePathSz int64
	XosName      string
}

func (xos *XOS) XOpen(zName gqliteFilename, flags int) (GqliteFile, error) {
	fullPath, err := xos.XFullPathName(zName)
	if err != nil {
		log.Fatal("Can't reach the file path.")
	}
	f, err := os.OpenFile(fullPath, flags, 0)
	gqliteFile := GqliteFile{File: f}
	return gqliteFile, err
}

func (xos *XOS) xDelete() {}

func (xos *XOS) xAccess() {}

func (xos *XOS) XFullPathName(fileName string) (string, error) {
	cwd, err := os.Getwd()
	cwd = filepath.Join(cwd, fileName)
	return cwd, err
}
