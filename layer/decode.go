package layer

import (
	"io"
	"errors"
	"strings"
)

var (
	errDecode = errors.New("failed to encode")
	errFormat = errors.New("error format")
)

type QueryDecoder struct {
}

func (qd *QueryDecoder) DecodeBytes(buf []byte) (*Query, error) {
	return nil, nil
}
func (qd *QueryDecoder) DecodeReader(reader io.Reader) (*Query, error) {
	return nil, nil
}

func (qd *QueryDecoder) DecodeHeader(bs []byte) (*DNSHeader, error) {
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

func (qd *QueryDecoder) DecodeQuestion(bs []byte, offset int) (*Question, int, error) {
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
	return &Question{QName: bh.String()}, l, nil
}
