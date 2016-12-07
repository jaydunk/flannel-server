package store

import (
	"errors"
	"fmt"
	"net"
)

type Store struct {
	data map[net.IPNet]net.IP
}

func (s *Store) GetLease(hostIP net.IP) (net.IPNet, error) {
	for subnet, owner := range s.data {
		if owner == hostIP {
			return subnet, nil
		}
	}

	for i := 0; i < 256; i++ {
		cidrString := fmt.Sprintf("10.255.%d.0/24", i)
		_, cidr, err := net.ParseCIDR(cidrString)
		if err != nil {
			panic("shouldn't ever happen")
		}

		_, ok := s.data[cidr]
		if !ok {
			s.data[cidr] = hostIP
			return cidr, nil
		}
	}
	return nil, errors.New("no lease available")
}
