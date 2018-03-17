package netlib

import (
	"fmt"
	"testing"
)

func TestAsIps(t *testing.T) {
	var (
		input = []string{
			"localhost",
			"localhost:5144",
			"localhost:21",
			"127.0.0.2",
			"127.0.0.3:47",
			"[::1]",
			"[::1]:443",
			"[FE80:0000:0000:0000:0202:B3FF:FE1E:8329]",
			"[FE80:0000:0000:0000:0202:B3FF:FE1E:8329]:445",
			"[FE80::0202:B3FF:FE1E:8329]",
			"[FE80::0202:B3FF:FE1E:8329]:80",
		}
		expected = []string{
			"127.0.0.1",
			"127.0.0.1:5144",
			"127.0.0.1:21",
			"127.0.0.2",
			"127.0.0.3:47",
			"[::1]",
			"[::1]:443",
			"[FE80:0000:0000:0000:0202:B3FF:FE1E:8329]",
			"[FE80:0000:0000:0000:0202:B3FF:FE1E:8329]:445",
			"[FE80::0202:B3FF:FE1E:8329]",
			"[FE80::0202:B3FF:FE1E:8329]:80",
		}
	)
	actual, err := AsIps(input...)
	if err != nil {
		t.Fatalf("asIps failed: %s", err)
	}
	if fmt.Sprintf("%+v", actual) != fmt.Sprintf("%+v", expected) {
		t.Fatalf("Expected output slice:\n\t%+v\nbut actual was:\n\t%+v", expected, actual)
	}
}
