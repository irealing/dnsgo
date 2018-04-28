package layer

import (
	"bytes"
	"encoding/binary"
	"io"
)

type queryEncoder struct {
}

func (qe *queryEncoder) Encode(q *Query) ([]byte, error) {
	bits := qe.encodeHeader(q.Header)
	curLen := len(bits)
	buf := bytes.NewBuffer(bits)
	qc := len(q.Questions)
	idx := make([]int, qc)
	if q.Questions != nil && len(q.Questions) > 0 {
		for i := 0; i < len(q.Questions); i++ {
			idx[i] = curLen
			bits = qe.encodeQuestion(q.Questions[i])
			curLen += len(bits)
			buf.Write(bits)
		}
	}
	if q.Answers != nil && len(q.Answers) > 0 {
		bits, err := qe.encodeAnswers(q.Answers, idx)
		if err != nil {
			return nil, err
		}
		buf.Write(bits)
	}
	return buf.Bytes(), nil
}
func (qe *queryEncoder) encodeQuestion(q *Question) []byte {
	buf := bytes.NewBuffer(encodeDomain(q.QName))
	ia := make([]byte, 2)
	binary.BigEndian.PutUint16(ia, uint16(q.Type))
	buf.Write(ia)
	binary.BigEndian.PutUint16(ia, q.Class)
	buf.Write(ia)
	return buf.Bytes()
}
func (qe *queryEncoder) encodeHeader(h *DNSHeader) []byte {
	buf := &bytes.Buffer{}
	cache := make([]byte, 2)
	arr := []uint16{h.ID, uint16(h.Opt), h.QDCount, h.AnCount, h.NsCount, h.ArCount}
	for i := 0; i < len(arr); i++ {
		binary.BigEndian.PutUint16(cache, arr[i])
		buf.Write(cache)
	}
	return buf.Bytes()
}
func (qe *queryEncoder) encodeAnswers(ans []*Answer, idxs []int) ([]byte, error) {
	if idxs != nil && len(ans) != len(idxs) {
		return nil, errEncode
	}
	useIndex := false
	if idxs != nil {
		useIndex = true
	}
	buf := &bytes.Buffer{}
	var idx int
	for i := 0; i < len(ans); i++ {
		if useIndex {
			idx = idxs[i]
		}
		qe.encodeAnswer(ans[i], buf, idx)
	}
	return buf.Bytes(), nil
}
func (qe *queryEncoder) encodeAnswer(answer *Answer, writer io.Writer, index int) {
	cache := make([]byte, 4)
	if index < 1 {
		writer.Write(encodeDomain(answer.Name))
	} else {
		binary.BigEndian.PutUint16(cache, 0xc000|uint16(index))
		writer.Write(cache)
	}
	binary.BigEndian.PutUint16(cache, uint16(answer.Type))
	writer.Write(cache[:2])
	binary.BigEndian.PutUint16(cache, answer.Class)
	writer.Write(cache[:2])
	binary.BigEndian.PutUint32(cache, answer.TTL)
	writer.Write(cache)
	binary.BigEndian.PutUint16(cache, answer.RDLen)
	writer.Write(cache[:2])
	writer.Write(answer.RData)
}
