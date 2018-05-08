package data

import (
	"io"
	"os"
	"encoding/binary"
)

type bstWriter struct {
}

func (bw *bstWriter) Write(tree *bsTree, writer io.Writer) error {
	bw.writeNode(tree.root, writer)
	f := func(i *node) {
		bw.writeNode(i, writer)
	}
	traverse(tree.Root().Left, f)
	traverse(tree.Root().Right, f)
	return nil
}
func (bw *bstWriter) writeNode(n *node, writer io.Writer) error {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, n.Value)
	_, err := writer.Write(buf)
	if err != nil {
		return err
	}
	data := []byte(n.Attch)
	binary.BigEndian.PutUint32(buf, uint32(len(data)))
	if _, err := writer.Write(buf); err != nil {
		return err
	}
	if _, err := writer.Write(data); err != nil {
		return err
	}
	return nil
}
func (bw *bstWriter) WriteFile(tree *bsTree, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return bw.Write(tree, f)
}
