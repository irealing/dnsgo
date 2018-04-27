package layer

import (
	"bytes"
	"encoding/binary"
)

type QueryEncoder struct {
}

func (qe *QueryEncoder) Encode(q *Query) []byte {
	buf := bytes.NewBuffer(qe.encodeHeader(q.Header))
	if q.Questions != nil && len(q.Questions) > 0 {
		for _, question := range q.Questions {
			buf.Write(qe.encodeQuestion(question))
		}
	}
	return buf.Bytes()
}
func (qe *QueryEncoder) encodeQuestion(q *Question) []byte {
	buf := bytes.NewBuffer(encodeDomain(q.QName))
	ia := make([]byte, 2)
	binary.BigEndian.PutUint16(ia, uint16(q.Type))
	buf.Write(ia)
	binary.BigEndian.PutUint16(ia, q.Class)
	buf.Write(ia)
	return buf.Bytes()
}
func (qe *QueryEncoder) encodeHeader(h *DNSHeader) []byte {
	buf := &bytes.Buffer{}
	cache := make([]byte, 2)
	arr := []uint16{h.ID, uint16(h.Opt), h.QDCount, h.AnCount, h.NsCount, h.ArCount}
	for i := 0; i < len(arr); i++ {
		binary.BigEndian.PutUint16(cache, arr[i])
		buf.Write(cache)
	}
	return buf.Bytes()
}
