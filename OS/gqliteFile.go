package OS

import (
	"io"
	"os"
)

type GqliteFile struct {
	File *os.File
	Path string
}

func (gqlF *GqliteFile) XClose() error {
	err := gqlF.File.Close()
	gqlF.File = nil
	return err
}

func (gqlF *GqliteFile) XRead(p []byte, offset int64) (int, error) {
	_, err := gqlF.File.Seek(offset, 0)
	if err != nil {
		return 0, err
	}
	n, err := gqlF.File.Read(p)
	if n < len(p) {
		copy(p[n:], make([]byte, len(p)-n))
	}
	if err == io.EOF {
		err = nil
	}
	return len(p), err
}

func (gqlF *GqliteFile) XWrite(p []byte, offset int64) error {
	n, err := gqlF.File.WriteAt(p, offset)
	if n < len(p) {
		return io.ErrShortWrite
	}
	return err
}

func (gqlF *GqliteFile) XFileName() (string, error) {
	return gqlF.Path, nil
}

func (gqlF *GqliteFile) XFileSize() (int64, error) {
	info, err := gqlF.File.Stat()
	if err != nil {
		return 0, err
	}
	sz := info.Size()
	return sz, nil
}
