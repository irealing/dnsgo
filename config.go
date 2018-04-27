package main

import (
	"fmt"
	"log"
)

var (
	cfg = new(Config)
)

type Config struct {
	Host string `param:"host"`
	Port int    `param:"port"`
	Addr string `param:"-"`
}

func (c *Config) Validate() error {
	c.Addr = fmt.Sprintf("%s:%d", c.Host, c.Port)
	return nil
}
func init() {
	cfg.Host = ""
	cfg.Port = 65533
	cfg.Validate()
	log.Printf("init config listen: %s", cfg.Addr)
}
