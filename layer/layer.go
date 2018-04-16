package layer

import (
	"bytes"
	"strings"
	"encoding/binary"
	"fmt"
)

type Query struct {
	Header    *DNSHeader
	Questions []Question
	Answers   []Answer
}
type DNSHeader struct {
	ID      uint16
	Opt     Option
	QDCount uint16
	AnCount uint16
	NsCount uint16
	ArCount uint16
}

func (h DNSHeader) String() string {
	f := "DNSHeader(ID= %d, Opt=(%v), QDCount= %d, AnCount= %d, NsCount= %d, ArCount= %d)"
	return fmt.Sprintf(f, h.ID, h.Opt.String(), h.QDCount, h.AnCount, h.NsCount, h.ArCount)
}

type Question struct {
	QName string
	Type  QType
	Class uint16
}

func (q *Question) Bytes() []byte {
	buf := bytes.Buffer{}
	buf.Write(q.encodeDomain())
	buf.Write(q.Type.Encode())
	ia := make([]byte, 2)
	binary.BigEndian.PutUint16(ia, q.Class)
	buf.Write(ia)
	return buf.Bytes()
}

func (q *Question) encodeDomain() []byte {
	c := bytes.Buffer{}
	keys := strings.Split(q.QName, ".")
	for i := 0; i < len(keys); i++ {
		arr := []byte(keys[i])
		c.WriteByte(byte(len(arr)))
		c.Write(arr)
	}
	c.WriteByte(0x00)
	return c.Bytes()
}

type Answer struct {
	Name  string
	Type  QType
	Class uint16
	TTL   uint32
	RDLen uint16
	RData []byte
}
