package server

import (
	"net"
	"sync"
	"sync/atomic"
	"errors"
	"dnsgo/layer"
	"log"
)

var (
	ErrClosed = errors.New("closed dns server")
)

type DNSServer interface {
	Serve() error
	Addr() *net.UDPAddr
	Shutdown()
}

type server struct {
	addr      *net.UDPAddr
	conn      *net.UDPConn
	closeOnce sync.Once
	closed    int32
	mutex     sync.Mutex
}

func NewServer(addr string) (DNSServer, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	var server DNSServer = &server{
		addr:      udpAddr,
		closeOnce: sync.Once{},
		mutex:     sync.Mutex{},
	}
	return server, nil
}
func (s *server) Shutdown() {
	s.closeOnce.Do(s.shutdown)
}
func (s *server) shutdown() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	atomic.CompareAndSwapInt32(&s.closed, 0, 1)
	if s.conn != nil {
		s.conn.Close()
	}
}

func (s *server) Serve() error {
	if atomic.LoadInt32(&s.closed) != 0 {
		return ErrClosed
	}
	conn, err := net.ListenUDP("udp", s.addr)
	if err != nil {
	}
	s.conn = conn
	return s.listen()
}

func (s *server) listen() error {
	qc := &layer.QueryCoding{}
	buf := make([]byte, 512)
	for {
		n, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			return err
		}
		q, err := qc.DecodeBytes(buf[:n])
		if err != nil {
			continue
		}
		go s.handleQuery(addr, q, qc)
	}
	return nil
}

func (s *server) handleQuery(addr *net.UDPAddr, query *layer.Query, qc *layer.QueryCoding) {
	log.Printf("recv dns query %s", addr.String())
	query.Header.Opt = layer.NewOption(layer.QROpt)
	s.conn.WriteToUDP(qc.EncodeQuery(query), addr)
}
func (s *server) Addr() *net.UDPAddr {
	return s.addr
}
