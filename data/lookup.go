package data

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"dnsgo/layer"
)

type Record struct {
	Name  string
	Type  uint32
	Class uint32
	Ac    uint32
	Raw   []*layer.Answer
}

func (i Record) String() string {
	return fmt.Sprintf("Record(Name=%s, Type=%d ,Class=%d)", i.Name, i.Type, i.Class)
}
func (i Record) Index() uint32 {
	bf := bytes.NewBuffer([]byte(i.Name))
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i.Type)
	bf.Write(bs)
	binary.BigEndian.PutUint32(bs, i.Class)
	bf.Write(bs)
	bits := bf.Bytes()
	return MurMurHash(bits, uint32(len(bits)))
}
