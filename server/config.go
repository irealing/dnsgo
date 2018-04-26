package main

var (
	cfg *Config = nil
)

type Config struct {
	Host string
	Port int
}

func init() {
	cfg = new(Config)
	cfg.Host = ""
	cfg.Port = 65533
}
