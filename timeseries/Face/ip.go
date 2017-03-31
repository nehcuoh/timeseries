package Face

import "net"

type IP struct {
	net.IP
}

func ParseIP(ip string) *IP {
	return &IP{IP: net.ParseIP(ip)}
}

func (i *IP) Key() (interface{}) {
	return i.String()
	//i.IP.To4()
	//return
}
