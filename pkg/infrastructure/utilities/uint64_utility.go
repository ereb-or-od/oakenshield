package utilities

import "encoding/binary"

func Bytes(id uint64) []byte{
	uid := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(uid, id)
	return uid
}
