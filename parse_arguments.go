package main

import (
	"errors"
	"flag"
	"os/user"
	"strconv"
	"strings"
)

// Is help requested by passing -h|--help as an argument?
func IsHelpRequested(argument string) bool {
	return HelpArgumentRegexp.MatchString(argument)
}

// Parse command line arguments, perform some initial validation and variable mutation
func ParseCommandLineArguments() (string, string, int, string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", "", 0, "", err
	}

	var sshServer = flag.String("s", "", "SSH server (host[:port]) to connect")
	var targetHost = flag.String("t", "127.0.0.1:22", "target host:port we will forward connections to")
	var exposePort = flag.Int("e", 8080, "port to expose and forward on the SSH server side")
	var sshUsername = flag.String("u", currentUser.Username, "username to connect SSH server")

	flag.Parse()

	if (*exposePort < TCPPortMIN) || (*exposePort > TCPPortMAX) {
		return "", "", 0, "", errors.New("forwarded port number is not in the valid range")
	}

	if !HostWithPortRegexp.MatchString(*sshServer) {
		*sshServer = strings.Join([]string{*sshServer, strconv.Itoa(DefaultSSHPort)}, ":")
	}

	if !HostWithPortRegexp.MatchString(*targetHost) {
		*targetHost = strings.Join([]string{*targetHost, strconv.Itoa(*exposePort)}, ":")
	}

	return *sshServer, *targetHost, *exposePort, *sshUsername, nil
}
