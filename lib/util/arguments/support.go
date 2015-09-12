package arguments

import (
	"net"
	"strconv"
)

func isSSHServerSet(s string) bool { return s != dummySSHServer }

func isSSHPasswordSet(s string) bool { return s != dummySSHPassword }

func isTCPPortValid(port int) bool { return !(port < 1) || (port > 65535) }

func isHostWithPort(hostPort string) bool {
	_, _, err := net.SplitHostPort(hostPort)

	return err == nil
}

func joinHostPort(host string, port int) string {
	return net.JoinHostPort(host, strconv.Itoa(port))
}

func mergeHostPort(hostPort string, port int) string {
	host, _, _ := net.SplitHostPort(hostPort)

	return joinHostPort(host, port)
}
