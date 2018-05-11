package main

import (
	"dnsgo/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s, err := server.NewServer(cfg.Addr, cfg.Src)
	if err != nil {
		log.Fatal(err)
	}
	go s.Serve()
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	v := <-stop
	s.Shutdown()
	log.Println("exit ", v)
}
