package main

import (
	"fmt"
	"log"
	"github.com/irealing/argsparser"
	"errors"
)

var (
	cfg = new(Config)
)

type Config struct {
	Host string `param:"host"`
	Port int    `param:"port"`
	Addr string `param:""`
	Src  string `param:"src"`
}

func (c *Config) Validate() error {
	if c.Host == "" || c.Port < 1 {
		return errors.New("params error")
	}
	c.Addr = fmt.Sprintf("%s:%d", c.Host, c.Port)
	return nil
}
func init() {
	ap := argsparser.New(cfg)
	ap.Init()
	if err := ap.Parse(); err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("init config listen: %s", cfg.Addr)
}
