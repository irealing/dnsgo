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
	log.Println("dns server shutdown")
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
		return err
	}
	s.conn = conn
	return s.listen()
}

func (s *server) listen() error {
	qc := layer.NewPacker()
	buf := make([]byte, 512)
	for {
		n, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			return err
		}
		q, err := qc.Decode(buf[:n])
		if err != nil {
			log.Printf("decode error %v", err)
			continue
		}
		go s.handleQuery(addr, q, qc)
	}
	return nil
}

func (s *server) handleQuery(addr *net.UDPAddr, query *layer.Query, p layer.Packer) {
	log.Printf("recv dns query %s", addr.String())
	var r *layer.Query
	var err error
	if finder, ok := interface{}(s).(DNSFinder); ok {
		r, err = finder.Find(addr, query)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		r = s.defaultFinder(query)
	}
	rs, err := p.Encode(r)
	if err != nil {
		log.Println(err)
		return
	}
	s.conn.WriteToUDP(rs, addr)
}

func (s *server) defaultFinder(query *layer.Query) *layer.Query {
	query.Header.Opt = layer.NewOption(layer.QROpt, layer.RCodeOPt(2))
	query.Header.ArCount = 0
	return query
}
func (s *server) Addr() *net.UDPAddr {
	return s.addr
}
