package layer

import (
	"io"
	"errors"
	"strings"
	"encoding/binary"
)

var (
	errDecode = errors.New("failed to encode")
	errFormat = errors.New("error format")
)

type queryDecoder struct {
}

func (qd *queryDecoder) Decode(bs []byte) (*Query, error) {
	if len(bs) < 12 {
		return nil, errFormat
	}
	header, err := qd.decodeHeader(bs[:12])
	if err != nil {
		return nil, err
	}
	q := new(Query)
	q.Header = header
	q.Questions, err = qd.decodeQuestions(bs[12:], int(header.QDCount))
	return q, err
}

func (qd *queryDecoder) DecodeReader(reader io.Reader) (*Query, error) {
	buf := make([]byte, 512)
	n, err := reader.Read(buf)
	if err != nil {
		return nil, err
	}
	return qd.Decode(buf[:n])
}

func (qd *queryDecoder) decodeHeader(bs []byte) (*DNSHeader, error) {
	if len(bs) != 12 {
		return nil, errDecode
	}
	id := decodeU16(bs[:2])
	opt := decodeU16(bs[2:4])
	qdCount := decodeU16(bs[4:6])
	anCount := decodeU16(bs[6:8])
	nsCount := decodeU16(bs[8:10])
	arCount := decodeU16(bs[10:])
	header := &DNSHeader{
		ID:      id, Opt: Option(opt), QDCount: qdCount, AnCount: anCount,
		NsCount: nsCount, ArCount: arCount,
	}
	return header, nil
}

func (qd *queryDecoder) decodeQuestions(bs []byte, num int) ([]*Question, error) {
	rs := make([]*Question, num)
	offset := 0
	for i := 0; i < num; i++ {
		q, l, err := qd.decodeQuestion(bs, offset)
		if err != nil {
			return nil, err
		}
		rs[i] = q
		offset += l
	}
	return rs, nil
}
func (qd *queryDecoder) decodeQuestion(bs []byte, offset int) (*Question, int, error) {
	c := offset
	l := 0
	bh := strings.Builder{}
	bl := len(bs)
	for {
		c = offset + l
		if c >= bl {
			return nil, l, errFormat
		}
		l += 1
		limit := int(bs[c])
		if limit == 0 {
			break
		}
		l += limit
		limit += c + 1
		if limit > bl || c+1 >= bl {
			return nil, l, errFormat
		}
		bh.Write(bs[c+1 : limit])
		bh.WriteRune('.')
	}
	qt := binary.BigEndian.Uint16(bs[l : l+2])
	l += 2
	qc := binary.BigEndian.Uint16(bs[l : l+2])
	l += 2
	return &Question{QName: bh.String(), Type: QType(qt), Class: qc}, l, nil
}
