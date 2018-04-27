package layer

import "io"

type Packer interface {
	Encode(query *Query) []byte
	Decode([]byte) (*Query, error)
	DecodeReader(reader io.Reader) (*Query, error)
}

type packer struct {
}
