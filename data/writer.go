package data

import (
	"io"
	"os"
	"bytes"
	"encoding/binary"
)

type bstWriter struct {
	orw *ObjectRW
}

func (bw *bstWriter) Write(tree IndexTree, writer io.Writer) error {
	if tree.Empty() {
		return nil
	}
	bw.writeIndex(tree.Root(), writer)
	f := func(i *Record) {
		bw.writeIndex(i, writer)
	}
	tree.TraverseLeft(f)
	tree.TraverseRight(f)
	return nil
}
func (bw *bstWriter) writeIndex(index *Record, writer io.Writer) error {
	bits := bw.SerializeRecord(index)
	_, err := writer.Write(bits)
	return err
}
func (bw *bstWriter) SerializeRecord(record *Record) []byte {
	buf := &bytes.Buffer{}
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, record.Index())
	buf.Write(bs)
	nb := []byte(record.Name)
	binary.BigEndian.PutUint32(bs, uint32(len(nb)))
	buf.Write(bs)
	buf.Write(nb)
	binary.BigEndian.PutUint32(bs, record.Type)
	buf.Write(bs)
	binary.BigEndian.PutUint32(bs, record.Class)
	buf.Write(bs)
	binary.BigEndian.PutUint32(bs, record.Ac)
	buf.Write(bs)
	for i := 0; i < len(record.Raw); i++ {
		bw.orw.Write(record.Raw[i], buf)
	}
	return buf.Bytes()
}
func (bw *bstWriter) WriteFile(tree IndexTree, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return bw.Write(tree, f)
}
