package layer

import "io"

type QueryDecoder struct {
}

func (qd *QueryDecoder) DecodeBytes(buf []byte) (*Query, error) {
	return nil, nil
}
func (qd *QueryDecoder) DecodeReader(reader io.Reader) (*Query, error) {
	return nil, nil
}
