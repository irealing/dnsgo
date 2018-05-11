package data

import (
	"io"
	"reflect"
	"errors"
	"encoding/binary"
)

type BSTReaderWriter interface {
	Write(tree IndexTree, writer io.Writer) error
	WriteFile(tree IndexTree, filename string) error
	Read(reader io.Reader) (IndexTree, error)
	ReadFile(filename string) (IndexTree, error)
}
type bstRWImpl struct {
	bstReader
	bstWriter
}

func NewBstRWImpl() BSTReaderWriter {
	owr := new(ObjectRW)
	return &bstRWImpl{
		bstReader{owr},
		bstWriter{owr},
	}
}

type ObjectRW struct {
}

func (srw *ObjectRW) Read(o interface{}, reader io.Reader) error {
	ov := reflect.ValueOf(o)
	if ov.Kind() != reflect.Ptr || ov.Elem().Type().Kind() != reflect.Struct {
		return errors.New("error type,Ptr type except")
	}
	ov = ov.Elem()
	ot := ov.Type()
	var err error
	for i := 0; i < ot.NumField(); i++ {
		oft := ot.Field(i)
		ofv := ov.Field(i)
		switch oft.Type.Kind() {
		case reflect.Slice:
			var v int64
			v, err = srw.readInt(reader)
			if err != nil {
				break
			}
			sv := reflect.MakeSlice(oft.Type, int(v), int(v))
			err = srw.injectSlice(sv, reader)
			if err == nil {
				ofv.Set(sv)
			}
		default:
			err = srw.injectValue(ofv, reader)
		}
		if err != nil {
			break
		}
	}
	return err
}
func (srw *ObjectRW) injectSlice(v reflect.Value, reader io.Reader) error {
	var err error
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		err = srw.injectValue(e, reader)
		if err != nil {
			break
		}
	}
	return err
}
func (srw *ObjectRW) injectValue(ofv reflect.Value, reader io.Reader) error {
	var err error
	switch ofv.Type().Kind() {
	case reflect.String:
		var s string
		s, err = srw.readString(reader)
		ofv.SetString(s)
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		var uv uint64
		uv, err = srw.readUint(reader)
		ofv.SetUint(uv)
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		var v int64
		v, err = srw.readInt(reader)
		ofv.SetInt(v)
	}
	return err
}
func (srw *ObjectRW) readString(reader io.Reader) (string, error) {
	buf := make([]byte, 4)
	n, err := reader.Read(buf)
	if err != nil || n != 4 {
		return "", errors.New("error format")
	}
	l := binary.BigEndian.Uint32(buf)
	bits := make([]byte, int(l))
	if n, err = reader.Read(bits); n != int(l) {
		return "", errors.New("error format")
	}
	return string(bits), nil
}
func (srw *ObjectRW) readBytes(reader io.Reader, n int) ([]byte, error) {
	bits := make([]byte, n)
	if s, err := reader.Read(bits); err != nil || s != n {
		return nil, errors.New("error format")
	}
	return bits, nil
}
func (srw *ObjectRW) readUint(reader io.Reader) (uint64, error) {
	bits, err := srw.readBytes(reader, 8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(bits), err
}
func (srw *ObjectRW) readInt(reader io.Reader) (int64, error) {
	bits, err := srw.readBytes(reader, 8)
	if err != nil {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(bits)), err
}
func (srw *ObjectRW) Write(o interface{}, writer io.Writer) error {
	ov := reflect.ValueOf(o)
	if ov.Kind() == reflect.Ptr {
		ov = ov.Elem()
	}
	ot := ov.Type()
	var err error
	for i := 0; i < ot.NumField(); i++ {
		sfv := ov.Field(i)
		sft := ot.Field(i)
		switch sft.Type.Kind() {
		case reflect.Slice:
			err = srw.writeSlice(sfv, writer)
		default:
			err = srw.writeValue(sfv, writer)
		}
		if err != nil {
			break
		}
	}
	return err
}

func (srw *ObjectRW) writeValue(sfv reflect.Value, writer io.Writer) error {
	var err error
	switch sfv.Type().Kind() {
	case reflect.String:
		s := sfv.String()
		err = srw.writeString(s, writer)
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		ui := sfv.Uint()
		err = srw.writeUint(ui, writer)
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		i := sfv.Int()
		err = srw.writeInt(i, writer)
	}
	return err
}
func (srw *ObjectRW) writeSlice(v reflect.Value, writer io.Writer) error {
	l := v.Len()
	srw.writeInt(int64(l), writer)
	var err error
	for i := 0; i < l; i++ {
		err = srw.writeValue(v.Index(i), writer)
		if err != nil {
			break
		}
	}
	return err
}
func (*ObjectRW) writeString(v string, writer io.Writer) error {
	tb := make([]byte, 4)
	sb := []byte(v)
	binary.BigEndian.PutUint32(tb, uint32(len(sb)))
	if _, err := writer.Write(tb); err != nil {
		return err
	}
	if _, err := writer.Write(sb); err != nil {
		return err
	}
	return nil
}
func (*ObjectRW) writeUint(v uint64, writer io.Writer) error {
	ub := make([]byte, 8)
	binary.BigEndian.PutUint64(ub, v)
	_, err := writer.Write(ub)
	return err
}
func (*ObjectRW) writeInt(v int64, writer io.Writer) error {
	ub := make([]byte, 8)
	binary.BigEndian.PutUint64(ub, uint64(v))
	_, err := writer.Write(ub)
	return err
}
