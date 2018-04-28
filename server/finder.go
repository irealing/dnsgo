package server

import (
	"dnsgo/layer"
	"net"
)

type DNSFinder interface {
	Find(remotAddr *net.UDPAddr, query *layer.Query) (*layer.Query, error)
}
