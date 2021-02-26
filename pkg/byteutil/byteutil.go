package byteutil

import (
	"bytes"
	"encoding/binary"
)

// IntToBytes int类型转bytes
func IntToBytes(n int) []byte {
	buf := bytes.NewBuffer([]byte{})
	x := int64(n)
	binary.Write(buf, binary.BigEndian, x)
	return buf.Bytes()
}

// BytesToInt bytes转int类型
func BytesToInt(b []byte) int {
	buf := bytes.NewBuffer(b)
	var n int64
	binary.Read(buf, binary.BigEndian, &n)
	return int(n)
}
