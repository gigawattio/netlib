package netlib

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
)

var (
	NoAddressFoundError = errors.New("no usable ip address found")
)

// ResolveLocalInterface takes an address string or regular expression and
// resolves/locates/validates it as a local IP address.  May also be a regular
// expression (e.g. "192.168.*").
// If bind parameter is empty, "0.0.0.0" or "0:0:0:0:0:0:0:0" the first address
// NOT matching 127.* will be returned.
func ResolveLocalInterface(bind string, excludeIps ...string) (net.IP, error) {
	var (
		expr        *regexp.Regexp
		bindHasSpec bool = !IsEmptyBindSpec(bind)
	)

	if bindHasSpec {
		log.Info("Locating bind network interface with IP address matching %q", bind)
		var err error
		if expr, err = regexp.Compile(bind); err != nil {
			return nil, fmt.Errorf("compiling bind address expression %q: %s", bind, err)
		}
	} else {
		log.Info("Auto-detecting bind IP address..")
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("net.Interfaces: %s", err)
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, fmt.Errorf("getting ip addresses from interface=%+v: %s", iface, err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ipStr := ip.String()
			if bindHasSpec && (bind == ipStr || expr.MatchString(ipStr)) {
				log.Info("Successfully verified bind IP address=%v", ipStr)
				return ip, nil
			}
			if strings.HasPrefix(ipStr, "127.") || strings.Contains(ipStr, "::") || (len(excludeIps) > 0 && strings.Contains(strings.Join(excludeIps, " ")+" ", ipStr+" ")) {
				continue
			}
			if expr == nil {
				log.Info("Auto-detected bind IP address=%v", ipStr)
				return ip, nil
			}
		}
	}
	return nil, NoAddressFoundError
}

func IsEmptyBindSpec(bind string) bool {
	isEmtpy := bind == "" || bind == "[::]" || bind == "0.0.0.0" || bind == "0:0:0:0:0:0:0:0"
	return isEmtpy
}
