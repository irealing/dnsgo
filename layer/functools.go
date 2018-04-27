package layer

import (
	"encoding/binary"
	"bytes"
	"strings"
)

func encodeU16(n uint16) []byte {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, n)
	return bs
}
func decodeU16(bs []byte) uint16 {
	return binary.BigEndian.Uint16(bs)
}

func encodeDomain(domain string) []byte {
	arr := strings.Split(domain, ".")
	buf := &bytes.Buffer{}
	for i := 0; i < len(arr); i++ {
		fragment := arr[i]
		if len(fragment) < 1 {
			continue
		}
		bin := []byte(fragment)
		buf.WriteByte(byte(len(bin)))
		buf.Write(bin)
	}
	buf.WriteRune(0)
	return buf.Bytes()
}
