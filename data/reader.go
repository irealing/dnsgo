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

func (br *bstReader) Read(reader io.Reader) (*bsTree, error) {
	bst := new(bsTree)
	var gerr error
	for {
		n, err := br.readNode(reader)
		switch err {
		case nil:
			bst.Insert(n.Value, n.Attch)
		case errStop:
			break
		default:
			gerr = err
			break
		}
	}
	return bst, gerr
}
func (br *bstReader) readNode(reader io.Reader) (*node, error) {
	buf := make([]byte, 8)
	n, err := reader.Read(buf)
	if err != nil || n != 8 {
		return nil, errFmt
	}
	v := binary.BigEndian.Uint32(buf[:4])
	l := binary.BigEndian.Uint32(buf[4:])
	data := make([]byte, l)
	n, err = reader.Read(data)
	if err != nil || n != int(l) {
		return nil, errFmt
	}
	return &node{Value: v, Attch: string(data)}, nil
}
func (br *bstReader) ReadFile(filename string) (*bsTree, error) {
	input, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer input.Close()
	return br.Read(input)
}
