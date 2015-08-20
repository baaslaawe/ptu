package main

import (
	"errors"
	"flag"
	"strconv"
	"strings"
)

// Is help requested by passing -h|--help as an argument?
func IsHelpRequested(argument string) bool {
	return HelpArgumentRegexp.MatchString(argument)
}

// Parse command line arguments, perform some initial validation and variable mutation
func ParseCommandLineArguments() (*Config, error) {
	var sshServer = flag.String("s", DefaultSSHServer, "SSH server (host[:port]) to connect")
	var sshUsername = flag.String("u", DefaultSSHUsername, "username to connect SSH server")
	var sshPassword = flag.String("p", DefaultSSHPassword, "password to authenticate against SSH server")
	var targetHost = flag.String("t", DefaultTargetHost, "target host:port we will forward connections to")
	var exposedHost = flag.String("l", DefaultExposedHost, "host/bind (listener) to expose on the SSH server side")
	var exposedPort = flag.Int("e", DefaultExposedPort, "port to expose and forward on the SSH server side")

	flag.Parse()

	if !isSSHServerSet(*sshServer) {
		return nil, errors.New("SSH server not defined")
	}

	if !isTCPPortValid(*exposedPort) {
		return nil, errors.New("exposed TCP port number is invalid")
	}

	if !isHostWithPort(*sshServer) {
		*sshServer = concatHostPort(*sshServer, DefaultSSHPort)
	}

	if !isHostWithPort(*targetHost) {
		*targetHost = concatHostPort(*targetHost, DefaultTargetPort)
	}

	config := &Config{
		sshServer:   *sshServer,
		sshUsername: *sshUsername,
		sshPassword: *sshPassword,
		sshUseAgent: !isSSHPasswordSet(*sshPassword),
		targetHost:  *targetHost,
		exposedHost: *exposedHost,
		exposedPort: *exposedPort,
		exposedBind: concatHostPort(*exposedHost, *exposedPort),
		connectTo:   mergeHostPort(*sshServer, *exposedPort),
	}

	return config, nil
}

func isSSHServerSet(s string) bool { return s != DummySSHServer }

func isSSHPasswordSet(s string) bool { return s != DummySSHPassword }

func isTCPPortValid(i int) bool { return !(i < TCPPortMIN) || (i > TCPPortMAX) }

func isHostWithPort(s string) bool { return HostWithPortRegexp.MatchString(s) }

func concatHostPort(host string, port int) string {
	return strings.Join([]string{host, strconv.Itoa(port)}, ":")
}

func mergeHostPort(host string, port int) string {
	bareHost := HostPortPartRegexp.ReplaceAllString(host, "")

	return concatHostPort(bareHost, port)
}
