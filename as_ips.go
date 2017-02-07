package netlib

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
)

// AsIps resolves hostname:port pairs to IPv4-address:port pairs.
func AsIps(hostnamePortPairs ...string) ([]string, error) {
	var (
		hasPortExpr  = regexp.MustCompile("[^:]:[0-9]+$")
		n            = len(hostnamePortPairs)
		hostnames    = make([]string, n)
		portMappings = make([]string, n)
		final        = []string{}
		err          error
	)
	for i := 0; i < n; i++ {
		if hasPortExpr.MatchString(hostnamePortPairs[i]) {
			if hostnames[i], portMappings[i], err = net.SplitHostPort(hostnamePortPairs[i]); err != nil {
				return nil, fmt.Errorf("AsIps: splitting hostname/port from %q: %s", hostnamePortPairs[i], err)
			}
		} else {
			hostnames[i] = hostnamePortPairs[i]
		}
		hostnames[i] = strings.Trim(hostnames[i], "[]")
		if ip := net.ParseIP(hostnames[i]); ip != nil {
			portMappings[i] = hostnamePortPairs[i]
			hostnames[i] = ""
		}
	}
	var (
		resolutions   = BulkResolver4(hostnames...)
		maybeColon    string
		reconstructed string
	)
	for i, r := range resolutions.Results {
		if r.DomainName == "" {
			// If DomainName is empty the entry is an ip, backfill the value from
			// dual-purposed portMappings.
			final = append(final, portMappings[i])
		} else {
			if r.Error == nil && len(r.Ips) > 0 {
				if len(portMappings[i]) > 0 {
					maybeColon = ":"
				} else {
					maybeColon = ""
				}
				reconstructed = fmt.Sprintf("%v%v%v", r.Ips[0], maybeColon, portMappings[i])
				final = append(final, reconstructed)
			}
		}
	}
	if n != 0 && len(final) == 0 {
		return nil, errors.New("AsIps: 0 ips were found in the bulk-dns-resolution result")
	}
	if err := resolutions.AnyErrors(); err != nil {
		return nil, err
	}
	return final, nil
}
