package layer

/*https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml*/
const (
	Unknown           QType = iota // unknown
	Adress                         // A
	NanmeServer                    //NS
	MailDestination                //MD
	MailForwarder                  //MF
	CName                          //CNAME
	SOA                            //SOA
	MailboxDomain                  //MB
	MailRenameDomain               //MR
	NullRR                         //NULL
	WellKnownService               //WKS
	DomainNamePointer              //PTR
	MailExchange                   //MX
	TextStrings                    //TEXT
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
	return opt&(1<<10) > 0
}
func (opt Option) TC() bool {
	return opt&(1<<9) > 0
}
func (opt Option) RD() bool {
	return opt&(1<<8) > 0
}
func (opt Option) RA() bool {
	return opt&1<<7 > 0
}
func (opt Option) RCode() uint8 {
	v := opt & 0xf
	return uint8(v)
}
func (opt Option) Z() uint8 {
	v := opt & (0x7 << 4)
	v = v >> 4
	return uint8(v)
}
func (opt Option) OPCode() uint8 {
	v := opt & (0xf << 11)
	v = v >> 11
	return uint8(v)
}
