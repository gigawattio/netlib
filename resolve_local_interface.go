package netlib

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
)

// ResolveLocalInterface takes an address string or regular expression and
// resolves/locates/validates it as a local IP address.  May also be a regular
// expression [e.g. '192.168.*'].
// If bind parameter is empty the first address NOT matching 127.* will be
// returned.
func ResolveLocalInterface(bind string) (net.IP, error) {
	var expr *regexp.Regexp

	if bind != "" {
		log.Info("Locating bind network interface with IP address matching %q", bind)
		var err error
		if expr, err = regexp.Compile(bind); err != nil {
			return nil, fmt.Errorf("compiling bind address expression: %s", err)
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
			if bind != "" && (bind == ipStr || expr.MatchString(ipStr)) {
				log.Info("Successfully verified bind IP address=%v", ipStr)
				return ip, nil
			}
			if strings.HasPrefix(ipStr, "127.") || strings.Contains(ipStr, "::") {
				continue
			}
			if expr == nil {
				log.Info("Auto-detected bind IP address=%v", ipStr)
				return ip, nil
			}
		}
	}
	return nil, errors.New("no usable ip address found")
}
