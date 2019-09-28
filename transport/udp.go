package transport

import "net"

type udp struct{}

// NewUDP returns a new udp instance
func NewUDP() udp {
	return udp{}
}

func (u udp) Listen(addr *net.UDPAddr) (*net.UDPConn, error) {
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func (u udp) IP(address net.Addr) net.IP {
	return address.(*net.UDPAddr).IP
}

func (u udp) Port(address net.Addr) uint16 {
	return uint16(address.(*net.TCPAddr).Port)
}

func (u udp) String() string {
	return "udp"
}
