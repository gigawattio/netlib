package netlib

import (
	"os/exec"
	"strings"
)

func IsTcpPortReachable(addr string) bool {
	cmd := exec.Command("nc", append([]string{"-w", "1"}, strings.Split(addr, ":")...)...)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
