package arguments

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	tcpPortMIN = 1
	tcpPortMAX = 65535
)

var (
	helpArgumentRegexp = regexp.MustCompile(`^(-h|--help)$`)
	hostWithPortRegexp = regexp.MustCompile(`.*:\d+$`)
	hostPortPartRegexp = regexp.MustCompile(`:\d+$`)
)

func isSSHServerSet(s string) bool { return s != dummySSHServer }

func isSSHPasswordSet(s string) bool { return s != dummySSHPassword }

func isTCPPortValid(i int) bool { return !(i < tcpPortMIN) || (i > tcpPortMAX) }

func isHostWithPort(s string) bool { return hostWithPortRegexp.MatchString(s) }

func concatHostPort(host string, port int) string {
	return strings.Join([]string{host, strconv.Itoa(port)}, ":")
}

func mergeHostPort(host string, port int) string {
	bareHost := hostPortPartRegexp.ReplaceAllString(host, "")

	return concatHostPort(bareHost, port)
}
