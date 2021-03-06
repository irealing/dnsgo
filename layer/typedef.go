package layer

import "fmt"

/**
  *https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml
*/
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
type OptCfg func(option Option) Option

func (opt Option) String() string {
	f := "QR: %v, OPCode: %v, AA: %v, TC: %v, RD: %v, RA: %v, Z: %v, RCode: %v"
	return fmt.Sprintf(f, opt.QR(), opt.OPCode(), opt.AA(), opt.TC(), opt.RD(), opt.RA(), opt.Z(), opt.RCode())
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

func NewOption(opts ... OptCfg) Option {
	var opt Option
	for _, cfg := range opts {
		opt = cfg(opt)
	}
	return opt
}
