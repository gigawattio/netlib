package netlib

import (
	"net"
	"sync"

	"gigawatt-common/pkg/errorlib"
)

type (
	DnsResolution struct {
		DomainName string
		Ips        []string
		Error      error
	}
	DnsResolutions struct {
		Results []DnsResolution
	}
)

// BulkResolver4 resolves a list of domain names into their ipv4 addresses.
func BulkResolver4(names ...string) DnsResolutions {
	var wg sync.WaitGroup
	wg.Add(len(names))
	r := DnsResolutions{Results: make([]DnsResolution, len(names))}
	for i, name := range names {
		go func(i int, name string) {
			ip4address, err := net.ResolveIPAddr("ip4", name)
			if err == nil {
				r.Results[i] = DnsResolution{name, []string{ip4address.String()}, nil}
			} else {
				r.Results[i] = DnsResolution{name, nil, err}
			}
			wg.Done()
		}(i, name)
	}
	wg.Wait()
	return r
}

func (rs *DnsResolutions) AnyErrors() (err error) {
	errors := []error{}
	for _, r := range rs.Results {
		if r.Error != nil {
			errors = append(errors, r.Error)
		}
	}
	if err = errorlib.Merge(errors); err != nil {
		return
	}
	return
}
