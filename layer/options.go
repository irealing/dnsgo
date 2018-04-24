package layer

func QROpt(opt Option) Option {
	return (1 << 15) | opt
}
func AAOpt(option Option) Option {
	return (1 << 10) | option
}
