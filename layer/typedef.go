package layer

const (
	AType     QType = 1
	NSType    QType = 2
	CNameType QType = 5
)

type QType uint16
type Option uint16

func (qt QType) Encode() []byte {
	return encodeU16(uint16(qt))
}
func (opt Option) Encode() []byte {
	return encodeU16(uint16(opt))
}
func (opt Option) QR() bool {
	return opt>>15 > 0
}
func (opt Option) AA() bool {
	return opt&1<<10 > 0
}
func (opt Option) TC() bool {
	return opt&1<<9 > 0
}
func (opt Option) RD() bool {
	return opt&1<<8 > 0
}
func (opt Option) RA() bool {
	return opt&1<<7 > 0
}
func (opt Option) RCode() uint8 {
	v := opt << 12 >> 12
	return uint8(v)
}
func (opt Option) Z() uint8 {
	v := opt << 6 >> 13
	return uint8(v)
}
func (opt Option) OPCode() uint8 {
	v := opt << 1
	v = opt >> 13
	return uint8(v)
}
