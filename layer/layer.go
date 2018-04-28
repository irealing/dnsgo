package layer

import (
	"fmt"
)

type Query struct {
	Header    *DNSHeader
	Questions []*Question
	Answers   []*Answer
}

type DNSHeader struct {
	ID  uint16
	Opt Option
	//QDCount 查询数
	QDCount uint16
	//AnCount 应答数
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
