package data

import (
	"io"
	"os"
	"errors"
	"encoding/binary"
	"dnsgo/layer"
)

var (
	errStop = errors.New("stop")
	errFmt  = errors.New("err format")
)

type bstReader struct {
	coder layer.Packer
}

func (br *bstReader) Read(reader io.Reader) (IndexTree, error) {
	bst := new(bsTree)
	var gerr error
outer:
	for {
		n, err := br.readRecord(reader)
		switch err {
		case nil:
			bst.Insert(n)
		case errStop:
			break outer
		default:
			gerr = err
			break outer
		}
	}
	return bst, gerr
}
func (br *bstReader) readRecord(reader io.Reader) (*Record, error) {
	bs := make([]byte, 12)
	n, err := reader.Read(bs[:8])
	if err != nil || n != 8 {
		if n > 0 {
			return nil, errFmt
		} else {
			return nil, errStop
		}
	}
	l := int(binary.BigEndian.Uint32(bs[4:8]))
	nb := make([]byte, l)
	if n, err := reader.Read(nb); err != nil || n != l {
		return nil, errFmt
	}
	var ac, t, c uint32
	if n, err := reader.Read(bs); n != 12 || err != nil {
		return nil, errFmt
	} else {
		ac = binary.BigEndian.Uint32(bs)
		t = binary.BigEndian.Uint32(bs[4:])
		c = binary.BigEndian.Uint32(bs[8:])
	}
	an, err := br.readAnswers(reader, int(ac))
	if err != nil {
		return nil, err
	}
	return &Record{Name: string(nb), Type: t, Class: c, Ac: ac, Raw: an}, nil
}
func (br *bstReader) readAnswers(reader io.Reader, n int) ([]*layer.Answer, error) {
	ret := make([]*layer.Answer, n)
	for i := 0; i < n; i++ {
		if r, err := br.readAnswer(reader); err != nil {
			return nil, err
		} else {
			ret[i] = r
		}
	}
	return ret, nil
}
func (br *bstReader) readAnswer(reader io.Reader) (*layer.Answer, error) {
	return nil, nil
}
func (br *bstReader) ReadFile(filename string) (IndexTree, error) {
	input, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer input.Close()
	return br.Read(input)
}
