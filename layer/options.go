package layer

func QROpt(option Option) Option {
	opt := (1 << 15) | option
	return opt
}
func AAOpt(option Option) Option {
	return (1 << 10) | option
}
func TCOpt(option Option) Option {
	return option | (1 << 9)
}
func RDOpt(option Option) Option {
	return option | (1 << 8)
}
func RAOpt(option Option) Option {
	return option | (1 << 7)
}
func RCodeOPt(v uint8) OptCfg {
	flag := 0xf & v
	return func(option Option) Option {
		return option | Option(flag)
	}
}
func ZOpt(v uint8) OptCfg {
	flag := Option((0x7 & v) << 4)
	return func(option Option) Option {
		return option | flag
	}
}
func OPCodeOpt(v uint8) OptCfg {
	flag := Option((0xf & uint16(v)) << 11)
	return func(option Option) Option {
		return option | flag
	}
}
