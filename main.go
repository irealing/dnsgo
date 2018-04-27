package main

import (
	"dnsgo/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s, err := server.NewServer(cfg.Addr)
	if err != nil {
		log.Fatal(err)
	}
	go s.Serve()
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	v := <-stop
	log.Println("exit ", v)
	s.Shutdown()
}
