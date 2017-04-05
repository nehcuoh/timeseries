package face

import (
	"fmt"
	"net"

	series "zeus/timeseries"
)

type IP struct {
	net.IP
}

func ParseIP(ip string) *IP {
	return &IP{IP: net.ParseIP(ip)}
}

func (*IP) ParseKey(key interface{}) (ip_obj series.FaceKey, err error) {
	k, e := key.(uint32)

	if e {
		info := fmt.Sprintf("Parse IP Fail,key: %d", key)
		return nil, &Error{info}
	}
	ip := uint32(k)
	bytesIP := &IP{make([]byte, 4)}
	bytesIP.IP[0] = byte(ip >> (3 * 8) & 0xff)
	bytesIP.IP[1] = byte(ip >> (2 * 8) & 0xff)
	bytesIP.IP[2] = byte(ip >> (1 * 8) & 0xff)
	bytesIP.IP[3] = byte(ip & 0xff)
	return bytesIP, nil
}

func (i *IP) Key() (interface{}) {
	return i.ToInt()
}

func (i*IP) String() string {
	return i.IP.String()
}

func (i*IP) ToInt() uint32 {
	ip := uint32(0)
	for _, b := range i.IP {
		ip = (ip << 8) | uint32(b)
	}
	return ip
}
