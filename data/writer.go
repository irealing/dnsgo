package data

import (
	"io"
	"os"
	"dnsgo/layer"
	"bytes"
	"encoding/binary"
)

type bstWriter struct {
	coder layer.Packer
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
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, record.Index())
	nb := []byte(record.Name)
	binary.BigEndian.PutUint32(bs, uint32(len(nb)))
	buf := bytes.NewBuffer(bs)
	buf.Write(bs)
	buf.Write(nb)
	binary.BigEndian.PutUint32(bs, record.Type)
	buf.Write(bs)
	binary.BigEndian.PutUint32(bs, record.Class)
	buf.Write(bs)
	binary.BigEndian.PutUint32(bs, record.Ac)
	//raw, _ := bw.coder.EncodeAnswers(record.Raw, nil)
	//buf.Write(raw)
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
