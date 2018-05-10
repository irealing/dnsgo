package data

import (
	"io"
	"dnsgo/layer"
)

type BSTReaderWriter interface {
	Write(tree IndexTree, writer io.Writer) error
	WriteFile(tree IndexTree, filename string) error
	Read(reader io.Reader) (IndexTree, error)
	ReadFile(filename string) (IndexTree, error)
}
type bstRWImpl struct {
	bstReader
	bstWriter
}

func NewBstRWImpl() BSTReaderWriter {
	return &bstRWImpl{
		bstReader{},
		bstWriter{layer.NewPacker()},
	}
}
