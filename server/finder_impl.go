package server

import (
	"dnsgo/layer"
	"net"
	"dnsgo/data"
	"log"
)

type ele struct {
	Prev *ele
	Data *layer.Answer
}

func (e *ele) Join(an []*layer.Answer) *ele {
	ret := e
	for i := 0; i < len(an); i++ {
		t := &ele{Prev: ret, Data: an[i]}
		ret = t
	}
	return ret
}
func (s *server) Find(remoteAddr *net.UDPAddr, query *layer.Query) (*layer.Query, error) {
	query.Header.Opt = layer.NewOption(layer.QROpt)
	c := len(query.Questions)
	var holder *ele
	for i := 0; i < c; i++ {
		q := query.Questions[i]
		r := &data.Record{
			Name:  q.QName,
			Type:  uint32(q.Type),
			Class: uint32(q.Class),
		}
		rs, err := s.rTree.Search(r.Index())
		if err != nil {
			log.Println(err)
			break
		}
		query.Header.AnCount += uint16(len(rs.Raw))
		if holder == nil {
			holder = &ele{Prev: nil, Data: nil}
		}
		holder = holder.Join(rs.Raw)
	}
	ret := make([]*layer.Answer, query.Header.AnCount)
	for i := 0; i < int(query.Header.AnCount); i++ {
		ret[i] = holder.Data
		holder = holder.Prev
	}
	query.Header.ArCount = 0
	query.Answers = ret
	return query, nil
}
