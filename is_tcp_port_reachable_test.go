package netlib_test

import (
	"testing"

	"gigawatt-common/pkg/netlib"
	"gigawatt-common/pkg/web"
)

func TestIsTcpPortReachable(t *testing.T) {
	server := web.NewWebServer(web.WebServerOptions{
		Addr: "127.0.0.1:0",
	})
	if err := server.Start(); err != nil {
		t.Fatalf("Error starting WebServer: %s", err)
	}
	addr := server.Addr().String()
	if expected, actual := true, netlib.IsTcpPortReachable(addr); actual != expected {
		t.Errorf("Expected address %v to be reachable but actual result=%v", addr, actual)
	}
	if err := server.Stop(); err != nil {
		t.Fatalf("Error stopping WebServer: %s", err)
	}
	if expected, actual := false, netlib.IsTcpPortReachable(addr); actual != expected {
		t.Errorf("Expected address %v to be unreachable after server is shutdown but actual=%v", addr, actual)
	}
}
