package storage

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

var (
	ID_SIZE         uint = uint(SizeOfAttr(Row{}.Id))
	USERNAME_SIZE   uint = uint(SizeOfAttr(Row{}.Username))
	EMAIL_SIZE           = uint(SizeOfAttr(Row{}.Email))
	ID_OFFSET       uint = 0
	USERNAME_OFFSET uint = ID_OFFSET + ID_SIZE
	EMAIL_OFFSET    uint = USERNAME_SIZE + USERNAME_OFFSET
	ROW_SIZE        uint = ID_SIZE + USERNAME_SIZE + EMAIL_SIZE
)

const (
	COLUMN_USERNAME_SIZE uint = 32
	COLUMN_EMAIL_SIZE    uint = 255
)

type Row struct {
	Id       int
	Username [COLUMN_USERNAME_SIZE]byte
	Email    [COLUMN_EMAIL_SIZE]byte
}

func (source *Row) Serialize(dest []byte) {
	binary.LittleEndian.PutUint32(dest[ID_OFFSET:ID_OFFSET+ID_SIZE], uint32(source.Id))
	copy(dest[USERNAME_OFFSET:USERNAME_OFFSET+USERNAME_SIZE], source.Username[:])
	copy(dest[EMAIL_OFFSET:EMAIL_OFFSET+EMAIL_SIZE], source.Email[:])
}

func (dest *Row) Deserialize(source []byte) {
	copy((*[1 << 30]byte)(unsafe.Pointer(&dest.Id))[:ID_SIZE:ID_SIZE], source[ID_OFFSET:ID_OFFSET+ID_SIZE])
	dest.Id = int(binary.LittleEndian.Uint32(source[ID_OFFSET : ID_OFFSET+ID_SIZE]))
	copy(dest.Username[:], source[USERNAME_OFFSET:USERNAME_OFFSET+USERNAME_SIZE])
	copy(dest.Email[:], source[EMAIL_OFFSET:EMAIL_OFFSET+EMAIL_SIZE])
}

func (row *Row) PrintRow() {
	fmt.Printf("(%d, %s, %s)\n", row.Id, row.Username, row.Email)
}

func SizeOfAttr(x any) uintptr {
	return unsafe.Sizeof(x)
}
