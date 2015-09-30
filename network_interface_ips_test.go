package netlib

import (
	"testing"
)

func TestNetworkInterfaceIps(t *testing.T) {
	ips, err := NetworkInterfaceIps()
	if err != nil {
		t.Fatal(err)
	}
	if l := len(ips); l == 0 {
		t.Errorf("Expected NetworkInterfaceIps to return more than 0 network interface IP addresses but len(ips)=%v", l)
	}
}
