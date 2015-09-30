package netlib

import (
	"fmt"
	"net"
)

// NetworkInterfaceIps returns all local network interface IP addresses.
func NetworkInterfaceIps() ([]net.IP, error) {
	addrToIp := func(addr net.Addr) (ip net.IP) {
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		return
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("getting net interfaces: %s", err)
	}

	ips := []net.IP{}
	for _, adapter := range interfaces {
		addrs, err := adapter.Addrs()
		if err != nil {
			return nil, fmt.Errorf("getting addresses for adapter=%+v: %s", adapter, err)
		}
		for _, addr := range addrs {
			ips = append(ips, addrToIp(addr))
		}
	}
	return ips, nil
}
