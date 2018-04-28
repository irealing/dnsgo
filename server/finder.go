package server

import "dnsgo/layer"

type DNSFinder interface {
	Find(query *layer.Query) (*layer.Query, error)
}
