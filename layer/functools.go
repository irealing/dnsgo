package layer

import "encoding/binary"

func encodeU16(n uint16) []byte {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, n)
	return bs
}
func decodeU16(bs []byte) uint16 {
	return binary.BigEndian.Uint16(bs)
}
