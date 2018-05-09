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
	bits := i.GetBytes()
	return MurMurHash(bits, uint32(len(bits)))
}
func (i Index) GetBytes() []byte {
	buf := bytes.NewBuffer([]byte(i.Name))
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i.Type)
	buf.Write(bs)
	binary.BigEndian.PutUint32(bs, i.Type)
	buf.Write(bs)
	binary.BigEndian.PutUint32(bs, i.Start)
	buf.Write(bs)
	binary.BigEndian.PutUint32(bs, i.End)
	buf.Write(bs)
	return buf.Bytes()
}
