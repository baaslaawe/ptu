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

	var sshUsername = flag.String("u", currentUser.Username, "username to connect SSH server")
	var sshServer = flag.String("s", "127.0.0.1:22", "SSH server (host[:port]) to connect")
	var forwardPort = flag.Int("p", 8080, "port to forward on the SSH server side")
	var remoteHost = flag.String("r", "127.0.0.1:3000", "remote host:port to forward connection")

	flag.Parse()

	if (*forwardPort < TCPPortMIN) || (*forwardPort > TCPPortMAX) {
		return "", "", 0, "", errors.New("forwarded port number is not in the valid range")
	}

	if !HostWithPortRegexp.MatchString(*sshServer) {
		*sshServer = strings.Join([]string{*sshServer, strconv.Itoa(DefaultSSHPort)}, ":")
	}

	if !HostWithPortRegexp.MatchString(*remoteHost) {
		*remoteHost = strings.Join([]string{*remoteHost, strconv.Itoa(*forwardPort)}, ":")
	}

	return *sshUsername, *sshServer, *forwardPort, *remoteHost, nil
}
