package Parser

import (
	"gqlite/REPL"
	"os"
	"strings"
	"unsafe"
)

type MetaCommandResult int

const (
	COLUMN_USERNAME_SIZE uint = 32
	COLUMN_EMEAIL_SIZE   uint = 255
)

const (
	META_SUCCESSFUL MetaCommandResult = iota
	META_UNRECOGNIZED
)

var (
	ID_SIZE         uint = uint(SizeOfAttr(Row{}.id))
	USERNAME_SIZE   uint = uint(SizeOfAttr(Row{}.username))
	EMAIL_SIZE           = uint(SizeOfAttr(Row{}.email))
	ID_OFFSET       uint = 0
	USERNAME_OFFSET uint = ID_OFFSET + ID_SIZE
	EMAIL_OFFSET    uint = USERNAME_SIZE + USERNAME_OFFSET
	ROW_SIZE        uint = ID_SIZE + USERNAME_SIZE + EMAIL_SIZE
)

type Row struct {
	id       int
	username [COLUMN_USERNAME_SIZE]byte
	email    [COLUMN_EMEAIL_SIZE]byte
}

func (r *Row) serialize(dest []byte) {
	copy(dest[ID_OFFSET:ID_OFFSET+ID_SIZE],
		(*[1 << 30]byte)(unsafe.Pointer(&r.id))[:ID_SIZE:ID_SIZE])
	copy(dest[USERNAME_OFFSET:USERNAME_OFFSET+USERNAME_SIZE], (*[1 << 30]byte)(unsafe.Pointer(&r.username))[:USERNAME_SIZE:USERNAME_SIZE])
	copy(dest[EMAIL_OFFSET:EMAIL_OFFSET+EMAIL_SIZE],
		(*[1 << 30]byte)(unsafe.Pointer(&r.email))[:EMAIL_SIZE:EMAIL_SIZE])
}

func ExecMetaCommand(IB *REPL.InputBuffer) MetaCommandResult {
	if strings.Compare(IB.Buffer, ".exit") == 0 {
		os.Exit(0)
		return META_SUCCESSFUL
	} else {
		return META_UNRECOGNIZED
	}
}

func SizeOfAttr(x any) uintptr {
	return unsafe.Sizeof(x)
}
