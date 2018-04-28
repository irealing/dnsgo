package layer

import (
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

type Question struct {
	QName string
	Type  QType
	Class uint16
}

func (q *Question) String() string {
	return fmt.Sprintf("Question(QName:%s, Type:%v,Class:%v)", q.QName, q.Type, q.Class)
}

func (q *Question) encodeDomain() []byte {
	return encodeDomain(q.QName)
}

type Answer struct {
	Name  string
	Type  QType
	Class uint16
	TTL   uint32
	RDLen uint16
	RData []byte
}

func (an *Answer) String() string {
	f := "Answer(Name=%s ,Type=%d ,Class=%d ,TTL=%d ,RDLen=%d, RData=...)"
	return fmt.Sprintf(f, an.Name, an.Type, an.Class, an.TTL, an.RDLen)
}
