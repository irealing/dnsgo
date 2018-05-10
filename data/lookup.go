package data

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Index struct {
	Name  string
	Type  uint32
	Class uint32
	Start uint32
	End   uint32
}

func (i Index) String() string {
	return fmt.Sprintf("Index(Name=%s, Type=%d ,Class=%d)", i.Name, i.Type, i.Class)
}
func (i Index) Index() uint32 {
	bf := bytes.NewBuffer([]byte(i.Name))
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i.Type)
	bf.Write(bs)
	binary.BigEndian.PutUint32(bs, i.Class)
	bf.Write(bs)
	bits := bf.Bytes()
	return MurMurHash(bits, uint32(len(bits)))
}
func (i Index) Serialize() []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i.Index())
	buf := bytes.NewBuffer(bs)
	bn := []byte(i.Name)
	buf.WriteByte(byte(len(bn)))
	buf.Write(bn)
	binary.BigEndian.PutUint32(bs, i.Type)
	buf.Write(bs)
	binary.BigEndian.PutUint32(bs, i.Class)
	buf.Write(bs)
	binary.BigEndian.PutUint32(bs, i.Start)
	buf.Write(bs)
	binary.BigEndian.PutUint32(bs, i.End)
	buf.Write(bs)
	return buf.Bytes()
}
