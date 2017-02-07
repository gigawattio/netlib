package netlib

import (
	"net"
	"sync"

	"github.com/gigawattio/errorlib"
)

type DnsResolution struct {
	DomainName string
	Ips        []string
	Error      error
}

type DnsResolutions struct {
	Results []DnsResolution
}

// BulkResolver4 resolves a list of domain names into their ipv4 addresses.
func BulkResolver4(names ...string) DnsResolutions {
	var (
		r = DnsResolutions{
			Results: make([]DnsResolution, len(names)),
		}
		wg sync.WaitGroup
	)
	for i, name := range names {
		if name != "" {
			wg.Add(1)
			go func(i int, name string) {
				if ip4address, err := net.ResolveIPAddr("ip4", name); err != nil {
					r.Results[i] = DnsResolution{
						DomainName: name,
						Ips:        nil,
						Error:      err,
					}
				} else {
					r.Results[i] = DnsResolution{
						DomainName: name,
						Ips:        []string{ip4address.String()},
						Error:      nil,
					}
				}
				wg.Done()
			}(i, name)
		}
	}
	wg.Wait()
	return r
}

// Return a single combined squashed together error if there were any errors in
// any of the resolution tasks.
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
