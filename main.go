package main

import (
	"net"
	"log"
	"dnsgo/layer"
	"os/signal"
	"os"
	"syscall"
)

func main() {
	go serve()
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	s := <-stop
	log.Println("exit", s)
}
func listen() (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", cfg.Addr)
	if err != nil {
		return nil, err
	}
	return net.ListenUDP("udp", addr)
}
func serve() {
	conn, err := listen()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listen", conn.RemoteAddr())
	buf := make([]byte, 512)
	qc := &layer.QueryCoding{}
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		q, err := qc.DecodeBytes(buf[:n])
		if err != nil {
			log.Println(addr, err)
			continue
		}
		go handleQuery(addr, conn, q, qc)
	}
	defer conn.Close()
}
func handleQuery(addr *net.UDPAddr, writer *net.UDPConn, query *layer.Query, coder *layer.QueryCoding) {
	log.Printf("recv dns query %s", addr.String())
	log.Println(query)
	query.Header.Opt = layer.NewOption(layer.QROpt)
	writer.WriteToUDP(coder.EncodeQuery(query), addr)
}
