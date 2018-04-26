package layer

import (
	"bytes"
	"strings"
	"encoding/binary"
	"fmt"
)

var globalID uint16 = 0

type Query struct {
	Header    *DNSHeader
	Questions []*Question
	Answers   []*Answer
}

func SimpleQuery(domain string) *Query {
	q := new(Query)
	q.Header = &DNSHeader{
		ID:      globalID,
		Opt:     NewOption(),
		QDCount: 1,
	}
	q.Questions = []*Question{
		{
			QName: domain,
			Type:  Adress,
			Class: 1,
		},
	}
	return q
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
func (h *DNSHeader) Bytes() []byte {
	buf := bytes.Buffer{}
	arr := []uint16{h.ID, uint16(h.Opt), h.QDCount, h.AnCount, h.NsCount, h.ArCount}
	bits := make([]byte, 2)
	for i := 0; i < len(arr); i++ {
		binary.BigEndian.PutUint16(bits, arr[i])
		buf.Write(bits)
	}
	return buf.Bytes()
}

type Question struct {
	QName string
	Type  QType
	Class uint16
}

func (q *Question) String() string {
	return fmt.Sprintf("Question(QName:%s, Type:%v,Class:%v)", q.QName, q.Type, q.Class)
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
		fragment := keys[i]
		if len(fragment) < 1 {
			continue
		}
		arr := []byte(fragment)
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
