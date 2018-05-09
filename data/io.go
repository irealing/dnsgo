package data

import "io"

type BSTReaderWriter interface {
	Write(tree *bsTree, writer io.Writer) error
	WriteFile(tree *bsTree, filename string) error
	Read(reader io.Reader) (*bsTree, error)
	ReadFile(filename string) (*bsTree, error)
}
type bstRWImpl struct {
	bstReader
	bstWriter
}
