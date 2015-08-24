package netlib

import (
	"errors"
	"strings"
)

// AsIps resolves hostname:port pairs to IPv4-address:port pairs.
func AsIps(hostnamePortPairs ...string) ([]string, error) {
	n := len(hostnamePortPairs)
	hostnames := make([]string, n)
	portMappings := make([]string, n)
	for i := 0; i < n; i++ {
		pieces := strings.Split(hostnamePortPairs[i], ":")
		if len(pieces) == 2 {
			portMappings[i] = ":" + pieces[1]
		}
		hostnames[i] = pieces[0]
	}
	resolutions := BulkResolver4(hostnames...)
	ips := []string{}
	for i, r := range resolutions.Results {
		if r.Error == nil && len(r.Ips) > 0 {
			reconstructed := r.Ips[0] + portMappings[i]
			ips = append(ips, reconstructed)
		}
	}
	if len(ips) == 0 && n != 0 {
		return nil, errors.New("AsIps: 0 ips were found in the bulk-dns-resolution result")
	}
	if err := resolutions.AnyErrors(); err != nil {
		return nil, err
	}
	return ips, nil
}
