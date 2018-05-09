package data

import (
	"io"
	"os"
	"errors"
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
	for {
		n, err := br.readIndex(reader)
		switch err {
		case nil:
			bst.Insert(n)
		case errStop:
			break
		default:
			gerr = err
			break
		}
	}
	return bst, gerr
}
func (br *bstReader) readIndex(reader io.Reader) (*Index, error) {
	//TODO implements readIndex method
}
func (br *bstReader) ReadFile(filename string) (IndexTree, error) {
	input, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer input.Close()
	return br.Read(input)
}
