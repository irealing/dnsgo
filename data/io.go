package data

import "io"

type BSTReaderWriter interface {
	Write(tree *bsTree, writer io.Writer) error
	WriteFile(tree *bsTree, filename string) error
	Reade(reader io.Reader) (*bsTree, error)
	ReadeFile(filename string) (*bsTree, error)
}
