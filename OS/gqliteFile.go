package OS

import (
	"os"
)

type GqliteFile struct {
	File *os.File
}

func (gqlF GqliteFile) XClose() error {
	err := gqlF.File.Close()
	return err
}
