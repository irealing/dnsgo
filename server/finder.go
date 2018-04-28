package server

import (
	"dnsgo/layer"
	"net"
)

type DNSFinder interface {
	Find(remoteAddr *net.UDPAddr, query *layer.Query) (*layer.Query, error)
}
