package data

import (
	"io"
	"os"
)

type bstWriter struct {
}

func (bw *bstWriter) Write(tree IndexTree, writer io.Writer) error {
	if tree.Empty() {
		return nil
	}
	bw.writeIndex(tree.Root(), writer)
	f := func(i *Index) {
		bw.writeIndex(i, writer)
	}
	tree.TraverseLeft(f)
	tree.TraverseRight(f)
	return nil
}
func (bw *bstWriter) writeIndex(index *Index, writer io.Writer) error {
	//TODO implements writeIndex method
}
func (bw *bstWriter) WriteFile(tree IndexTree, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return bw.Write(tree, f)
}
