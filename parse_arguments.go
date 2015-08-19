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
func ParseCommandLineArguments() (string, string, int, string, error) {
	var sshServer = flag.String("s", DefaultSSHServer, "SSH server (host[:port]) to connect")
	var targetHost = flag.String("t", DefaultTargetHost, "target host:port we will forward connections to")
	var exposedPort = flag.Int("e", DefaultExposedPort, "port to expose and forward on the SSH server side")
	var sshUsername = flag.String("u", DefaultSSHUsername, "username to connect SSH server")

	flag.Parse()

	if *sshServer == DummySSHServer {
		return "", "", 0, "", errors.New("SSH server not defined")
	}

	if (*exposedPort < TCPPortMIN) || (*exposedPort > TCPPortMAX) {
		return "", "", 0, "", errors.New("forwarded port number is not in the valid range")
	}

	if !HostWithPortRegexp.MatchString(*sshServer) {
		*sshServer = strings.Join([]string{*sshServer, strconv.Itoa(DefaultSSHPort)}, ":")
	}

	if !HostWithPortRegexp.MatchString(*targetHost) {
		*targetHost = strings.Join([]string{*targetHost, strconv.Itoa(*exposedPort)}, ":")
	}

	return *sshServer, *targetHost, *exposedPort, *sshUsername, nil
}
