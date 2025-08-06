package OS

import (
	"fmt"
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
	err = xos.XFileExists(zName)
	if err != nil {
		return GqliteFile{}, fmt.Errorf("file does not exist: %w", err)
	}
	f, err := os.OpenFile(fullPath, flags, 0)
	gqliteFile := GqliteFile{File: f}
	return gqliteFile, err
}

func (xos *XOS) xDelete(zName gqliteFilename, dirSync bool) error {

	if err := xos.XFileExists(zName); err != nil {
		return fmt.Errorf("file does not exist: %w", err)
	}
	if err := os.Remove(zName); err != nil {
		return fmt.Errorf("error while deleting the file: %w", err)
	}
	if dirSync {
		dir := filepath.Dir(string(zName))
		f, err := os.Open(dir)
		if err != nil {
			return fmt.Errorf("error while opening the directory: %w", err)
		}
		defer func(f *os.File) {
			if Err := f.Close(); Err != nil {
				log.Fatal("error while closing the directory")
			}
		}(f)
		if err := f.Sync(); err != nil {
			return fmt.Errorf("failed to sync directory: %w", err)
		}
	}
	return nil
}

func (xos *XOS) XFileExists(zName gqliteFilename) error {
	_, err := os.Stat(zName)
	return err
}

func (xos *XOS) xAccess() {}

func (xos *XOS) XFullPathName(fileName string) (string, error) {
	cwd, err := os.Getwd()
	cwd = filepath.Join(cwd, fileName)
	return cwd, err
}
