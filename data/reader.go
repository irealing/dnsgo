package data

import (
	"io"
	"os"
	"errors"
	"encoding/binary"
)

var (
	errStop = errors.New("stop")
	errFmt  = errors.New("err format")
)

type bstReader struct {
}

func (br *bstReader) Read(reader io.Reader) (IndexTree, error) {
	bst := new(bsTree)
	var gerr error
outer:
	for {
		_, n, err := br.readIndex(reader)
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
func (br *bstReader) readIndex(reader io.Reader) (uint32, *Index, error) {
	bs := make([]byte, 16)
	n, err := reader.Read(bs[:5])
	if err != nil || n < 5 {
		if n > 0 {
			return 0, nil, errFmt
		}
		return 0, nil, errStop
	}
	v := binary.BigEndian.Uint32(bs)
	l := int(bs[4])
	nbs := make([]byte, l)
	n, err = reader.Read(nbs)
	if err != nil || n != l {
		return 0, nil, errFmt
	}
	n, err = reader.Read(bs)
	if err != nil || n != 16 {
		return 0, nil, errFmt
	}
	t := binary.BigEndian.Uint32(bs)
	c := binary.BigEndian.Uint32(bs[4:])
	s := binary.BigEndian.Uint32(bs[8:])
	e := binary.BigEndian.Uint32(bs[12:])
	return v, &Index{Name: string(nbs), Type: t, Class: c, Start: s, End: e}, nil
}
func (br *bstReader) ReadFile(filename string) (IndexTree, error) {
	input, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer input.Close()
	return br.Read(input)
}
